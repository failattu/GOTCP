package main

import (
	"fmt"
	"net"
	"os"
)
func server(){
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		fmt.Print(conn)

		
	}
}
func client(){
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Print(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "message\n")
}

func main() {
	fmt.Printf("hello, world\n")
	
	if len(os.Args) > 1 {
		if os.Args[1] == "server" {
			server()
		}else if os.Args[1] == "client"{
			client()
		}else{
			fmt.Printf("Goodbye\n")
		}
	}
}
