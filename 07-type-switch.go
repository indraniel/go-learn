package main

import (
    "fmt"
)

func Foo(x interface{}) {
    fmt.Printf("%#v\n", x)
    switch t := x.(type) {
    default:
        fmt.Printf("x is of type: %T\n", t)
    }
}

type Baz struct {
    entry1 string
    entry2 int
}

type Bizz struct {
    blarg int
    boo string
}

func main () {
    a := Baz{entry1: "hello", entry2: 0}
    b := Bizz{blarg: 100, boo: "Yo Yo"}

    Foo(a)
    Foo(&a)
    Foo(b)
    Foo(&b)
}
