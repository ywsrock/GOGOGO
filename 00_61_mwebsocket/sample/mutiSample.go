package sample

import (
	"fmt"
	"net"
	"os"
)

type Client struct {
	conn net.Conn
}

func main() {
	ip := "0.0.0.0" // バインドするIPアドレス (0.0.0.0は全てのネットワークインターフェース)
	port := "8080"  // バインドするポート番号

	// TCPリスナーの作成
	listener, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("リスナーの作成に失敗しました:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("クライアントからの接続を待機しています...")

	clients := make(map[string]*Client) // 接続されたクライアントを管理するマップ

	for {
		// クライアントからの接続を受け入れる
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接続の受け入れに失敗しました:", err)
			os.Exit(1)
		}

		client := &Client{conn: conn}
		clientAddr := conn.RemoteAddr().String()

		clients[clientAddr] = client // クライアントをマップに追加

		go handleClient(client, clients) // クライアントごとの処理をゴルーチンで実行
	}
}

func handleClient(client *Client, clients map[string]*Client) {
	for {
		buffer := make([]byte, 1024)
		n, err := client.conn.Read(buffer)
		if err != nil {
			fmt.Printf("クライアント %s との接続が切断されました。\n", client.conn.RemoteAddr().String())
			delete(clients, client.conn.RemoteAddr().String()) // 切断されたクライアントをマップから削除
			return
		}

		message := string(buffer[:n])
		fmt.Printf("クライアント %s からの受信したデータ: %s\n", client.conn.RemoteAddr().String(), message)

		// クライアントに返信
		reply := "Hello, Client! I received your message."
		_, err = client.conn.Write([]byte(reply))
		if err != nil {
			fmt.Printf("クライアント %s への返信に失敗しました: %s\n", client.conn.RemoteAddr().String(), err.Error())
		}
	}
}
