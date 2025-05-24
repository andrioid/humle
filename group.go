package humle

type Group []Node

func (g *Group) Type() NodeType {
	return NodeGroup
}
