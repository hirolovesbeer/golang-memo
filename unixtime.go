package main

import (
    "fmt"
    "time"
)

func main() {
   fmt.Println(time.Now().Unix()) 
   fmt.Println(time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC).Unix())
//   fmt.Println(time.Now().UnixNano()) 
}
