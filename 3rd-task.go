package main

import (
	"fmt"
	"os/exec"
	"net"
	"bufio"
	"io"
	"strings"
	"bytes"
)

func main() {
	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":8082")

	conn, _ := ln.Accept()

	r := bufio.NewReader(conn)
	for {
		// wcCmd := exec.Command("/usr/bin/wc", "-l")
		wcCmd := exec.Command("/usr/bin/wc")

		message, _, err := r.ReadLine()
		fmt.Println("Message Received:", string(message))

		if err == io.EOF {
			break
		}

		wcCmd.Stdin = strings.NewReader(string(message))
		var out bytes.Buffer
		wcCmd.Stdout = &out
		err2 := wcCmd.Run()
		if err2 != nil {
			continue
		}
		fmt.Printf("in all caps: %q\n", out.String())

		// conn.Write([]byte("from server\n"))
		conn.Write([]byte(out.String() + "\n"))
	}
}

