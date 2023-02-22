package main

import (
	"fmt"
	"runtime"
	"sync"
)

func calc(id, price int, interestsRate float64, year int) {
	months := year * 12
	interest := 0
	for i := 0; i < months; i++ {
		balance := price * (months - i) / months
		interest += int(float64(balance) * interestsRate / 12)
	}
	fmt.Printf("year=%d total=%d interest=%d id=%d\n", year, price+interest, interest, id)
}

func worker(id, price int, interestsRate float64, years chan int, wg *sync.WaitGroup) {
	for year := range years {
		calc(id, price, interestsRate, year)
		wg.Done()
	}
}

func main() {
	price := 40000000
	interestsRate := 0.011

	years := make(chan int, 35)
	for i := 1; i < 36; i++ {
		years <- i
	}

	var wg sync.WaitGroup
	wg.Add(35)

	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, price, interestsRate, years, &wg)
	}

	close(years)
	wg.Wait()
}
