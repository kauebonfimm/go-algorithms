package two_numbers_test

import (
	"slices"
	"testing"

	"github.com/kauebonfimm/go-algorithms/two_numbers"
)

func TestTwoSums(t *testing.T) {
	indexes := two_numbers.TwoSum([]int{10, 55, 28, 78, 90}, 100)
	if !(slices.Contains(indexes, 0) && slices.Contains(indexes, 4)) {
		t.Error("invalid test two sum")
	}
}
