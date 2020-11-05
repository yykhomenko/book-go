package search

import "sync"

// DLS is depth-limited search
func DLS(s string, depth int, p int, seen map[string]bool, f func(string) []string) {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	sm := make(chan struct{}, p)
	if seen == nil {
		seen = make(map[string]bool)
	}

	wg.Add(1)

	var dls func(s string, depth int)
	dls = func(s string, depth int) {
		defer wg.Done()
		if depth == 0 {
			return
		}

		sm <- struct{}{}
		strs := f(s)
		<-sm

		for _, link := range strs {
			mu.Lock()
			if seen[link] {
				mu.Unlock()
				continue
			}
			seen[link] = true
			mu.Unlock()
			wg.Add(1)
			go dls(link, depth-1)
		}
	}

	dls(s, depth)
	wg.Wait()
}
