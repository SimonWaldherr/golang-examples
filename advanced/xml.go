package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func (c *classAccessesMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	c.Map = map[string]string{}
	val := ""

	for {
		t, _ := d.Token()
		switch tt := t.(type) {

		case xml.StartElement:

		case xml.EndElement:
			if tt.Name == start.Name {
				return nil
			}
			c.Map[tt.Name.Local] = val
		default:
			val = strings.TrimSpace(fmt.Sprintf("%s", tt))
		}
	}
}

func ParseDocumentXML(xmlString string) Variables {
	m := Variables{}
	xml.Unmarshal([]byte(xmlString), &m)
	return m
}

type Variables struct {
	Head Head
	Data classAccessesMap `xml:"Data"`
}

type Head struct {
	Name string
	IP   string
}

type classAccessesMap struct {
	Map map[string]string
}

func main() {
	txt := `
        <Variables>
       <Head>
           <Name>WKS001</Name>
           <IP>192.168.178.23</IP>
       </Head>
       <Data>
           <POSNR>10</POSNR>
           <AUFNR>103002</AUFNR>
           <KUNNR>332401</KUNNR>
           <ZTXT1>lorem</ZTXT1>
           <ZTXT2>ipsum</ZTXT2>
           <ZTXT3>dolor</ZTXT3>
           <ZTXT4>sit a</ZTXT4>
       </Data>
    </Variables>
        `
	m := Variables{}
	xml.Unmarshal([]byte(txt), &m)

	fmt.Printf("xml: %#v\n", m)
	for a, b := range m.Data.Map {
		fmt.Printf("key: %s, value: %s\n", a, b)
	}
}
