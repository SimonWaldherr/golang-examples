package main

import "fmt"

type exampledata struct {
	a, b int
	c, d string
	e, f bool
	g, h float64
}

func printx(str string, data exampledata) {
	fmt.Printf("%"+str+":\t "+str+"\n", data)
}

func main() {
	x := exampledata{0, 1, "a", "b", true, false, 3.1415926535, 2.357111317192329}

	printx("%v", x)
	//%v:	 {0 1 a b true false 3.1415926535 2.357111317192329}

	printx("%+v", x)
	//%+v:	 {a:0 b:1 c:a d:b e:true f:false g:3.1415926535 h:2.357111317192329}

	printx("%#v", x)
	//%#v:	 main.exampledata{a:0, b:1, c:"a", d:"b", e:true, f:false, g:3.1415926535, h:2.357111317192329}

	printx("%t", x)
	//%t:	 {%!t(int=0) %!t(int=1) %!t(string=a) %!t(string=b) true false %!t(float64=3.1415926535) %!t(float64=2.357111317192329)}

	printx("%T", x)
	//%T:	 main.exampledata

	printx("%d", x)
	//%d:	 {0 1 %!d(string=a) %!d(string=b) %!d(bool=true) %!d(bool=false) %!d(float64=3.1415926535) %!d(float64=2.357111317192329)}

	printx("%b", x)
	//%b:	 {0 1 %!b(string=a) %!b(string=b) %!b(bool=true) %!b(bool=false) 7074237751826244p-51 5307742824889076p-51}

	printx("%c", x)
	//%c:	 {  %!c(string=a) %!c(string=b) %!c(bool=true) %!c(bool=false) %!c(float64=3.1415926535) %!c(float64=2.357111317192329)}

	printx("%x", x)
	//%x:	 {0 1 61 62 %!x(bool=true) %!x(bool=false) 0x1.921fb54411744p+01 0x1.2db5d2da2faf4p+01}

	printx("%f", x)
	//%f:	 {%!f(int=0) %!f(int=1) %!f(string=a) %!f(string=b) %!f(bool=true) %!f(bool=false) 3.141593 2.357111}

	printx("%e", x)
	//%e:	 {%!e(int=0) %!e(int=1) %!e(string=a) %!e(string=b) %!e(bool=true) %!e(bool=false) 3.141593e+00 2.357111e+00}

	printx("%E", x)
	//%E:	 {%!E(int=0) %!E(int=1) %!E(string=a) %!E(string=b) %!E(bool=true) %!E(bool=false) 3.141593E+00 2.357111E+00}

	printx("%s", x)
	//%s:	 {%!s(int=0) %!s(int=1) a b %!s(bool=true) %!s(bool=false) %!s(float64=3.1415926535) %!s(float64=2.357111317192329)}

	printx("%q", x)
	//%s:	 {%!s(int=0) %!s(int=1) a b %!s(bool=true) %!s(bool=false) %!s(float64=3.1415926535) %!s(float64=2.357111317192329)}

	printx("%p", x)
	//%p:	 %!p(main.exampledata={0 1 a b true false 3.1415926535 2.357111317192329})

	fmt.Println("\n\n", fmt.Sprintf("%+v", x))
	// {a:0 b:1 c:a d:b e:true f:false g:3.1415926535 h:2.357111317192329}
}
