package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloatToInt(t *testing.T) {
	a := FloatToInt(54.5)
	assert.EqualValues(t, 54, a)
}
func TestIntToFloat(t *testing.T) {
	a := IntToFloat(58)
	assert.EqualValues(t, 58, a)
}
func TestIntToString(t *testing.T) {
	a := IntToString(67)
	assert.EqualValues(t, "67", a)
}
func TestStringToInt(t *testing.T) {
	a := StringToInt("21")
	assert.EqualValues(t, 21, a)
}
