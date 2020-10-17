package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClosure(t *testing.T) {
	series := closure()
	assert.EqualValues(t, series(), 1)
	assert.EqualValues(t, series(), 2)
	assert.NotNil(t, series())
}
