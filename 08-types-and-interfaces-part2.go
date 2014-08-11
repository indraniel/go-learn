package main

/*  
    It's possible to have an interface type as an anonymous field.
    However, you'll have to do a type assertion to access the underlying 
    real type attributes that are assigned to anonymous interface field.

    Assigning an interface type as an anonymous field to a struct is a
    technique used in "github.com/petar/GoLLRB"
*/

import "fmt"

type Foo interface {
    abracadabra() string
}

type Node struct {
    Foo
    name  string
    value int
}

type myStuff struct {
    id int
    label string
}

func (mi myStuff) abracadabra() string {
    return "abracadabra"
}

func main() {
    var thing myStuff = myStuff{id: 44, label: "Baz"}
    n := Node{Foo: thing,
              name: "junk",
              value: 100}

    // this won't work -- it won't compile
//    fmt.Printf("thing is: %d", n.Foo.id)

    // this will work -- via type assertion
    fmt.Printf("the 'id' on thing is: %d\n", n.Foo.(myStuff).id)
    fmt.Printf("the 'label' on thing is: %s\n", n.Foo.(myStuff).label)
}
