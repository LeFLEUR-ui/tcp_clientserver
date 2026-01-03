package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    fmt.Println("Connected to server. Type a message:")
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')

    fmt.Fprintf(conn, text)
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Server reply:", message)
}
