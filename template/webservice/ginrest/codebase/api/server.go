package api
import(
"github.com/gin-gonic/gin"
)
//WebServer struct
type WebServer struct {
	// gin server reference
	server *gin.Engine
	// channel for cleanup notification
	cleanUp chan int
}


//Create this will load default server config
func Create(cleanup chan int) *WebServer {

	server := gin.Default()
	
	server.Use(gin.Recovery())

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
	ws.server.Run(port)
}

