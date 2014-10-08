package main

import "fmt"

// 参照: http://otiai10.hatenablog.com/entry/2014/01/15/220136

type MyBase struct {
    count int
}

func (b *MyBase)Increment() int {
    b.count++
    return b.count
}

type MyProperty struct {
    Value string
}

type AnotherProperty struct {
    AnotherValue string
}

// child class
type SubStruct struct {
    *MyBase //embed
    SomeProperty *MyProperty // alias?
    *AnotherProperty //embed
}

func (s *SubStruct)IncrementByTwo() int {
    s.Increment()
    return s.Increment()
}

func main() {
    sub := &SubStruct{
        &MyBase{0},
        &MyProperty{"プロパティ1"},
        &AnotherProperty{"プロパティ2"},
    }

    fmt.Printf("sub.Increment => %d\n", sub.Increment()) // 1
    fmt.Printf("sub.IncrementByTwo => %d\n", sub.IncrementByTwo()) // 3
    fmt.Printf("sub.IncrementByTwo => %d\n", sub.IncrementByTwo()) // 5
    fmt.Printf("sub.SomeProperty.Value => %s\n", sub.SomeProperty.Value) // プロパティ1
    fmt.Printf("sub.AnotherValue => %s\n", sub.AnotherValue) // プロパティ2
}
