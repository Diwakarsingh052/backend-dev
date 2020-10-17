package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMap(t *testing.T) {
	eng := map[string]string{
		"up":   "above",
		"down": "below",
	}
	out := CreateMap()
	assert.NotNil(t, out)
	assert.EqualValues(t, eng, out)
}
