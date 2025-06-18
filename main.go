package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {

	if size <= 0 || size > SIZE {
		return nil // panic("wrong slice size in generateRandomElements")
	}

	arrayOfRand := make([]int, size)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		arrayOfRand[i] = rnd.Int()
	}

	return arrayOfRand

}

// maximum returns the maximum number of elements.
func maximum(data []int) int {

	if len(data) <= 0 || len(data) > SIZE {
		return 0 // panic("wrong slice size in maximum")
	}

	max := data[0]

	for _, value := range data {
		if value > max {
			max = value
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {

	if len(data) <= 0 || len(data) > SIZE {
		return 0 // panic("wrong slice size in maxChunks")
	}

	chunkSize := len(data) / CHUNKS
	maxInChanks := make([]int, CHUNKS)

	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			maxInChanks[i] = maximum(data[start:end])
		}(start, end)
	}

	wg.Wait()

	finalMax := maximum(maxInChanks)

	return finalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	array := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(array)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(array)
	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())

}
