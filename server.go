package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    defer ln.Close()
    fmt.Println("Server listening on port 8080...")

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Connection error:", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message received:", message)
    conn.Write([]byte("Hello from server!\n"))
}
