// This program is based on reading:
// http://golang.org/pkg/bufio/#Scanner

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("hello-world.txt")
    if err != nil {
        panic(fmt.Sprintf("Yikes! Couldn't open file! : %s", err))
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    line_num := 1
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Printf("%d: %s\n", line_num, line)
        line_num++
        if err := scanner.Err(); err != nil {
            fmt.Println(fmt.Sprintf("Yikes! Trouble reading file : %s"))
        }
    }
}
