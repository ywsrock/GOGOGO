package frontandserver

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// オリジンのチェックを無効化して、クロスオリジン通信を許可します
		return true
	},
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	defer conn.Close()

	for {
		// クライアントからのメッセージを受信します
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}

		log.Println("Received message:", string(message))

		// クライアントにメッセージを送信します
		response := []byte("Server received: " + string(message))
		err = conn.WriteMessage(websocket.TextMessage, response)
		if err != nil {
			log.Println("Failed to send message:", err)
			break
		}
	}
}
