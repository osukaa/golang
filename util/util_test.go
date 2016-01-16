package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloworld(t *testing.T) {
	message := Helloworld()
	assert.Equal(t, "Hello to the world", message, "Deberian ser iguales")
}
