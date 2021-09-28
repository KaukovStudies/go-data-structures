package set

type HashSet map[interface{}]struct{}

func NewHashSet() *HashSet {
	return &HashSet{}
}

func (hs HashSet) Add(item interface{}) {
	if _, exists := hs[item]; !exists {
		hs[item] = struct{}{}
	}
}

func (hs HashSet) Has(item interface{}) bool {
	_, exists := hs[item]

	return exists
}

func (hs HashSet) Delete(item interface{}) {
	delete(hs, item)
}
