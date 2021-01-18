package main

import "fmt"

func main() {
	seq := []int{1, 2, 3, 4, 5}
	target := 5
	r, s := solve(seq, target)
	if r {
		fmt.Printf("%d回目の検索で「%d」が見つかりました。", s+1, target)
	} else {
		fmt.Printf("対象「%d」が見つかりませんでした。", target)
	}

}

func solve(argSeq []int, argTarget int) (bool, int) {
	exist := false
	index := 0
	for i := 0; i < len(argSeq); i++ {
		if argSeq[i] == argTarget {
			exist = true
			index = i
		}
	}
	return exist, index
}
