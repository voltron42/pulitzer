package model

type Node interface {
	// todo --
}

type NodeList []Node

type Block struct {
	Formatting BlockFormat
	Children   NodeList
}

type Inline struct {
	Formatting InlineFormat
	Children   NodeList
}

type Text string

type BlockFormat int

const (
	BlockQuote BlockFormat = iota
	CodeBlock
	Paragraph
	HTML
)

type InlineFormat int

const (
	Emphasis      InlineFormat = 1
	Strong        InlineFormat = 2
	Strikethrough InlineFormat = 4
	Code          InlineFormat = 8
	Html          InlineFormat = 16
)

type Break int

const (
	Hard Break = iota
	Soft
	Hr
)

type Heading struct {
	Level    int
	Children NodeList
}

type Image struct {
	Src   string
	Title string
}

type Link struct {
	Href     string
	Title    string
	Target   string
	Children NodeList
}

type Table struct {
	Head TableHead
	Body TableBody
}

type TableHead []ColumnHeader

type ColumnHeader NodeList

type TableBody []TableRow

type TableRow []TableCell

type TableCell NodeList

type List struct {
	Type  ListType
	Items []ListItem
}

type ListItem NodeList

type ListType bool

const (
	Ordered ListType = true
	Bullet  ListType = false
)
