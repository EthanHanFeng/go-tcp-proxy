package main

import "fmt"
import "net"
import "bufio"
import "os"

func main() {
	serverPort := "6001"
	fmt.Println("Launching client ...")
	fmt.Println("Server Port:", serverPort)
	
	// loop forever to continuously send messages
	for {
		conn, err := net.Dial("tcp", ":"+serverPort)
		if err != nil {
			if _, t := err.(*net.OpError); t {
				fmt.Println("Connection error.")
			} else {
				fmt.Println("Unknown error: " + err.Error())
			}
			os.Exit(1)
		}
		reader := bufio.NewReader(os.Stdin)
		// read line by line
		text, _ := reader.ReadString('\n')
		fmt.Printf("Sending: %s\n", text)
		
		_, writeErr := conn.Write([]byte(text))
		if writeErr != nil {
			fmt.Println("Error writing to stream")
			break
		}
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message from server: " + message)
		fmt.Println("---------------------")
		conn.Close()
	}
}
