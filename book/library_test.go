package book

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	aa := CurrentLibrary()
	assert.Equal(t, len(aa.books), 2, "they should be equal")
}
