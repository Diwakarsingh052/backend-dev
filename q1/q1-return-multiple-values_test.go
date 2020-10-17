package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultipleValues(t *testing.T) {
	msg, sum := multipleValues()
	assert.EqualValues(t, msg, "Total")
	assert.NotNil(t, sum)
}
