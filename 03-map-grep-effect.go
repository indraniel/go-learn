package main

import "fmt"

// Based on my readings so far, this is how I could do a 
// simple "map", or Perl-like "grep" feature on a array or slice.

func main () {
    items := []string{"one", "two", "three"}
    filtered := []string{}

    for _, item := range items {
        if len(item) == 3 {
            filtered = append(filtered, item)
        }
    }

    fmt.Println("items: ", items)
    fmt.Println("filtered: ", filtered)
}
