package golearn

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAbsolute(t *testing.T) {

	t.Run("negative test case", func(t *testing.T) {

		if testing.Short() {
			t.Skip()
		}

		<-time.After(5 * time.Second)

		c := Absolute(-5)
		assert.Equal(t, 5, c)
	})

	t.Run("positive test case", func(t *testing.T) {
		c := Absolute(5)
		assert.Equal(t, 5, c)
	})

}

func TestAdd_WithTestTable(t *testing.T) {

	testCases := []struct {
		name           string
		a, b, expected int
	}{
		{
			name:     "negative and negative",
			a:        -1,
			b:        -1,
			expected: -2,
		},
		{
			name:     "negative and positive",
			a:        -1,
			b:        1,
			expected: 0,
		},
		{
			name:     "positive and positive",
			a:        1,
			b:        1,
			expected: 2,
		},
		{
			name:     "positive and negative",
			a:        1,
			b:        -1,
			expected: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Add(tc.a, tc.b)
			assert.Equal(t, tc.expected, c)
		})
	}
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
