package engine

type Vector struct {
	X, Y float64
}
type Circle struct {
	Center *Vector
	Radius float64
}
type Message struct {
	Code int
	Data string
}
