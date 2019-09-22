package main

import (
	"fmt"
	"net"
	"io"
	"flag"
)

func main() {
	var target string 
	var port int

	flag.StringVar(&target, "remoteSrv", "127.0.0.1:6000", "the remote server (<host>:<port>)")
	flag.IntVar(&port, "proxyPort", 6001, "The proxy port")
	flag.Parse()

	fmt.Println("Launching proxy ...")
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Invalid port!")
	}
	targetAddr, err := net.ResolveTCPAddr("tcp", target)
	if err != nil {
		fmt.Println("Invalid remote server address!")
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println("Failed to start proxy!")
	}
	fmt.Println("Proxy server is started on port:", port)
	defer listener.Close()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Failed to accept connections.")
		}
		go handleConnection(conn, targetAddr)
	}
	
}

func handleConnection(conn *net.TCPConn, targetAddr *net.TCPAddr) {
	defer conn.Close()
	fmt.Println("Connected to client:", conn.RemoteAddr())
	client, err := net.DialTCP("tcp", nil, targetAddr)
	if err != nil {
		fmt.Println("Could not connect to remote server:", err)
		return
	}
	defer client.Close()
	fmt.Printf("Connection to server '%v' established!\n", client.RemoteAddr())
	
	stop := make(chan bool)

	go func() {
		_, err := io.Copy(client, conn)
		fmt.Println(err)
		stop <- true
	}()

	_, cpErr := io.Copy(conn, client)
	if cpErr != nil {
		fmt.Println("Could not copy from server to client")
	}
	<-stop

}