package shape

import "fmt"

// Rectangle is one of the shapes
type Rectangle struct {
	Width   int
	Height  int
	Breadth int
}

// Description for Rectangle
func (r Rectangle) Description() {
	fmt.Printf("This is a rectangle. w=%v, h=%v\n", r.Width, r.Height)
}

// Area of a rectange
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// Volume for a rectangle
func (r Rectangle) Volume() int {
	return r.Width * r.Height * r.Breadth
}
