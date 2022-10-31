package main

import "fmt"

func cuboidDraw(drawX, drawY, drawZ int) {
	fmt.Printf("Cuboid %d %d %d:\n", drawX, drawY, drawZ)
	cubeLine(drawY+1, drawX, 0, "+-")
	for i := 1; i <= drawY; i++ {
		cubeLine(drawY-i+1, drawX, i-1, "/ |")
	}
	cubeLine(0, drawX, drawY, "+-|")
	for i := 4*drawZ - drawY - 2; i > 0; i-- {
		cubeLine(0, drawX, drawY, "| |")
	}
	cubeLine(0, drawX, drawY, "| +")
	for i := 1; i <= drawY; i++ {
		cubeLine(0, drawX, drawY-i, "| /")
	}
	cubeLine(0, drawX, 0, "+-\n")
}

func cubeLine(n, drawX, drawY int, cubeDraw string) {
	fmt.Printf("%*s", n+1, cubeDraw[:1])
	for d := 9*drawX - 1; d > 0; d-- {
		fmt.Print(cubeDraw[1:2])
	}
	fmt.Print(cubeDraw[:1])
	fmt.Printf("%*s\n", drawY+1, cubeDraw[2:])
}

func main() {
	fmt.Println("Enter 3 dimensions of Cuboid : ")
	var l, b, h int
	fmt.Scanln(&l)
	fmt.Scanln(&b)
	fmt.Scanln(&h)
	cuboidDraw(l, b, h)
}
