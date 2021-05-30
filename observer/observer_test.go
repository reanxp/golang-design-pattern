package observer

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestObserver(t *testing.T) {
	convey.Convey("Observer", t, func() {
		convey.Convey("Success", func() {
			obj1 := NewConcreateObserver("obj1")
			obj2 := NewConcreateObserver("obj2")

			subject := NewConcreateSubject()
			subject.Attach(obj1)
			subject.Attach(obj2)
			subject.SetState(2)
		})
	})
}
