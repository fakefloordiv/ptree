package ptree

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPTree(t *testing.T) {
	t.Run("compare to map values", func(t *testing.T) {
		values := map[string]string{
			"Hello":         "World",
			"Hallo":         "deutsch",
			"DifferentSize": "what's then?",
			"SomeVeryVeryVeryVeryLooooooongKeyThatIsAlsoPossibleInRealProduction": "a",
		}
		tree := NewTree[string]()
		tree.InsertMap(values)

		for key, value := range values {
			actual, found := tree.Get(key)
			if assert.True(t, found) {
				assert.Equal(t, value, actual)
			}
		}
	})

	// this test will anyway fail, as common prefix collisions cannot be easily
	// solved by the algorithm. But I expected these values to be at least inserted:(
	t.Run("common prefix collision", func(t *testing.T) {
		tree := NewTree[string]()
		tree.Insert("Hell", "heaven")
		tree.Insert("Hello", "world")
		actual, found := tree.Get("Hell")
		require.True(t, found, "key is presented, but not found")
		require.Equal(t, "heaven", actual)
		actual, found = tree.Get("Hello")
		require.True(t, found, "key is presented, but not found")
		require.Equal(t, "world", actual)
	})
}
