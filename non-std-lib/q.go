package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ryboe/q"
)

func main() {
	// q.Q writes its output to a temp file, you should run "tail -f $TMPDIR/q" in a separate terminal/cli-window
	var strings []string
	var jsonstring = `["lorem", "ipsum", "dolor", "sit", "amet"]`

	// convert json bytes-string to object
	err := json.Unmarshal([]byte(jsonstring), &strings)

	if err != nil {
		fmt.Println("error while unmarshalling")
		os.Exit(2)
	}

	// debug strings variable with q.Q
	q.Q(strings)

	//convert object to bytes-string
	func(str []string) {
		jsonData, err := json.Marshal(str)

		if err != nil {
			fmt.Println("error while marshalling")
			os.Exit(2)
		}

		// debug jsonData variable in an anonymous function with q.Q
		q.Q(string(jsonData))
	}(strings)

	/* you will see something like this in your other terminal/cli-window:
	 *
	 * [18:00:00 non-std-lib/q.go:25 main.main]
	 * 0.000s strings=[]string{"lorem", "ipsum", "dolor", "sit", "amet"}
	 *
	 * [18:00:00 non-std-lib/q.go:37 main.main.func1]
	 * 0.000s string(jsonData)=["lorem","ipsum","dolor","sit","amet"]
	 *
	 * it contains infos about the folder, file and function it was executed in,
	 * the runtime and the data passed to q.Q
	 *
	 */
}
