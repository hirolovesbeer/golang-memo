package main

import (
    "encoding/json"
    "fmt"
    "time"
)

type Color int

type Personal struct {
    Number   int `json:"ReferenceNumber"`
    *City
    Company  string
    DateTime time.Time
    color    Color
}

type City struct {
    CityNumber string
    CityName   string
    Address    string
    PostalCode int
}

const (
    Red Color = iota
    Green
    Blue
    Yellow
    White
    Black
)

func (person *Personal) ChangeColor(c Color) {
    person.color = c
}

func (person *Personal) ToJson() ([]byte, error) {
    b, err := json.Marshal(person)
    return b, err
}

func main() {
    var personJson = []byte(
        `{"ReferenceNumber":12345,
        "CityNumber":"67890",
        "CityName":"Kasukabe",
        "Address":"central",
        "PostalCode":1234567,
        "Company":"hogehoge",
        "DateTime":"2014-10-29T17:23:00Z",
        "Color":2}`)

    var person Personal

    err := json.Unmarshal(personJson, &person)

    if err != nil {
        fmt.Println("error:", err)
    }

    fmt.Printf("%+v\n", person)
    fmt.Printf("%+v\n", person.City)
}
