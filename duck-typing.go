package main

import "fmt"

// 参照: http://otiai10.hatenablog.com/entry/2014/01/15/210445

// ふくたろうさんこんにちは
// interfaceはメソッドリストを持つ
type Fukutaro interface {
    Hello() string
}

// バクさんさようなら
// interfaceはメソッドリストを持つ
type Bakusan interface {
    Goodbye() string
}

// ふくたろうさんはバクさんとのダブルマン
type DoubleMan struct {
    greetingHello string
    greetingGoodbye string
}

func (d DoubleMan)Hello() string {
    return d.greetingHello
}

func (d DoubleMan)Goodbye() string {
    return d.greetingGoodbye
}

// 引数はふくたろうさん
// ふくたろうさん以外の引数は取らない
func SayHello(fuku Fukutaro) string {
    return fuku.Hello()
}

// 引数はバクさん
// バクさん以外の引数は取らない
func SayGoodbye(baku Bakusan) string {
    return baku.Goodbye()
}

func main() {
    /* 構造体 DoubleManに文字列をセット */
    dm := &DoubleMan{
        "こんにちは",
        "さようなら",
    }

    fmt.Printf("ダブルマン => %+v\n", dm)
    // ダブルマンなので、こんにちはが言える
    fmt.Printf("call dm.Hello: %s\n", dm.Hello())
    // ダブルマンなので、さようならが言える
    fmt.Printf("call dm.Goodbye: %s\n", dm.Goodbye())

    // 引数がふくたろうさんなのにダブルマンを渡すとこんにちはが言える
    fmt.Println(
        "Say Hello", SayHello(dm),
    )

    // 引数がバクさんなのにダブルマンを渡すとさようならが言える
    fmt.Println(
        "Say Goodbye", SayGoodbye(dm),
    )
}
