package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.NumCPU())
    fmt.Println(runtime.GOMAXPROCS(0))
    fmt.Println(runtime.NumGoroutine())
}
