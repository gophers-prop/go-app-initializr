package engine

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"errors"
	"fmt"
	"go-initializer/consts"
	"go-initializer/logger"
	"go-initializer/types"
	"go-initializer/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

//GenerateTemplateRequest request payload for generate template
type GenerateTemplateRequest struct {
	AppName             string `form:"appname" json:"appname" xml:"appname"  binding:"required"`
	AppType             string `form:"apptype" json:"apptype" xml:"apptype"  binding:"required"`
	Library             string `form:"library" json:"library" xml:"library"  binding:"required"`
	DependencyManagment string `form:"dependencies" json:"dependencies" xml:"dependencies" `
	LoggingFramework    string `form:"loggingframework" json:"loggingframework"`
	OutputFormat        string `form:"outputformat" json:"outputformat"`
	requestTime         string
	outputFolder        string
	sourceFolder        string
	outputArchive       string
	Functionalities     []Functionality `json:"functionalities" binding:"required"`
}

func GenerateTemplate(ctx *gin.Context) {

	request, err := validateGenerateTemplateRequest(ctx)
	if err != nil {
		ctx.JSON(err.HTTPStatus, gin.H{"error": err.Error})
		ctx.Abort()
		return
	}
	err = createOuputFolder(request)
	if err != nil {
		ctx.JSON(err.HTTPStatus, gin.H{"error": err.Error})
		ctx.Abort()
		return
	}

	if request.OutputFormat == "tar" {
		ctx.Header("Content-Type", "application/octet-stream")
		err = createTar(request)
	} else {
		ctx.Header("Content-Type", "application/zip")
		err = createZip(request)
	}

	if err != nil {
		ctx.JSON(err.HTTPStatus, gin.H{"error": err.Error})
		ctx.Abort()
		return
	}
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")

	// zip is default output format . This is to support cli feature

	ctx.Header("Content-Disposition", "attachment; filename="+request.AppName+"."+request.OutputFormat)

	ctx.Header("File-name", request.AppName+"."+request.OutputFormat)
	ctx.File(request.outputArchive)

	cleanUpError := request.Cleanup()
	if cleanUpError != nil {
		logger.Log(request.AppName).Warnf("Error while cleaning up outputs for request %+v", request)
	}
}

func validateGenerateTemplateRequest(ctx *gin.Context) (*GenerateTemplateRequest, *types.Error) {
	var request GenerateTemplateRequest
	if err := ctx.ShouldBind(&request); err != nil {
		logger.Log(request.AppName).Errorf("Error while reading request %+v", request)
		return nil, utils.GetError(err, http.StatusBadRequest)
	}
	fmt.Println(request)
	logger.Log(request.AppName).Errorf("%+v", request)
	logger.Log(request.AppName).Infof("Received request for app generation %+v", request)

	request.requestTime = fmt.Sprintf("%d", time.Now().Unix())

	sourcePath, _ := utils.GetWorkingDir()
	request.outputFolder = filepath.Join(sourcePath, consts.OUTPUT_FOLDER, request.AppName+request.requestTime)
	request.outputArchive = filepath.Join(sourcePath, consts.OUTPUT_ZIP, request.AppName+"."+request.OutputFormat)

	if !utils.AppTypeExists(request.AppType) {
		return nil, utils.GetError(errors.New("requested apptype does not exists"), http.StatusBadRequest)
	}

	request.Library = strings.ReplaceAll(request.Library, "/", string(os.PathSeparator))

	if !utils.LibExists(filepath.Join(request.AppType, request.Library)) {
		return nil, utils.GetError(errors.New("requested library does not exists"), http.StatusBadRequest)
	}

	request.sourceFolder = filepath.Join(sourcePath, "template", request.AppType, request.Library, "codebase")

	return &request, nil
}

func createOuputFolder(request *GenerateTemplateRequest) *types.Error {

	err := os.Mkdir(request.outputFolder, 0777)
	if err != nil {
		logger.Log(request.AppName).Errorf("Error while creating output folder %s", err.Error())
		return utils.GetError(err, http.StatusInternalServerError)
	}

	config := getConfiguration(request)
	//check if fun is a valid functionality or not and inject the functionalities accordingly
	for _, fun := range request.Functionalities {
		if fun.Name != "" {
			fun.OutputFolder = request.outputFolder
			featureMap[fun.Name](fun)
		}
	}

	err = filepath.Walk(request.sourceFolder, func(filePath string, info os.FileInfo, err error) error {

		outputFileName := strings.TrimPrefix(filePath, request.sourceFolder)

		if outputFileName == "" {
			return nil
		}

		outputFileName = outputFileName[1:]
		if info.IsDir() {
			err := os.Mkdir(filepath.Join(request.outputFolder, outputFileName), 0777)
			if err != nil {
				return err
			}
		} else {
			t, err := template.ParseFiles(filePath)
			if err != nil {
				logger.Log(request.AppName).Errorf("Error while parsing file %s , Error : %+v", filePath, err)
				return err
			}

			f, err := os.Create(filepath.Join(request.outputFolder, outputFileName))
			if err != nil {
				logger.Log(request.AppName).Errorf("Error while creating file file %s , Error : %+v", f.Name(), err)
				return err
			}

			err = t.Execute(f, config)
			if err != nil {
				logger.Log(request.AppName).Errorf("Error while generating tempalte file  Error : %+v", err)
				return err
			}
			f.Close()

		}
		return nil
	})
	if err != nil {
		logger.Log(request.AppName).Debugf("Error while generating output folder. Error : %s", err.Error())
		return utils.GetError(err, http.StatusInternalServerError)
	}

	return nil

}

func createZip(request *GenerateTemplateRequest) *types.Error {

	zipfile, err := os.Create(request.outputArchive)
	if err != nil {
		logger.Log(request.AppName).Errorf("Error while creating output zip file %s", err.Error())
		return utils.GetError(err, http.StatusInternalServerError)
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	err = filepath.Walk(request.outputFolder, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}
		if request.outputFolder == path {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, request.outputFolder)[1:]

		if info.IsDir() {
			header.Name += string(os.PathSeparator)
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)

		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)

		file.Close()
		return nil

	})
	if err != nil {
		logger.Log(request.AppName).Debugf("Error while generating zip file. Error : %s", err.Error())
		return utils.GetError(err, http.StatusInternalServerError)
	}
	return nil
}

// tarrer walks paths to create tar file tarName
func createTar(request *GenerateTemplateRequest) *types.Error {
	tarFile, err := os.Create(request.outputArchive)
	if err != nil {
		logger.Log(request.AppName).Errorf("Error while creating output tar file %s", err.Error())
		return utils.GetError(err, http.StatusInternalServerError)
	}
	defer func() {
		err = tarFile.Close()
	}()

	absTar, err := filepath.Abs(request.outputArchive)
	if err != nil {
		logger.Log(request.AppName).Errorf("Error while creating output tar file %s", err.Error())
		return utils.GetError(err, http.StatusInternalServerError)
	}

	// enable compression if file ends in .gz
	tw := tar.NewWriter(tarFile)
	if strings.HasSuffix(request.outputArchive, ".gz") || strings.HasSuffix(request.outputArchive, ".gzip") {
		gz := gzip.NewWriter(tarFile)
		defer gz.Close()
		tw = tar.NewWriter(gz)
	}
	defer tw.Close()

	var paths []string

	paths = append(paths, request.outputFolder)

	for _, path := range paths {
		// validate path
		path = filepath.Clean(path)
		absPath, err := filepath.Abs(path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if absPath == absTar {
			fmt.Printf("tar file %s cannot be the source\n", request.outputArchive)
			continue
		}
		if absPath == filepath.Dir(absTar) {
			fmt.Printf("tar file %s cannot be in source %s\n", request.outputArchive, absPath)
			continue
		}

		walker := func(file string, finfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// fill in header info using func FileInfoHeader
			hdr, err := tar.FileInfoHeader(finfo, finfo.Name())
			if err != nil {
				return err
			}

			relFilePath := file
			if filepath.IsAbs(path) {
				relFilePath, err = filepath.Rel(path, file)
				if err != nil {
					return err
				}
			}
			// ensure header has relative file path
			hdr.Name = relFilePath

			if err := tw.WriteHeader(hdr); err != nil {
				return err
			}
			// if path is a dir, dont continue
			if finfo.Mode().IsDir() {
				return nil
			}

			// add file to tar
			srcFile, err := os.Open(file)

			if err != nil {
				return err
			}
			defer srcFile.Close()
			_, err = io.Copy(tw, srcFile)
			if err != nil {
				return err
			}
			return nil
		}

		// build tar
		if err := filepath.Walk(path, walker); err != nil {
			logger.Log(request.AppName).Debugf("Error while generating output folder. Error : %s", err.Error())
			return utils.GetError(err, http.StatusInternalServerError)
		}
	}
	return nil
}

func getConfiguration(req *GenerateTemplateRequest) types.Configuration {
	var res types.Configuration
	res.AppName = req.AppName
	for _, feat := range req.Functionalities {
		switch feat.Name {
		case LOGGING:
			loggingframework, err := readLogJson(feat.Library)
			if err != nil {
				fmt.Println(err)
				res.Logging = types.LoggingFramework{}
			} else {
				res.Logging = *loggingframework
			   
			}
		}
	}

	return res

}

//Cleanup perfoming cleanup activities
func (request *GenerateTemplateRequest) Cleanup() error {

	//cleaning output folder

	err := os.RemoveAll(request.outputFolder)
	if err != nil {
		return fmt.Errorf("Error cleaning up output folder for %s", request.outputFolder)
	}
	// removing zip file

	err = os.RemoveAll(request.outputArchive)
	if err != nil {
		return fmt.Errorf("Error cleaning up Zip file for %s", request.outputArchive)
	}
	return nil

}
