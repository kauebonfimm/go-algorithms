package bubble_sorted_test

import (
	"testing"

	"github.com/kauebonfimm/go-algorithms/bubble_sorted"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSorted(t *testing.T) {
	assert := assert.New(t)

	input := []int{35, 14, 22, 48, 1}

	input = bubble_sorted.BubbleSorted(input)

	assert.Equal([]int{1, 14, 22, 35, 48}, input)
}
