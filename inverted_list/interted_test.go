package inverted_list_test

import (
	"testing"

	"github.com/kauebonfimm/go-algorithms/inverted_list"
	"github.com/stretchr/testify/assert"
)

func TestIntertedList(t *testing.T) {

	assert := assert.New(t)
	arr := inverted_list.Invert([]inverted_list.Users{
		{Id: 1}, {Id: 2}, {Id: 3}, {Id: 4}, {Id: 5}, {Id: 6}})
	assert.Equal(arr, []inverted_list.Users{
		{Id: 6}, {Id: 5}, {Id: 4}, {Id: 3}, {Id: 2}, {Id: 1}})

}
