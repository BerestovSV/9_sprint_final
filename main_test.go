package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testData = []struct {
	input []int
	want  int
}{
	{input: []int{1, 2, 3, 2, 1}, want: 3},
	{input: []int{-1, -2, -3, -4, 1}, want: 1},
	{input: []int{3, 3, 3, 3}, want: 3},
	{input: []int{}, want: 0},
	{input: []int{100}, want: 100},
}

var testSizes = []struct {
	inputLen int
	wantLen  int
}{
	{inputLen: 0, wantLen: 0},
	{inputLen: 1, wantLen: 1},
	{inputLen: SIZE + 1, wantLen: 0},
	{inputLen: SIZE, wantLen: SIZE},
	{inputLen: rndSize, wantLen: rndSize},
}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
var rndSize = rnd.Intn(SIZE + 1)

func TestGenerateRandomElements(t *testing.T) {
	t.Parallel()

	for _, testSize := range testSizes {
		testArr := generateRandomElements(testSize.inputLen)
		assert.Equal(t, testSize.wantLen, len(testArr), "input: %d", testSize.inputLen)

		sum := 0
		for _, v := range testArr {
			sum += v
		}

		if len(testArr) > 0 {
			require.NotZero(t, sum)
		}
	}
}

func TestMaximum(t *testing.T) {
	t.Parallel()

	for _, testArr := range testData {
		assert.Equal(t, testArr.want, maximum(testArr.input), "input: %v", testArr.input)
	}
}

func TestMaxChunks(t *testing.T) {
	t.Parallel()

	for _, testArr := range testData {
		assert.Equal(t, testArr.want, maxChunks(testArr.input), "input: %v", testArr.input)
		if len(testArr.input) > 1 {
			assert.Equal(t, maxChunks(testArr.input), maximum(testArr.input))
		}
	}
}
