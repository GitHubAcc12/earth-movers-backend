package math

import (
	"github.com/cannona/choose"
	"log"
	"strconv"
)

func Compositions(n int, k int) [][]int {
	n_len := int(choose.Choose(int64(n+k-1), int64(k-1)))
	log.Print("Len: " + strconv.Itoa(n_len))
	all := [][]int{}
	recCompositions(n, k, []int{}, &all)
	return all
}

func recCompositions(n int, k int, current []int, all_comps *[][]int){
	if len(current) == k && sum(current) == n {
		*all_comps = append(*all_comps, current)
	} else if len(current) < k {
		for i := 0; i < n - sum(current) + 1; i++ {
			tmp_current := make([]int, len(current))
			copy(tmp_current, current)
			tmp_current = append(tmp_current, i)
			recCompositions(n, k, tmp_current, all_comps)
		}
	}
}


func sum(input []int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return sum
}