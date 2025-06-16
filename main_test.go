package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	t.Parallel()

	emptyArr := generateRandomElements(0)
	shortArr := generateRandomElements(1)
	longArr := generateRandomElements(SIZE + 1)
	maxArr := generateRandomElements(SIZE)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rndSize := rnd.Int()
	rndArr := generateRandomElements(rndSize)

	switch {
	case rndSize > SIZE:
		assert.Equal(t, len(rndArr), 0)
	default:
		assert.Equal(t, len(rndArr), rndSize)
	}

	assert.Equal(t, len(emptyArr), 0)
	assert.Equal(t, len(shortArr), 0)
	assert.Equal(t, len(longArr), 0)
	assert.Equal(t, len(maxArr), SIZE)

	sum := 0
	for _, v := range maxArr {
		sum += v
	}
	require.NotZero(t, sum)
}

func TestMaximum(t *testing.T) {
	t.Parallel()

	var testArr = []int{1, 2, 3, 2, 1}
	assert.Equal(t, maximum(testArr), 3, "%d", testArr)

	testArr = []int{-1, -2, -3, -4, 1}
	assert.Equal(t, maximum(testArr), 1)

	testArr = []int{3, 3, 3, 3}
	assert.Equal(t, maximum(testArr), 3)

	testArr = []int{}
	assert.Equal(t, maximum(testArr), 0)

	testArr = []int{100}
	assert.Equal(t, maximum(testArr), 100)
}

func TestMaxChunks(t *testing.T) {
	t.Parallel()

	var testArr = []int{1, 2, 3, 2, 1}
	assert.Equal(t, maxChunks(testArr), 3)
	assert.Equal(t, maxChunks(testArr), maxChunks(testArr))

	testArr = []int{-1, -2, -3, -4, 1}
	assert.Equal(t, maxChunks(testArr), 1)
	assert.Equal(t, maxChunks(testArr), maxChunks(testArr))

	testArr = []int{3, 3, 3, 3}
	assert.Equal(t, maxChunks(testArr), 3)
	assert.Equal(t, maxChunks(testArr), maxChunks(testArr))

	testArr = []int{}
	assert.Equal(t, maxChunks(testArr), 0)
	assert.Equal(t, maxChunks(testArr), maxChunks(testArr))

	testArr = []int{100}
	assert.Equal(t, maxChunks(testArr), 100)
	assert.Equal(t, maxChunks(testArr), maxChunks(testArr))
}
