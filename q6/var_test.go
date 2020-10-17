package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVariable(t *testing.T) {
	age, name := variables()
	assert.NotNil(t, age)
	assert.NotNil(t, name)

}
