package httl

type Node interface {
	Accept(v Visitor) error
	Parent() Node
	Children() []Node
}
