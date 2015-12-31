package pulitzer

import (
	"../xtraml/choice"
	"encoding/xml"
	"fmt"
	"github.com/golang-commonmark/markdown"
	"reflect"
)

type Node interface {
	String() string
}

type NodeList []Node

func (n NodeList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return choice.WrapList(e, xml.Name{Local: "body"}, n)
}

type EnclosureConfig interface {
	CfgString() string
}

type Open struct {
	XMLName xml.Name        `xml:"open"`
	Config  EnclosureConfig `xml:",any"`
	Type    EnclosureType   `xml:"type,attr"`
}

func (n Open) String() string {
	return fmt.Sprintf("%v", n)
}

type Close struct {
	XMLName xml.Name      `xml:"close"`
	Type    EnclosureType `xml:"type,attr"`
}

func (n Close) String() string {
	return fmt.Sprintf("%v", n)
}

type Content struct {
	XMLName xml.Name    `xml:"content"`
	Type    ContentType `xml:"type,attr"`
	Content string      `xml:"text,attr"`
}

func (n Content) String() string {
	return fmt.Sprintf("%v", n)
}

type Break struct {
	XMLName xml.Name  `xml:"break"`
	Type    BreakType `xml:"text,attr"`
}

func (n Break) String() string {
	return fmt.Sprintf("%v", n)
}

type Fence struct {
	XMLName xml.Name `xml:"fence"`
	Params  string   `xml:"params,attr"`
	Content string   `xml:"text,attr"`
}

func (n Fence) String() string {
	return fmt.Sprintf("%v", n)
}

type Image struct {
	XMLName  xml.Name `xml:"image"`
	Src      string   `xml:"src,attr"`
	Title    string   `xml:"text,attr"`
	Children []Node   `xml:",any"`
}

func (n Image) String() string {
	return fmt.Sprintf("%v", n)
}

type Inline struct {
	XMLName  xml.Name `xml:"inline"`
	Content  string   `xml:"content,attr"`
	Children []Node   `xml:",any"`
}

func (n Inline) String() string {
	return fmt.Sprintf("%v", n)
}

type HeadingConfig struct {
	XMLName xml.Name `xml:"heading-config"`
	HLevel  int      `xml:"h-level,attr"`
}

func (n HeadingConfig) CfgString() string {
	return fmt.Sprintf("%v", n)
}

type LinkConfig struct {
	XMLName xml.Name `xml:"heading-config"`
	Href    string   `xml:"href,attr"`
	Title   string   `xml:"title,attr"`
	Target  string   `xml:"target,attr"`
}

func (n LinkConfig) CfgString() string {
	return fmt.Sprintf("%v", n)
}

type OrderedListConfig struct {
	XMLName xml.Name `xml:"ordered-list-config"`
	Order   int      `xml:"order,attr"`
}

func (n OrderedListConfig) CfgString() string {
	return fmt.Sprintf("%v", n)
}

type ParagraphConfig struct {
	XMLName xml.Name `xml:"paragraph-config"`
	Tight   bool     `xml:"tight,attr"`
}

func (n ParagraphConfig) CfgString() string {
	return fmt.Sprintf("%v", n)
}

type CellConfig struct {
	XMLName xml.Name `xml:"empty-config"`
	IsHead  bool     `xml:"is-head,attr"`
	Align   string   `xml:"align,attr"`
}

func (n CellConfig) CfgString() string {
	return fmt.Sprintf("%v", n)
}

type EmptyConfig struct {
	XMLName xml.Name `xml:"empty-config"`
}

func (n EmptyConfig) CfgString() string {
	return fmt.Sprintf("%v", n)
}

var alignments = []string{"", "L", "C", "R"}

type EnclosureType int

const (
	Blockquote EnclosureType = iota
	BulletList
	Emphasis
	ListItem
	Strikethrough
	Strong
	Table
	TableBody
	TableHead
	TableRow
	Heading
	Link
	OrderedList
	Paragraph
	TableCell
)

var enclosureTypes = []string{"Blockquote", "BulletList", "Emphasis", "ListItem", "Strikethrough", "Strong", "Table", "TableBody", "TableHead", "TableRow", "Heading", "Link", "OrderedList", "Paragraph", "TableCell"}

func (t EnclosureType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: enclosureTypes[t]}, nil
}

type ContentType int

const (
	CodeBlock ContentType = iota
	CodeInline
	HTMLBlock
	HTMLInline
	Text
)

var contentTypes = []string{"CodeBlock", "CodeInline", "HTMLBlock", "HTMLInline", "Text"}

func (t ContentType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: contentTypes[t]}, nil
}

type BreakType int

const (
	Soft BreakType = iota
	Hard
	HardRule
)

var breakTypes = []string{"soft", "hard", "hardrule"}

func (t BreakType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: breakTypes[t]}, nil
}

func convertAll(tokens ...markdown.Token) []Node {
	nodes := []Node{}
	for _, token := range tokens {
		nodes = append(nodes, convert(token))
	}
	return nodes
}

func convert(token markdown.Token) Node {
	name := reflect.TypeOf(token).Elem().Name()
	switch name {
	case "BlockquoteClose":
		return Close{Type: Blockquote}
	case "BlockquoteOpen":
		return Open{Type: Blockquote, Config: EmptyConfig{}}
	case "BulletListClose":
		return Close{Type: BulletList}
	case "BulletListOpen":
		return Open{Type: BulletList, Config: EmptyConfig{}}
	case "EmphasisClose":
		return Close{Type: Emphasis}
	case "EmphasisOpen":
		return Open{Type: Emphasis, Config: EmptyConfig{}}
	case "ListItemClose":
		return Close{Type: ListItem}
	case "ListItemOpen":
		return Open{Type: ListItem, Config: EmptyConfig{}}
	case "StrikethroughClose":
		return Close{Type: Strikethrough}
	case "StrikethroughOpen":
		return Open{Type: Strikethrough, Config: EmptyConfig{}}
	case "StrongClose":
		return Close{Type: Strong}
	case "StrongOpen":
		return Open{Type: Strong, Config: EmptyConfig{}}
	case "TableClose":
		return Close{Type: Table}
	case "TableOpen":
		return Open{Type: Table, Config: EmptyConfig{}}
	case "TbodyClose":
		return Close{Type: TableBody}
	case "TbodyOpen":
		return Open{Type: TableBody, Config: EmptyConfig{}}
	case "TheadClose":
		return Close{Type: TableHead}
	case "TheadOpen":
		return Open{Type: TableHead, Config: EmptyConfig{}}
	case "TrClose":
		return Close{Type: TableRow}
	case "TrOpen":
		return Open{Type: TableRow, Config: EmptyConfig{}}
	case "HeadingClose":
		return Close{Type: Heading}
	case "HeadingOpen":
		head, _ := token.(*markdown.HeadingOpen)
		return Open{Type: Heading, Config: HeadingConfig{HLevel: head.HLevel}}
	case "LinkClose":
		return Close{Type: Link}
	case "LinkOpen":
		link, _ := token.(*markdown.LinkOpen)
		return Open{Type: Link, Config: LinkConfig{Href: link.Href, Title: link.Title, Target: link.Target}}
	case "OrderedListClose":
		return Close{Type: OrderedList}
	case "OrderedListOpen":
		orderedList, _ := token.(*markdown.OrderedListOpen)
		return Open{Type: OrderedList, Config: OrderedListConfig{Order: orderedList.Order}}
	case "ParagraphClose":
		return Close{Type: Paragraph}
	case "ParagraphOpen":
		paragraph, _ := token.(*markdown.ParagraphOpen)
		return Open{Type: Paragraph, Config: ParagraphConfig{Tight: paragraph.Tight}}
	case "TdClose":
		return Close{Type: TableCell}
	case "TdOpen":
		cell, _ := token.(*markdown.TdOpen)
		return Open{Type: TableCell, Config: CellConfig{IsHead: false, Align: alignments[cell.Align]}}
	case "ThClose":
		return Close{Type: TableCell}
	case "ThOpen":
		cell, _ := token.(*markdown.ThOpen)
		return Open{Type: TableCell, Config: CellConfig{IsHead: true, Align: alignments[cell.Align]}}
	case "CodeBlock":
		content, _ := token.(*markdown.CodeBlock)
		return Content{Content: content.Content, Type: CodeBlock}
	case "CodeInline":
		content, _ := token.(*markdown.CodeInline)
		return Content{Content: content.Content, Type: CodeInline}
	case "HTMLBlock":
		content, _ := token.(*markdown.HTMLBlock)
		return Content{Content: content.Content, Type: HTMLBlock}
	case "HTMLInline":
		content, _ := token.(*markdown.HTMLInline)
		return Content{Content: content.Content, Type: HTMLInline}
	case "Text":
		content, _ := token.(*markdown.Text)
		return Content{Content: content.Content, Type: Text}
	case "Softbreak":
		return Break{Type: Soft}
	case "Hardbreak":
		return Break{Type: Hard}
	case "Hr":
		return Break{Type: HardRule}
	case "Fence":
		fence, _ := token.(*markdown.Fence)
		return Fence{Params: fence.Params, Content: fence.Content}
	case "Image":
		image, _ := token.(*markdown.Image)
		return Image{Children: convertAll(image.Tokens...), Title: image.Title, Src: image.Src}
	case "Inline":
		inline, _ := token.(*markdown.Inline)
		return Inline{Children: convertAll(inline.Children...), Content: inline.Content}
	default:
		fmt.Println(name)
		return nil
	}
}
