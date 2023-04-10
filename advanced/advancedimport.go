// Description: Advanced import examples
// Tags: import, package, namespace, alias, dot, underscore
package main

import (
	// You can import a package to your own namespace, so you don't have to write the package
	// name in front of the imported function name. Use this import method wisely.
	// By wisely i mean don't use it unless you have a really good reason for.
	. "./packages"
	// If you wan't to import a package but don't need any function of the package, you can
	// do it like this (you only need it very rarely, e.g. for database "drivers"
	// or for image format "drivers"):
	_ "./packages/withInit"
	// If you wan't to import a package with a different name, you can write the alias and then
	// the package name like this:
	f "fmt"
)

func main() {
	f.Printf("use function PlusOne from \"./packages/\" without namespace: %d\n", PlusOne(2))
}
