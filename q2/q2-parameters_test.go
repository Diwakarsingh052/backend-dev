package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassByValue(t *testing.T) {
	var (
		name string
		age  int
	)
	name, age = passByValue(name, age)
	assert.EqualValues(t, "raj", name)
	assert.EqualValues(t, 21, age)
}
