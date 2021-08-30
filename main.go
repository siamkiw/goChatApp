package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

type socketData struct {
	IP string `json:"IP"`
	Message string `json:"Message"`

}

func initRoutes(){
	http.HandleFunc("/ws", wsEndPoint)
	http.HandleFunc("/chat", indexPage)
	http.HandleFunc("/", homePage)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Home Page")
}

func indexPage(w http.ResponseWriter, r *http.Request){

	http.ServeFile(w, r, "./static/html/index.html")

}

func reader(conn *websocket.Conn, r *http.Request){
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := socketData{
			IP: r.RemoteAddr,
			Message: string(p),
		}

		value, err := json.Marshal(message)
		if err != nil {
			log.Println(err)
		}

		if err = conn.WriteMessage(messageType, value); err != nil {
			log.Println(err)
		}

	}
}

func wsEndPoint(w http.ResponseWriter, r *http.Request){
	upgrader.CheckOrigin = func(r *http.Request) bool {return true}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("IP address is : %v \n", r.RemoteAddr)

	log.Print("Client Successfully Connected.")

	reader(ws, r)

}

func main() {
	fmt.Println("Go WebSockets")
	initRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}