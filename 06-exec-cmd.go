/* This is exercise Q29, in Chapter 7: Communications, from "Learning Go"
   by Miek Gieben.

   Write a program that takes a list of all running processes and prints
   how many child processes each parent has spawned.  The output should
   look like:

       Pid 0 has 2 children: [1 2]
       Pid 490 has 2 children: [1199 26524]
       Pid 1824 has 1 child: [7293]

   This was an example of running external commands via golang.  Other
   interesting reads are:

       http://golang.org/pkg/os/exec
       http://golang.org/pkg/os
       http://gobyexample.com/execing-processes
       http://stackoverflow.com/questions/10027477/golang-fork-process
       https://code.google.com/p/go/issues/detail?id=227
       https://github.com/golang-samples/exec

       http://github.com/dinedal/textql/blob/master/main.go

       You can detach your exec command from the go program itself by using
       StartProcess and setting stdin, stdout, and stderr to nil.  The above
       textql github example uses this approach to enter into a standalone
       sqlite3 session.
*/

package main

import (
    "fmt"
    "os/exec"
    "sort"
    "strconv"
    "strings"
)

func main() {
    ps := exec.Command("ps", "-e", "-opid,ppid,comm")
    output, _ := ps.Output()
    child := make(map[int][]int)
    for i, s := range strings.Split(string(output), "\n") {
        if i == 0 { continue } // skip header
        if len(s) == 0 { continue } // skip last line, itself
        f := strings.Fields(s)
        fpp, _ := strconv.Atoi(f[1]) // parent's pid
        fp, _ := strconv.Atoi(f[0])  // child's pid
        child[fpp] = append(child[fpp], fp)
    }

    schild := make([]int, len(child))
    i := 0
    for k, _ := range child { schild[i] = k; i++ }
    sort.Ints(schild)
    for _, ppid := range schild {
        fmt.Printf("Pid %d has %d child", ppid, len(child[ppid]))
        if len(child[ppid]) == 1 {
            fmt.Printf(": %v\n", child[ppid])
            continue
        }
        fmt.Printf("ren: %v\n", child[ppid])
    }
}
