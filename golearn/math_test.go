package golearn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsolute(t *testing.T) {

	t.Run("negative test case", func(t *testing.T) {
		c := Absolute(-5)
		assert.Equal(t, 5, c)
	})

	t.Run("positive test case", func(t *testing.T) {
		c := Absolute(5)
		assert.Equal(t, 5, c)
	})

}

func TestAdd(t *testing.T) {

	t.Run("negative and negative", func(t *testing.T) {
		c := Add(-1, -2)
		assert.Equal(t, -3, c)
	})

	t.Run("positive and negative", func(t *testing.T) {
		c := Add(1, -2)
		assert.Equal(t, -1, c)
	})

	t.Run("positive and positive", func(t *testing.T) {
		c := Add(1, 2)
		assert.Equal(t, 3, c)
	})

	t.Run("negative and positive", func(t *testing.T) {
		c := Add(-1, 2)
		assert.Equal(t, 1, c)
	})

}
