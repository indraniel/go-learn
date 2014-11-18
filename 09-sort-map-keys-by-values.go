package main

import (
	"fmt"
	"sort"
)

// A data structure to hold key/value pairs
type Pair struct {
	Key   string
	Value int
}

// A slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func main() {
	noble := map[string]int{
		"Radon":   86,
		"Xenon":   54,
		"Krypton": 36,
		"Argon":   18,
		"Neon":    10,
		"Helium":  2,
	}

	p := make(PairList, len(noble))

	i := 0
	for k, v := range noble {
		p[i] = Pair{k, v}
		i++
	}

	fmt.Printf("Pre-sorted: ")
	for _, k := range p {
		fmt.Printf("%s ", k.Key)
	}
	fmt.Println("")

	sort.Sort(p)

	fmt.Printf("Post-sorted: ")
	for _, k := range p {
		fmt.Printf("%s ", k.Key)
	}
	fmt.Println("")
}
