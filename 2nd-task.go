package main

import (
	"fmt"
	"os/exec"
	"net"
	"bufio"
//	"io/ioutil"
	"io"
	"strings"
	"bytes"
)

func main() {
/*
	grepCmd := exec.Command("/usr/bin/grep", "abe")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
*/

	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":8081")

	conn, _ := ln.Accept()

	r := bufio.NewReader(conn)
	for {
		grepCmd := exec.Command("/usr/bin/grep", "abe")

		message, _, err := r.ReadLine()
		fmt.Println("Message Received:", string(message))

		if err == io.EOF {
			break
		}

		grepCmd.Stdin = strings.NewReader(string(message))
		var out bytes.Buffer
		grepCmd.Stdout = &out
		err2 := grepCmd.Run()
		if err2 != nil {
			continue
			// fmt.Println("err2: ", err2)
		}
		fmt.Printf("in all caps: %q\n", out.String())

		// conn.Write([]byte("from server\n"))
		conn.Write([]byte(out.String() + "\n"))
	}
}