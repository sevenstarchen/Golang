package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node Node) Print() {

	fmt.Println(node.Value)
}
func (node *Node) SetValue(value int) {
	if node == nil {
		println("setting value to nil node")
		return
	}
	node.Value = value
}

/*func (node *Node) Traverse(){
	if node == nil{
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
	//node.print()
}*/
/*func main(){
	var root Node
	root = Node{Value:3}
	root.Left=&Node{}
	root.Right=&Node{5,nil,nil}
	root.Right.Left=new(Node)
	root.Left.Right=CreateNode(2)
	root.Left.SetValue(3)
	root.Traverse()

}*/
