package ptree

import "testing"

func BenchmarkPTree_Get(b *testing.B) {
	values := map[string]string{
		"Hello":         "World",
		"Hallo":         "deutsch",
		"DifferentSize": "what's then?",
		"SomeVeryVeryVeryVeryLooooooongKeyThatIsAlsoPossibleInRealProduction": "a",
	}
	tree := NewTree[string]()
	tree.InsertMap(values)

	b.Run("Hello", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = tree.Get("Hello")
		}
	})

	b.Run("DifferentSize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = tree.Get("DifferentSize")
		}
	})

	b.Run("Long", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = tree.Get("SomeVeryVeryVeryVeryLooooooongKeyThatIsAlsoPossibleInRealProduction")
		}
	})
}
