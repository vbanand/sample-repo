package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplayHello(t *testing.T) {
	str := displayHello()
	assert.Equal(t, str, "Hello World", "Expected Hello World")
}
