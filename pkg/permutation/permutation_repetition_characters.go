package permutation

import (
	"math"
	"sync"
)

func output(data []string, out chan<- string, wg *sync.WaitGroup) {
	for _, d := range data {
		out <- d
		wg.Done()
	}
}

func permutationChan(wg *sync.WaitGroup, out chan string, data string, std []string, n int) {
	var result []string
	for _, s := range std {
		var d = data
		d += s

		if len(d) == n {
			result = append(result, d)
		} else {
			permutationChan(wg, out, d, std, n)
		}
	}

	go output(result, out, wg)
}

func PermutationChan(std []string, n int) <-chan string {
	var (
		out   = make(chan string)
		wg    sync.WaitGroup
		count = int(math.Pow(float64(len(std)), float64(n)))
	)

	wg.Add(count)
	go func() {
		permutationChan(&wg, out, "", std, n)
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func Permutation(result []string, data string, std []string, n int) []string {
	for _, s := range std {
		var d = data
		d += s

		if len(d) == n {
			result = append(result, d)
		} else {
			result = Permutation(result, d, std, n)
		}
	}

	return result
}
