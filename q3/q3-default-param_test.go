package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassByValue(t *testing.T) {
	var u user
	u = defaultParam(u)
	assert.EqualValues(t, "Ajay", u.name)
	assert.EqualValues(t, 29, u.age)
}
