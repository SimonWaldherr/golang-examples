package foobar

func privateFunc() int { //only functions with a capitalized name are usable outside the package
	return 23
}

func PublicFunc(value int) int { //this is a exported function
	return value + privateFunc()
}

func PlusOne(value int) int { //and another one
	return value + 1
}
