package main

import (
    "flag"
    "bufio"
    "fmt"
    "os"
    "encoding/json"
)

type Request struct {
    Name string     `json:"name"`
    Colums []string `json:"colums"`
    Points [][]string `json:"points"`
}

func main() {
    flag.Parse()

    fp, err := os.Open(flag.Arg(0))

    if err != nil {
        fmt.Println("Open error")
        return
    }

    lines := [][]string{}
    scanner := bufio.NewScanner(fp)

    i := 0
    for scanner.Scan() {
        var line []string
        line = append(line, scanner.Text())
//        fmt.Println(line)
        lines = append(lines, line)
//        fmt.Println(scanner.Text()) // Println will add back the final '\n'
        i++

        if i == 10000 {
            break
        }
    }
    fmt.Println(lines)

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }

    /* build json requests */
    request := new(Request)
    request.Name = "time_value_datas"
    request.Colums = []string{"time", "line"}
    request.Points = lines

    bytes, err := json.MarshalIndent(request, "", "\t")
    fmt.Printf("%s", bytes)
}
