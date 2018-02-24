package main

import (
	"fmt"
)

// BST binary search tree
type BST struct {
	Val   *int
	Left  *BST
	Right *BST
}

// Find
func (b *BST) Find(v int) *BST {
	if *b.Val == v {
		return b
	}
	if *b.Val < v {
		if b.Right == nil {
			return nil
		}
		return b.Right.Find(v)
	}
	if *b.Val > v {
		if b.Left == nil {
			return nil
		}
		return b.Left.Find(v)
	}
	return nil
}

// Insert
func (b *BST) Insert(v int) {
	var t *BST = b
	for {
		// existed
		if *t.Val == v {
			return
		}
		if *t.Val < v {
			if t.Right == nil {
				t.Right = &BST{Val: &v}
				return
			}
			t = t.Right
		} else {
			if t.Left == nil {
				t.Left = &BST{Val: &v}
				return
			}
			t = t.Left
		}
	}
}

// Delete
func (b *BST) Delete(v int) {
	var t *BST = b
	var p *BST
	for t != nil {
		// find out and delete
		if *t.Val == v {
			if p == nil {

				// root
				// 1.no children
				if b.Left == nil && b.Right == nil {
					b.Val = nil
					return
				}
				// 2.only left child
				if b.Right == nil {
					var (
						_p   *BST = b
						rmid *BST = b.Left
					)
					for rmid.Right != nil {
						_p = rmid
						rmid = rmid.Right
					}
					b.Val = POI(*rmid.Val)
					// delete duplicate
					if _p != b {
						_p.Right = nil
					} else {
						_p.Left = nil
					}
					return
				}
				// 3.only right child or double children
				var (
					_p  *BST = b
					mid *BST = b.Right
				)
				for mid.Left != nil {
					_p = mid
					mid = mid.Left
				}
				b.Val = POI(*mid.Val)
				// delete duplicate
				if _p != b {
					_p.Left = nil
				} else {
					_p.Right = nil
				}
				return

			} else {

				// child
				if t == p.Right {
					// right child
					if t.Left == nil {
						p.Right = t.Right
					} else if t.Right == nil {
						p.Right = t.Left
					} else {
						// double children
						var (
							_p  *BST = t
							mid *BST = t.Right
						)
						for mid.Left != nil {
							_p = mid
							mid = mid.Left
						}
						t.Val = POI(*mid.Val)
						// delete duplicate
						if _p != t {
							_p.Left = nil
						} else {
							_p.Right = nil
						}
					}
				} else {
					// left child
					if t.Left == nil {
						p.Left = t.Right
					} else if t.Right == nil {
						p.Left = t.Left
					} else {
						// double children
						var (
							_p  *BST = t
							mid *BST = t.Left
						)
						for mid.Right != nil {
							_p = t
							t = t.Right
						}
						t.Val = POI(*mid.Val)
						// delete duplicate
						if _p != t {
							_p.Right = nil
						} else {
							_p.Left = nil
						}
					}
				}
			}
		}
		// compare and loop
		p = t
		if *t.Val < v {
			t = t.Right
		} else {
			t = t.Left
		}
	}
}

func (b *BST) String() string {
	var (
		left  string = "nil"
		rgiht string = "nil"
	)
	if b.Left != nil {
		left = b.Left.String()
	}
	if b.Right != nil {
		rgiht = b.Right.String()
	}
	return fmt.Sprintf("%d -< [%s, %s]", *b.Val, left, rgiht)
}
