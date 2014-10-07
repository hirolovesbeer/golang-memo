package main

import (
    "net"
    "os"
)

func main() {
    strEcho := "Hello echo"
    servAddr := "localhost:9999"

    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    }

   for i := 0; i < 1000; i++ {
        _, err = conn.Write([]byte(strEcho))
        if err != nil {
            println("Write to server failed:", err.Error())
            os.Exit(1)
        }

//        println("write to server = ", strEcho)
    }

    println("write to server finished ")
    conn.Close()
}
