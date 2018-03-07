package main

import (
	"fmt"
	"github.com/wangbin/jiebago"
	"io/ioutil"
	"os"
	"sort"
)

var seg jiebago.Segmenter
var result map[string]int

type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func xinit() {
	seg.LoadDictionary("dict.txt")
}

func print(ch <-chan string) {
	for word := range ch {
		if v, ok := result[word]; ok {
			v++
			result[word] = v
		} else {
			result[word] = 1
		}
	}

	pairResult := sortMapByValue(result)

	for _, v := range pairResult {
		word, value := v.Key, v.Value
		pos, _ := seg.Pos(word)
		fmt.Printf("[%s / %s / %d]\n", word, pos, value)
	}

	fmt.Println()
}

func main() {
	xinit()

	result = make(map[string]int)

	f, err := os.Open("weicheng.txt")
	if err != nil {
		return
	}
	str, _ := ioutil.ReadAll(f)
	txt := string(str[:])
	print(seg.Cut(txt, false))
}
