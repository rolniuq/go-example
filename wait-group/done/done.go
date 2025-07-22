package done

import "sync"

type wgDone struct {
	wg *sync.WaitGroup
}

func (w *wgDone) Go(worker func()) {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		worker()
	}()
}
