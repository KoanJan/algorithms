package BH

// BinomialTree binomial tree
type BinomialTree struct {
	Child   *BinomialTree // left
	Sibling *BinomialTree // right
	Parent  *BinomialTree // parent
	Head    *BinomialTree // head

	Degree int // the number of children

	key int // key
}
