package main

import (
	"fmt"
	"regexp"
)

/*
Date and time expressed according to ISO 8601:

Date:							2014-11-01
Combined date and time in UTC:	2014-11-01T10:01:13+00:00
								2014-11-01T10:01:13Z
Week:							2014-W44
Date with week number:			2014-W44-6
Ordinal date:					2014-305
*/

var input string
var r = regexp.MustCompile("(?P<date>\\d{4}(\\-W\\d{1,2}(\\-\\d)?|\\-\\d{2}\\-\\d{2}(T\\d{2}:\\d{2}(:\\d{2}(\\+\\d{2}:\\d{2}Z?)?)?)?|\\-\\d{1,3}))")

func main() {
	input = `
2014-06-06
2014-11-01
2014-11-01T10:01:13+00:00
2014-11-01T10:01:13Z
2014-W44
2014-W44-6
2014-305
2014-23
`
	date := r.FindAllStringSubmatch(input, -1)
	for i := 0; i < len(date); i++ {
		fmt.Printf("ISO8601-Date: %#v\n", date[i])
	}
}
