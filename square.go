package gopkg

import "fmt"

// Square is one of the shapes
type Square struct {
	Width   int
	Height  int
	Breadth int
}

// Description for Square
func (r Square) Description() {
	fmt.Printf("This is a Square. w=%v, h=%v\n", r.Width, r.Height)
}

// Area of a rectange
func (r Square) Area() int {
	return r.Width * r.Height
}

// Volume for a Square
func (r Square) Volume() int {
	return r.Width * r.Height * r.Breadth
}
