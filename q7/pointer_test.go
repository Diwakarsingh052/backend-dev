package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointers(t *testing.T) {
	a := pointers()
	assert.EqualValues(t, 10, a)
	assert.NotNil(t, a)
}
