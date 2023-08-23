package helpers

import (
	"score_cat/types"
	"sync"
)

func RunAsync(funcs ...func() (interface{}, error)) []types.Result {
	var wg sync.WaitGroup
	results := make([]types.Result, len(funcs))

	for i, f := range funcs {
		wg.Add(1)

		go func(index int, fn func() (interface{}, error)) {
			defer wg.Done()
			data, err := fn()
			results[index] = types.Result{
				Data:  data,
				Error: err,
			}
		}(i, f)
	}

	wg.Wait()

	return results
}
