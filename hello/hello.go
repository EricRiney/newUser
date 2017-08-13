package main

import "net"
import "fmt"
//import "bufio"
//import "os"

func main() {

  // connect to this socket
  // conn, _ := net.Dial("tcp", "127.0.0.1:9090")
  for { 
	conn, _ := net.Dial("tcp", "127.0.0.1:9090")
    // read in input from stdin
    // reader := bufio.NewReader(os.Stdin)
    // fmt.Print("Text to send: ")
	var text = "" 
	//http.ListenAndServe(":8080", nil)
	//, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text)
    // listen for reply
	var message = "abc"
	//, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: "+message)
  }
}