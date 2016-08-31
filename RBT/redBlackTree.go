package main

// Color of tree node
const (
	RedColor   = 1
	BlackColor = 2
)

// RBT
type RBT struct {
	Node  *node
	Color int
}

type node struct {
	Val    int
	Left   *RBT
	Right  *RBT
	Parent *RBT
}

// Insert
func (r *RBT) Insert(v int) {
	// 1. find out location

	// 2. fix up (red-black)
}

func leftRotate(p, x, y *RBT) {
	if x.Node.Right != y {
		panic("left rotate failed, x.Node.Right != y")
		return
	}
	if p.Node.Left == x {
		p.Node.Left = y
	} else if p.Node.Right == x {
		p.Node.Right = y
	} else {
		panic("left rotate failed, p is not the parent of x")
		return
	}
	y.Node.Parent = p
	x.Node.Right = y.Node.Left
	y.Node.Left.Node.Parent = x
	y.Node.Left = x
	x.Node.Parent = y
	return
}

func rightRotate(p, x, y *RBT) {
	if x.Node.Left != y {
		panic("left rotate failed, x.Node.Left != y")
		return
	}
	if p.Node.Left == x {
		p.Node.Left = y
	} else if p.Node.Right == x {
		p.Node.Right = y
	} else {
		panic("left rotate failed, p is not the parent of x")
		return
	}
	y.Node.Parent = p
	x.Node.Left = y.Node.Right
	y.Node.Right.Node.Parent = x
	y.Node.Right = x
	x.Node.Parent = y
	return
}

func fixUp(t, z *RBT) {
	for z.Node.Parent.Color == RedColor {
		if z.Node.Parent == z.Node.Parent.Node.Parent.Node.Left {
			// parent is left-child of grandparent
			y := z.Node.Parent.Node.Parent.Node.Right
			if y.Color == RedColor {
				z.Node.Parent.Color = BlackColor
				y.Color = BlackColor
				z.Node.Parent.Node.Parent.Color = RedColor
				z = z.Node.Parent.Node.Parent
			} else if z == z.Node.Parent.Node.Right {
				z = z.Node.Parent
				leftRotate(z.Node.Parent, z, z.Node.Right)
			}
			z.Node.Parent.Color = BlackColor
			z.Node.Parent.Node.Parent.Color = RedColor
			rightRotate(t, z.Node.Parent.Node.Parent, z.Node.Parent)
		} else {
			//
		}
	}
	t.Color = BlackColor
}
