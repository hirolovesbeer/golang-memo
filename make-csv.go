package main

import (
    "encoding/csv"
    "log"
    "os"
    "math/rand"
    "time"
    "strconv"
//    "fmt"
    "sync"
    "runtime"
)

func failOnError(err error) {
    if err != nil {
        log.Fatal("Error:", err)
    }
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU()) 
/*
    file, err := os.OpenFile("./time-serise.csv", os.O_WRONLY|os.O_CREATE, 0600)
    failOnError(err)
    defer file.Close()

    err = file.Truncate(0)
    failOnError(err)

    writer := csv.NewWriter(file)
*/

    limit := 60 * 60 * 24 * 365

    start_time := time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
    sources := []string{"32", "33", "34", "35", "36", "37", "38", "39"}

    var wg sync.WaitGroup

    for _, sid := range sources {
        wg.Add(1)
        stime := start_time

        go func(sid2 string, stime2 int64, limit2 int){
            defer wg.Done()
            fname := sid2 + "-time-series.csv"
            file, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, 0600)
            failOnError(err)
            defer file.Close()
            err = file.Truncate(0)
            failOnError(err)
            writer := csv.NewWriter(file)

            for i := 0; i < limit2; {
                rand.Seed(time.Now().UTC().UnixNano())
                num := rand.Intn(1000)
                writer.Write([]string{strconv.FormatInt(stime2, 10), sid2, strconv.Itoa(num)})
                i++
                stime2++
            }
            writer.Flush()
//            wg.Done()
        }(sid, stime, limit)
    }
    wg.Wait()
}
