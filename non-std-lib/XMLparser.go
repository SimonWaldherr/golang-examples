package main

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/antchfx/xpath"
	"strings"
)

func main() {
	s := `<?xml version="1.0" encoding="UTF-8" ?>
<breakfast_menu>
	<food>
		<name price="10">Berry-Berry Belgian Waffles</name>
		<description>Light Belgian waffles</description>
		<calories>900</calories>
	</food>
	<food>
		<name price="20">French Toast</name>
		<description>Thick slices</description>
		<calories>600</calories>
	</food>
	<food>
		<name price="30">Homestyle Breakfast</name>
		<description>Two eggs, bacon or sausage</description>
		<calories>950</calories>
	</food>	
</breakfast_menu>`

	doc, err := xmlquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}

	root := xmlquery.FindOne(doc, "//breakfast_menu")
	if n := root.SelectElement("//food/name"); n != nil {
		fmt.Printf("Name #%s\n", n.InnerText())
	}

	if n := root.SelectElement("//food[2]/name"); n != nil {
		fmt.Printf("Name #%s\n", n.InnerText())
	}

	for i, n := range xmlquery.Find(doc, "//food/name/@price") {
		fmt.Printf("Price #%d %s\n", i, n.InnerText())
	}

	for i, n := range xmlquery.Find(doc, "//food/calories") {
		fmt.Printf("Calories #%d %s\n", i, n.InnerText())
	}

	if n := root.SelectElement("//food[2]/name"); n != nil {
		fmt.Printf("Attr #%s\n", n.Attr)
	}

	if n := root.SelectElement("//food[2]/name"); n != nil {
		fmt.Printf("Data #%s\n", n.Data)
	}

	node := xmlquery.FindOne(doc, "//breakfast_menu/food[2]")
	if n := node.SelectElement("//description"); n != nil {
		fmt.Printf("Description #%s\n", n.InnerText())
	}

	expr, err := xpath.Compile("sum(//breakfast_menu/food/name/@price)")
	price := expr.Evaluate(xmlquery.CreateXPathNavigator(doc)).(float64)

	fmt.Printf("Total price: %f\n", price)

	countexpr, err := xpath.Compile("count(//breakfast_menu/food)")
	count := countexpr.Evaluate(xmlquery.CreateXPathNavigator(doc))

	fmt.Printf("Food Node Counts: %f\n", count)
}
