package main

import (
    "fmt"
    "sync"
)

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func countPrimes(start, end int, wg *sync.WaitGroup, mu *sync.Mutex, count *int) {
    defer wg.Done()
    localCount := 0
    for i := start; i <= end; i++ {
        if isPrime(i) {
            mu.Lock()
            *count++
            mu.Unlock()
            localCount++
        }
    }
    fmt.Printf("Goroutine from %d to %d found %d primes.\n", start, end, localCount)
}

func main() {
    const numWorkers = 5
    const maxNumber = 50

    var wg sync.WaitGroup
    var mu sync.Mutex
    count := 0

    wg.Add(numWorkers)
    for i := 0; i < numWorkers; i++ {
        start := i * (maxNumber / numWorkers)
        end := start + (maxNumber / numWorkers) - 1
        go countPrimes(start, end, &wg, &mu, &count)
    }

    wg.Wait()
    fmt.Printf("Found %d primes in total.\n", count)
}
