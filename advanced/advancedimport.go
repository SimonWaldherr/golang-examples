package main

import (
	// You can import a package to your own namespace, so you don't have to write the package
	// name in front of the imported function name. Use this import method wisely.
	// By wisely i mean don't use it unless you have a realy good reason for.
	. "./packages"
	// If you wan't to import a package but don't need any function of the package, you can
	// do it as
	_ "./packages/withInit"
	f "fmt"
)

func main() {
	f.Printf("use function PlusOne from \"./packages/\" without namespace: %d\n", PlusOne(2))
}
