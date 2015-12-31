package pulitzer

import (
	"../xtraml/choice"
	"encoding/xml"
	"github.com/golang-commonmark/markdown"
)

type TokenList []markdown.Token

func (t TokenList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return choice.WrapList(e, xml.Name{Local: "tokens"}, t)
}

func Convert(tokens TokenList) NodeList {
	return NodeList(convertAll(tokens...))
}
