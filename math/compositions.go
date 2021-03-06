package math

import (
	"github.com/cannona/choose"
	"log"
	"strconv"

	"earth-movers-backend/tools"
)

func Compositions(n int, k int) [][]int {
	n_len := int(choose.Choose(int64(n+k-1), int64(k-1)))
	log.Print("Will compute " + strconv.Itoa(n_len) + " distributions.")
	all := [][]int{}
	recCompositions(n, k, []int{}, &all)
	return all
}

func recCompositions(n int, k int, current []int, all_comps *[][]int){
	if len(current) == k && tools.Sum(current) == n {
		*all_comps = append(*all_comps, current)
	} else if len(current) < k {
		for i := 0; i < n - tools.Sum(current) + 1; i++ {
			tmp_current := make([]int, len(current))
			copy(tmp_current, current)
			tmp_current = append(tmp_current, i)
			recCompositions(n, k, tmp_current, all_comps)
		}
	}
}


