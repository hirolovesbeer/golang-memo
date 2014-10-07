package main

import (
    "net"
    "bufio"
    "fmt"
    "os"
)

func main() {
    listener, err := net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Printf("couldn't start listening: %s ", err)
        os.Exit(1)
    }
    conns := clientConns(listener)

    for {
        go handleConn(<-conns)
    }
}

func clientConns(listener net.Listener) chan net.Conn {
    ch := make(chan net.Conn)
    i := 0
    go func() {
        for {
            client, err := listener.Accept()
            if client == nil {
                fmt.Printf("couldn't accept: %s", err)
                continue
            }
            i++
            fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
            ch <- client
        }
    }()
    return ch
}

func handleConn(client net.Conn) {
    b := bufio.NewReader(client)
    for {
        line, err := b.ReadBytes('\n')
        if err != nil {
            break
        }
        client.Write(line)
    }
}
