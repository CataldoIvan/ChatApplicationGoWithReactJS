package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader=websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,

	// We'll need to check the origin of our connection
  // this will allow us to make requests from our React
  // development server to here.
  // For now, we'll do no checking and just allow any connection
	CheckOrigin:func (r *http.Request) bool {
		return true
	},
}

var clientes []websocket.Conn
// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn)  {
	
	for{
		//read message
		messageType,p,err:=conn.ReadMessage()
		if err != nil {
			fmt.Println("",err)
			return
		}
		clientes=append(clientes, *conn)
		fmt.Println("llego a este punto:",string(p))
		//Write messaje
		for _,client:=range clientes{

			if err:=client.WriteMessage(messageType,p);err!=nil{
				fmt.Println("fg",err)
				return
			}
			conn.WriteJSON(messageType)
		}

		
	}
}

//define our Websocjet endpoint
func serveWs(c *gin.Context){
	// upgrade this connection to a WebSocket
  // connection
  data,_:=c.GetRawData()
  fmt.Println(data)
	ws,err:=upgrader.Upgrade(c.Writer,c.Request,nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	 // listen indefinitely for new messages coming
  // through on our WebSocket connection
	reader(ws)	
}	

func main()  {
	router:=gin.Default()
	router.GET("/",func (c *gin.Context)  {
		c.JSON(
			200,
			gin.H{
			"mensaje":"hola mundo",
			})
	})

	router.GET("/ws",serveWs)
	router.Run()
	fmt.Println("chat App V0.01")
	
}