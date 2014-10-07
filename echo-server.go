package main

import (
    "net"
    "fmt"
    "bufio"
)

func main() {
    ln, err := net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Printf("Error: %s", err)
        return
    }

    for {
        conn, _ := ln.Accept()
        go func(conn net.Conn) {
            reader := bufio.NewReader(conn)
            for {
                line, _, e := reader.ReadLine()
                if e != nil {
                    conn.Close()
                    return
                }
                fmt.Printf("> %s\n", line)
            }
        }(conn)
    }
}
