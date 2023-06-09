package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// オリジンのチェックを無効化して、クロスオリジン通信を許可します
		return true
	},
}

type Client struct {
	conn *websocket.Conn
	w    http.ResponseWriter
}

type RequestMsg struct {
	Message  string   `json:"message,omitempty"`
	ToClient []string `json:"to,omitempty"`
}

type ResponseMsg struct {
	Message  string   `json:"message"`
	Clients  string   `json:"clients,omitempty"`
	ToClient []string `json:"to,omitempty"`
}

var clients = make(map[string]*Client)

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

	remoteAdd := conn.RemoteAddr().String()
	log.Println("current client :", remoteAdd)
	if err != nil {
		// delete client from clients
		// delete(clients, remoteAdd)
		deleteClient(remoteAdd)
		log.Println("Failed to upgrade connection:", err)
		return
	}
	// add a client to  clients
	_, ok := clients[remoteAdd]
	if ok {
		// delete(clients, remoteAdd)
		deleteClient(remoteAdd)
	}

	clients[remoteAdd] = &Client{conn: conn, w: w}

	// get opened client list
	s := getClients()
	res := &ResponseMsg{Message: fmt.Sprintf("[%s] ログインしました。", remoteAdd), Clients: s}
	send2client(res, "")
	boardCastWebs(conn)
}

func boardCastWebs(srcConn *websocket.Conn) {
	defer srcConn.Close()
	for {
		_, message, err := srcConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Println("close message from the client: ", err)
			}
			log.Println("Failed to read message:", err)
			// delete(clients, srcConn.RemoteAddr().String())
			deleteClient(srcConn.RemoteAddr().String())
			break
		}
		log.Println("Received message:", string(message))

		r := &RequestMsg{}
		json.Unmarshal(message, r)

		// res := []byte(message)
		s := getClients()
		res := &ResponseMsg{Message: r.Message, Clients: s, ToClient: r.ToClient}
		send2client(res, srcConn.RemoteAddr().String())
	}

}

// send message to opened client
func send2client(r *ResponseMsg, src string) {
	res, err := json.Marshal(r)
	if err != nil {
		log.Println("response marshal err:", err)
		return
	}
Labels_Clients:
	for k, c := range clients {
		// defer c.conn.Close()
		// send to specified address
		// if no specified address,it will broadcast to all opened client exclude me
		if len(r.ToClient) != 0 {
			v := false
			for _, tc := range r.ToClient {
				if k == tc {
					v = true
				}
			}
			if !v {
				continue Labels_Clients
			}
		} else {
			if src == k {
				continue
			}
		}

		// send message to client
		err := c.conn.WriteMessage(websocket.TextMessage, res)
		if err != nil {
			log.Println("error:", err)
			// delete(clients, c.conn.RemoteAddr().String())
			deleteClient(c.conn.RemoteAddr().String())
			continue
		}
	}
}

// Get opened client
func getClients() string {
	sb := strings.Builder{}
	for client := range clients {
		sb.WriteString(client)
		sb.WriteString(`<br>`)
	}
	s := strings.TrimSuffix(sb.String(), "<br>")
	return s
}

func deleteClient(client string) {
	delete(clients, client)
	res := &ResponseMsg{Message: fmt.Sprintf("%sログアウトしました", client), Clients: getClients()}
	send2client(res, "")
}
