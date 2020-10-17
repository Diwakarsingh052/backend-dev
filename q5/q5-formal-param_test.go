package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintSlice(t *testing.T) {
	a := []int{40, 60, 10, 2, 5}
	total := sum(a)
	assert.EqualValues(t, 117, total)
	assert.NotNil(t, total)

}
