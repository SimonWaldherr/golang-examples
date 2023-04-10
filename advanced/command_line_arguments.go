// Description: Command line arguments in Go
// Tags: command, line, arguments, flag, flag, flag p
package main

import (
	"flag"
	"fmt"
)

func main() {
	env := flag.String("env", "dev", "Environment(dev, qa, stg, prod)")
	cron := flag.Bool("consumer", false, "boolean")
	//Parse parses the command-line flags
	flag.Parse()
	fmt.Println("The environment set is", *env)
	fmt.Println("The consumer flag retrieved from command line is", *cron)
}
