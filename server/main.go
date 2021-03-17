package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    //CheckOrigin:  func(r *http.Request) bool { return true },
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        conn,  err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
        if err != nil {
            fmt.Println(err)
        }

        for {
            // Read message from browser
            msgType, msg, err := conn.ReadMessage()
            if err != nil {
                return
            }

            // Print the message to the console
            fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

            // Write message back to browser
            if err = conn.WriteMessage(msgType, msg); err != nil {
                return
            }
        }
    })

    http.HandleFunc("/client", func(w http.ResponseWriter, r *http.Request) { // I dont need this 
    http.ServeFile(w, r, "client.html")
  })

    http.ListenAndServe(":8080", nil)
}
