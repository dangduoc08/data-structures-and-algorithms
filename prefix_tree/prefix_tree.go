package prefixTree

type PrefixTree struct {
	Leaf        map[string]*PrefixTree
	IsEndOfWord bool
}
