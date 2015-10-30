package main

import (
	"fmt"
	"os/exec"
)

func main() {
	output, err := exec.Command("/usr/bin/head", "-n", "10", "/var/log/system.log").Output()

	if err != nil {
		fmt.Println("error ", err)
	}

	fmt.Println("> output")
	fmt.Println(string(output))

	output2, err2 := exec.Command("sh", "-c", "/usr/bin/head -n 10 /var/log/system.log").Output()

        if err2 != nil {
                fmt.Println("error ", err2)
        }

        fmt.Println("> output2")
        fmt.Println(string(output2))
}
