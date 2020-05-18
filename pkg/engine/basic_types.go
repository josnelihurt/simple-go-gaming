package engine

// Vector is the repesentation for a point x-y
type Vector struct {
	X, Y float64
}

// Rect is thre representation for a rectangle with x0,y0 as initial point and Width and Height
type Rect struct {
	X, Y, Width, Height float64
}

// Circle is a representation of a circle with a center and its radius
type Circle struct {
	Center *Vector
	Radius float64
}

// Message holds the data in the message passing
type Message struct {
	Sender              *Element
	RelatedTo           []*Element
	Code                int
	Data                string
	SendToOtherElements bool
}
