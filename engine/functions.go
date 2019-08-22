package engine

import (
	"fmt"

	"github.com/josnelihurt/simple-go-gaming/engine/util"
)

func isSingleAndEmpty(a []string) bool {
	return len(a) == 1 && a[0] == ""
}
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// containsInt missing templates I wonder how can I do it
func containsInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// BindMessage connects a message-channel from one element to other
func BindMessage(source *Element, destination *Element) {
	source.RegisterEmmiterCallback(func(message *Message) error {
		util.Logger <- fmt.Sprintf("from:%v to:%v msg:%v", source.Tag, destination.Tag, message)
		destination.BroadcastMessageToComponents(message)
		return nil
	})
}

//BindMessages connects a message-channel from many elements to another element
func BindMessages(sources []*Element, destination *Element) {
	for _, source := range sources {
		BindMessage(source, destination)
	}
}
