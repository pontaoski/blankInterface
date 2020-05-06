package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"regexp"

	"github.com/alecthomas/participle"
)

func main() {
	parser := participle.MustBuild(&ProtocolDescription{})
	protocol := ProtocolDescription{}
	parser.Parse(os.Stdin, &protocol)
	out, _ := xml.MarshalIndent(protocol.ToProtocol(), "", "\t")
	r := regexp.MustCompile(`></(arg|entry)>`)
	out = r.ReplaceAll(out, []byte(" />"))
	fmt.Println(string(out))
}
