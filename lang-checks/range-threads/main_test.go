package rangethreads

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	const threads = 5
	var (
		wg  sync.WaitGroup
		sem int
		mu  sync.Mutex
	)
	for i := 0; i < threads; i++ {
		wg.Add(1)
		sem = threads - i - 1
		go threadfunc(i, &sem, &wg, &mu)
	}
	wg.Wait()
}

func threadfunc(i int, sem *int, wg *sync.WaitGroup, mu sync.Locker) {

	time.Sleep(100 * time.Millisecond)
	rndInt := rand.Intn(100)

	time.Sleep(time.Duration(rndInt) * time.Millisecond)
	for i != *sem {
		time.Sleep(0)
	}

	mu.Lock()
	(*sem)++
	fmt.Printf("%d ", i)
	mu.Unlock()

	wg.Done()
}
