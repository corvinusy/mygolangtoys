package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type table struct {
	forks []sync.Mutex
}

type philosopher struct {
	name  string
	left  int
	right int
}

func (p *philosopher) eat(t *table) {
	defer wg.Done()
	t.forks[p.left].Lock()
	defer t.forks[p.left].Unlock()
	t.forks[p.right].Lock()
	defer t.forks[p.right].Unlock()

	fmt.Println(p.name, "is eating.")

	time.Sleep(1 * time.Second)

	fmt.Println(p.name, "finished eating.")
}

func main() {
	philosophers := [...]philosopher{
		philosopher{"Джудит Батлер", 0, 1},
		philosopher{"Рая Дунаевская", 1, 2},
		philosopher{"Наталья Зарубина", 2, 3},
		philosopher{"Эмма Гольдман", 3, 4},
		philosopher{"Анна Шмидт", 0, 4},
	}
	t := table{forks: make([]sync.Mutex, len(philosophers))}

	for i := range philosophers {
		wg.Add(1)
		go philosophers[i].eat(&t)
	}
	wg.Wait()
}
