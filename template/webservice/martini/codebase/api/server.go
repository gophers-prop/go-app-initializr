package api
import(
	"github.com/go-martini/martini"
)
//WebServer struct
type WebServer struct {
	// gin server reference
	server *martini.ClassicMartini
	// channel for cleanup notification
	cleanUp chan int
}


//Create this will load default server config
func Create(cleanup chan int) *WebServer {

	server :=  martini.Classic()
	
	server.Use(martini.Recovery())
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
	ws.server.RunOnAddr(port)
}

