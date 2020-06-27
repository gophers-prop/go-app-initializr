package api
import(
	"goji.io"
	"net/http"
	{{ if .Logging.ImportPath }}
	"{{ .Logging.ImportPath }}"
	{{end}}
)
//WebServer struct
type WebServer struct {
	server *goji.Mux
	// channel for cleanup notification
	cleanUp chan int
}


//Create this will load default server config
func Create(cleanup chan int) *WebServer {

	server :=  goji.NewMux()
	
	return &WebServer{
		server:  server,
		cleanUp: cleanup,
	}
}

//Start registers all routes and then starts the server
func (ws *WebServer)Start(port string){
	//registering route
	ws.registerRoute()
	//starting server
	serverErr := http.ListenAndServe(port, ws.server)
	if serverErr != nil{
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Server.ERROR}} {{end}}
	return
	}
	{{ if .Logging.ImportPath }} {{ .Logging.Messages.Server.SUCCESS}} {{end}}
}

