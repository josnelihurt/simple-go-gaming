package engine

type Vector struct {
	X, Y float64
}
type Circle struct {
	Center *Vector
	Radius float64
}
type Message struct {
	Sender              *Element
	RelatedTo           []*Element
	Code                int
	Data                string
	SendToOtherElements bool
}
