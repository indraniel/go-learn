// This program was created to better understand:
// http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

package main

import (
    "fmt"
)

type Pets interface {
    Speak() string
}

type Dog struct {}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {}

func (c *Cat) Speak() string {
    return "Meow!"
}

func main() {
    d := Dog{}
    dptr := new(Dog)
    c := Cat{}

    fmt.Println(d.Speak())

    // from http://golang.org/ref/spec#Calls
    //   Selects automatically deference pointers to structs.  If x is a 
    //   pointer to a struct, x.y is shorthand for (*x).y;  if the field y
    //   is also a pointer to a struct, x.y.z is shorthand for (*(*x).y).z,
    //   and so on.  If x contains an anonymous field of type *A, where A is also
    //   a struct type, x.f is shorthand for (*x.A).f.
    fmt.Println(dptr.Speak())

    // from http://golang.org/ref/spec#Calls
    //   if x is addressable and &x's method set contains m,
    //   x.m() is shorthand for (&x).m()
    fmt.Println(c.Speak())

    pets := []Pets{d, &c}

    // the slice elements of pets are of type "Pets", not the
    // types "Dog" and "Cat"
    // based on the function defintions 
    //   -- a "Dog" type has a "Speak" method
    //   -- a "*Cat" type has a "Speak" method, not a "Cat" type
    //        (we defined the method to work on pointers)
    // from the Spec
    // For a variable x of type I where I is an interface type,
    // x.f denotes the actual method with name f of the value assigned
    // to x.  If there is no method with name f in the method set of I, the
    // selector expression is illegal

    for _, pet := range pets {
        fmt.Println(pet.Speak())
    }

    // Go cannot easily convert to a slice of interfaces.  Just converting
    // to an interface is easy, but to a slice is much more costly.  Bottom line:
    // you need to explicitly convert a slice into a type of interfaces
    // References:
    //  - http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
    //  - https://groups.google.com/forum/#!topic/golang-nuts/Il-tO1xtAyE
    //  - http://go-book.appspot.com/interfaces.html

    // Type assertions & Type switches
    // See -- http://golang.org/ref/spec#Type_assertions
}
