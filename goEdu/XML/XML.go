package main

import (
	"encoding/xml"
	"fmt"
)

//type Human struct {
//	XMLName  xml.Name `xml:"human"`
//	Name     string   `xml:"name,attr"`
//	Age      uint8    `xml:"age,attr"`
//	Features []string `xml:"Features"`
//}

type Human struct {
	XMLName  xml.Name `xml:"human"`
	Name     string   `xml:"name,attr"`
	Age      uint8    `xml:"age,attr"`
	Features []string `xml:"features>feature"` // injection each of <feature></feature> in <features></features>
}

func main() {
	human1 := Human{Name: "Aleksandr", Age: 21, Features: []string{"JAVA", "GO"}}
	human1XML, _ := xml.MarshalIndent(human1, " ", " ")
	human1XMLWithHeader := xml.Header + string(human1XML)
	fmt.Printf("human1XMLWithHeader:\n%s\n", human1XMLWithHeader)

	// DECODE XML:
	var human2 Human
	if ok := xml.Unmarshal(human1XML, &human2); ok != nil {
		panic(ok.Error())
	}
	fmt.Println("human2:", human2)

}
