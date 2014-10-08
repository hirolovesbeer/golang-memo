package main

import "fmt"

type Person struct {
	name string
}

// 値渡し
func (p Person) ChangeName() string {
	p.name = p.name + "くん"
	return p.name
}

// 参照渡し
func (p *Person) ChangeNameMore() string {
	p.name = p.name + "くんさん"
	return p.name
}

func main() {
	abe := Person{"ひろし"}
	fmt.Printf("%T\n", abe)
	fmt.Printf("%+v\n", abe)

	// 値渡し
	var hiroshi string
	hiroshi = abe.ChangeName()
	fmt.Println(hiroshi)
	// 値渡しなので、構造体のnameは変わらない
	fmt.Printf("%+v\n", abe)

	// ポインタ渡し
	hiroshi2 := abe.ChangeNameMore()
	fmt.Println(hiroshi2)
	// ポインタ渡しなので、構造体のnameは変わる
	fmt.Printf("%+v\n", abe)
}
