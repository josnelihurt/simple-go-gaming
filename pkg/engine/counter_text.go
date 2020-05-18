package engine

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

//CounterText represents component for counters
type CounterText struct {
	parent               *Element
	textRenderer         *TextRenderer
	currentValue         int
	initialValue         int
	incrementStep        int
	finalValue           int
	textFormat           string
	incrementCondition   func(message *Message, parent *Element) bool
	resetCondition       func(message *Message, parent *Element) bool
	onFinalValueDetected func(parent *Element)
}

//NewCounterText creates a CounterText instance
func NewCounterText(
	parent *Element, textRenderer *TextRenderer,
	initialValue, incrementStep, finalValue int,
	textFormat string,
	incrementCondition func(message *Message, parent *Element) bool,
	resetCondition func(message *Message, parent *Element) bool,
	onFinalValueDetected func(parent *Element),
) *CounterText {
	return &CounterText{
		parent:               parent,
		currentValue:         initialValue,
		initialValue:         initialValue,
		textRenderer:         textRenderer,
		incrementStep:        incrementStep,
		incrementCondition:   incrementCondition,
		textFormat:           textFormat,
		onFinalValueDetected: onFinalValueDetected,
		resetCondition:       resetCondition,
	}
}

func (context *CounterText) OnUpdate() error {
	if context.onFinalValueDetected != nil && context.currentValue == context.finalValue {
		context.onFinalValueDetected(context.parent)
	}
	context.textRenderer.SetNewText(fmt.Sprintf(context.textFormat, context.currentValue))
	return nil
}
func (context *CounterText) OnMessage(message *Message) error {
	if context.incrementCondition != nil && context.incrementCondition(message, context.parent) {
		context.currentValue += context.incrementStep
	}
	if context.resetCondition != nil && context.resetCondition(message, context.parent) {
		context.currentValue = context.initialValue
	}
	return nil
}

func (context *CounterText) OnDraw(renderer *sdl.Renderer) error { return nil }
func (context *CounterText) OnCollision(other *Element) error    { return nil }
