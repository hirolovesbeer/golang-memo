package main

import (
    "flag"
    "bufio"
    "fmt"
    "os"
    "net/http"
    "encoding/json"
    "bytes"
    "errors"
    "strings"
    "strconv"
)

/*
  reference
  https://gowalker.org/github.com/rossdylan/influxdbc
*/

type InfluxDB struct {
    host     string
    database string
    username string
    password string
}

func (db InfluxDB) SeriesURL() string {
    return fmt.Sprintf("http://%s/db/%s/series?u=%s&p=%s", db.host, db.database, db.username, db.password)
}

func (db InfluxDB) WriteSeries(s []Series) error {
    url := db.SeriesURL()
    _, err := PostStruct(url, s)
    return err
}

func PostStruct(url string, reqStruct interface{}) (string, error) {
    fmt.Println(url)
    marshalled, err := json.Marshal(reqStruct)
    marshalled = bytes.ToLower(marshalled)
    if err != nil {
        panic(err)
    }

    buf := bytes.NewBuffer(marshalled)
    result, err := http.Post(url, "application/json", buf)
    if err != nil {
        panic(err)
    }

    defer result.Body.Close()
    result_buf := new(bytes.Buffer)
    result_buf.ReadFrom(result.Body)
    if result.StatusCode != 200 {
        return "", errors.New(result_buf.String())
    }

    return result_buf.String(), nil
}

type Series struct {
    Name       string
    Columns    []string
//    Points     [][]string
    Points     [][]int
}

func NewSeries(name string, cols ...string) *Series {
    s := new(Series)
    s.Name = name
    s.Columns = cols
//    s.Points = make([][]string, 0)
    s.Points = make([][]int, 0)
    return s
}

// func (s *Series) AddPoint(point ...string) {
func (s *Series) AddPoint(point ...int) {
    s.Points = append(s.Points, point)
}


func main() {
    flag.Parse()

    fp, err := os.Open(flag.Arg(0))

    if err != nil {
        fmt.Println("Open error")
        return
    }

    scanner := bufio.NewScanner(fp)

    /* influxdb connection info */
    database := InfluxDB{"hostname or IPaddr:8086", "time_value_datas", "root", "root"}
    series := NewSeries("test", "time", "sid", "value")

    i := 0
    limit := 100000 //31536000
    for scanner.Scan() {
        line := scanner.Text()
        cols := strings.Split(line, ",")

        time, _ := strconv.Atoi(cols[0])
        sid, _  := strconv.Atoi(cols[1])
        value, _  := strconv.Atoi(cols[2])

        series.AddPoint(time, sid, value)
        i++

        if i == limit {
            sarray := make([]Series, 1)
            sarray[0] = *series
            err = database.WriteSeries(sarray)

            i = 0
            series = NewSeries("test", "time", "sid", "value")
            continue
//            break
        }
    }

/*
    sarray := make([]Series, 1)
    sarray[0] = *series
    err = database.WriteSeries(sarray)
*/

    if err != nil {
        fmt.Println(err)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}
