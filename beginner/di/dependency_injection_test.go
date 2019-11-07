package di

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrintCheckList(t *testing.T) {
	Convey("Given a checklist", t, func() {
		buffer := &bytes.Buffer{}
		checklist := []string{
			"[✓] Get milk",
			"[✓] Learn Go",
			"[✗] Book holidays",
		}

		PrintCheckList(buffer, checklist)

		got := buffer.String()
		want := `[✓] Get milk
[✓] Learn Go
[✗] Book holidays
`
		Convey("expect list items to be printed", func() {
			So(got, ShouldEqual, want)
		})
	})
}
