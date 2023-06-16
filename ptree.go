package ptree

type PTree[T any] struct {
	transitions [][]int
	values      []T
}

func NewTree[T any]() *PTree[T] {
	return &PTree[T]{
		transitions: [][]int{make([]int, 256)},
	}
}

func (p *PTree[T]) InsertMap(m map[string]T) {
	for key, value := range m {
		p.Insert(key, value)
	}
}

func (p *PTree[T]) Insert(key string, value T) {
	transTable := 0

	for i := 0; i < len(key)-1; i++ {
		next := p.transitions[transTable][key[i]]

		if next == 0 {
			p.transitions = append(p.transitions, make([]int, 256))
			p.transitions[transTable][key[i]] = len(p.transitions) - 1
			next = len(p.transitions) - 1
		}

		transTable = next
	}

	p.values = append(p.values, value)
	p.transitions[transTable][key[len(key)-1]] = -len(p.values)
}

func (p *PTree[T]) Get(key string) (value T, found bool) {
	var (
		current, next int
	)

	trans := p.transitions

	for i := 0; i < len(key)-1; i++ {
		next = trans[current][key[i]]
		if next <= 0 {
			return value, false
		}

		current = next
	}

	valueIndex := trans[current][key[len(key)-1]]
	if valueIndex >= 0 {
		return value, false
	}

	return p.values[-valueIndex-1], true
}
