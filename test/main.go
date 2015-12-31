package main

import (
	"../"
	"encoding/xml"
	"github.com/golang-commonmark/markdown"
	"io/ioutil"
	"os"
)

func main() {
	bytes, err := ioutil.ReadFile("../../food/monster/docConvert/instructions.md")
	if err != nil {
		panic(err)
	}
	md := markdown.New(markdown.HTML(true), markdown.XHTMLOutput(true), markdown.Typographer(false))
	tokens := md.Parse(bytes)
	tokenList := pulitzer.TokenList(tokens)
	nodes := pulitzer.Convert(tokenList)
	bytes, err = xml.MarshalIndent(nodes, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("tokens.xml", bytes, os.ModePerm)
	if err != nil {
		panic(err)
	}
	html := md.RenderTokensToString(tokens)
	err = ioutil.WriteFile("result.html", []byte(html), os.ModePerm)
	if err != nil {
		panic(err)
	}

}
