// This is program is based on reading:
// http://blog.golang.org/go-slices-usage-and-internals
// http://openmymind.net/The-Minimum-You-Need-To-Know-About-Arrays-And-Slices-In-Go/
//
// tl;dr
//   - passing a slice into a function is akin to passing by pointer address value

package main

import "fmt"

func main() {
    xPtr := new(int)
    *xPtr = 5

    fmt.Println("-------- Variable Playground -----------")
    fmt.Printf("The value of xPtr is : %#v\n", xPtr)
    fmt.Printf("The value of *xPtr is: %#v\n", *xPtr)
    fmt.Println("")

    fmt.Println("-------- Slice Playground -----------")
    r := []int{0,1,2,3}

    fmt.Println("+++ Inspection Tests +++")
    fmt.Printf("The address of r is : %#p\n", &r)
    fmt.Printf("The contents of r is: %v\n", r)
    inspect_slice(r)
    fmt.Println("")
    inspect_slice_ptr(&r)
    fmt.Println("")

    fmt.Println("+++ Update Tests +++")
    fmt.Println("-> By regular value")
    fmt.Printf("before: %v\n", r)
    update_pure_slice_regular(r)
    fmt.Printf("after: %v\n", r)

    fmt.Println("-> By pointer")
    fmt.Printf("before: %v\n", r)
    update_pure_slice_pointer(&r)
    fmt.Printf("after: %v\n", r)
    fmt.Println("")

    fmt.Println("-------- Array Playground -----------")
    a := [4]int{4,5,6,7}

    fmt.Println("+++ Inspection Tests +++")
    fmt.Printf("The address of a is : %#p\n", &a)
    fmt.Printf("The contents of a is: %v\n", a)
    inspect_array(a)
    fmt.Println("")
    inspect_array_ptr(&a)
    fmt.Println("")

    fmt.Println("+++ Update Tests (pure array) +++\n")
    fmt.Println("-> By regular value")
    fmt.Printf("before: %v\n", a)
    update_array_regular(a)
    fmt.Printf("after: %v\n\n", a)

    fmt.Println("-> By pointer")
    fmt.Printf("before: %v\n", a)
    update_array_pointer(&a)
    fmt.Printf("after: %v\n\n", a)


    fmt.Println("+++ Update Tests (via slice) +++\n")
    fmt.Println("-> By regular value")
    sl := a[:]
    fmt.Printf("(a) before: %v\n", a)
    fmt.Printf("(sl) before: %v\n", sl)
    update_array_slice_regular(sl)
    fmt.Printf("(a) after: %v\n", a)
    fmt.Printf("(sl) after: %v\n\n", sl)

    fmt.Println("-> By pointer")
    sl = a[:]
    fmt.Printf("(a) before: %v\n", a)
    fmt.Printf("(sl) before: %v\n", sl)
    update_array_slice_pointer(&sl)
    fmt.Printf("(a) after: %v\n", a)
    fmt.Printf("(sl) after: %v\n", sl)
}

func update_array_slice_regular(stuff []int) {
    stuff[2] = 55
}

func update_array_slice_pointer(stuff *[]int) {
    s := *stuff
    s[2] = 66
}

func update_pure_slice_regular(stuff []int) {
    stuff[2] = 11
}

func update_pure_slice_pointer(stuff *[]int) {
//    s := *stuff
//    s[2] = 11
    (*stuff)[2] = 22
}

func update_array_regular(stuff [4]int) {
    stuff[2] = 33
}

func update_array_pointer(stuff *[4]int) {
    stuff[2] = 44
}

func inspect_slice(stuff []int) {
    fmt.Println("inside inspect_slice")
    fmt.Printf("The address of stuff is : %p\n", &stuff)
    fmt.Printf("The contents of stuff is: %v\n", stuff)
}

func inspect_slice_ptr(stuff *[]int) {
    fmt.Println("inside inspect_slice_pointer")
    fmt.Printf("The address of stuff is : %p\n", stuff)
    fmt.Printf("The contents of stuff is:%v\n", stuff)
}

func inspect_array(argh [4]int) {
    fmt.Println("inside inspect_array")
    fmt.Printf("The address of argh is : %p\n", &argh)
    fmt.Printf("The contents of argh is: %v\n", argh)
}

func inspect_array_ptr(argh *[4]int) {
    fmt.Println("inside inspect_array_ptr")
    fmt.Printf("The address of argh is : %p\n", argh)
    fmt.Printf("The contents of argh is: %v\n", *argh)
}

/* ------ Output ------
    $ go run 04-array-slices-and-pointers.go
    -------- Variable Playground -----------
    The value of xPtr is : (*int)(0x20817e170)
    The value of *xPtr is: 5
    
    -------- Slice Playground -----------
    +++ Inspection Tests +++
    The address of r is : 2081b2000
    The contents of r is: [0 1 2 3]
    inside inspect_slice
    The address of stuff is : 0x2081b2060
    The contents of stuff is: [0 1 2 3]
    
    inside inspect_slice_pointer
    The address of stuff is : 0x2081b2000
    The contents of stuff is:&[0 1 2 3]
    
    +++ Update Tests +++
    -> By regular value
    before: [0 1 2 3]
    after: [0 1 11 3]
    -> By pointer
    before: [0 1 11 3]
    after: [0 1 22 3]
    
    -------- Array Playground -----------
    +++ Inspection Tests +++
    The address of a is : 2081b2140
    The contents of a is: [4 5 6 7]
    inside inspect_array
    The address of argh is : 0x2081b2180
    The contents of argh is: [4 5 6 7]
    
    inside inspect_array_ptr
    The address of argh is : 0x2081b2140
    The contents of argh is: [4 5 6 7]
    
    +++ Update Tests (pure array) +++
    
    -> By regular value
    before: [4 5 6 7]
    after: [4 5 6 7]
    
    -> By pointer
    before: [4 5 6 7]
    after: [4 5 44 7]
    
    +++ Update Tests (via slice) +++
    
    -> By regular value
    (a) before: [4 5 44 7]
    (sl) before: [4 5 44 7]
    (a) after: [4 5 55 7]
    (sl) after: [4 5 55 7]
    
    -> By pointer
    (a) before: [4 5 55 7]
    (sl) before: [4 5 55 7]
    (a) after: [4 5 66 7]
    (sl) after: [4 5 66 7]
*/

