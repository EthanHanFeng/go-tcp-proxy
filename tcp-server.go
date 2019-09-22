package main

import "fmt"
import "net"
import "bufio"

func main() {
	port := "6000"
	fmt.Printf("Launching server...\nPort:%s\n", port)
	// listen on port 6000 for client connection
	listen, _ := net.Listen("tcp", ":"+port)
	defer listen.Close()
	
	// loop forever to handle client connections
	for {
		// accept connection
		conn, _ := listen.Accept()

		// will listen for message util the first '\n', read message line by line
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Read error.")
			conn.Close()
		}
		fmt.Println("Received: ", string(msg))
		conn.Write([]byte(msg))
		conn.Close()
	}

}
