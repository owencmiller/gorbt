package main

import "fmt"

type Node struct {
	val int
	left *Node
	right *Node
}

type Tree struct{
	root *Node
}

func make_node(val int) *Node{
	return &Node{val:val}
}

func (tree Tree) preorder(){
	fmt.Printf("---Preorder Traversal---\n")
	preorder_traversal(tree.root)
}
func preorder_traversal(node *Node){
	if node != nil{
		fmt.Printf("%d \n", node.val)
		preorder_traversal(node.left)
		preorder_traversal(node.right)
	}
}

func (tree *Tree) insert(val int){
	node := make_node(val)
	tree.root = insert_with_node(tree.root, node)
}
func insert_with_node(root *Node, node *Node) *Node{
	if root == nil{
		return node
	}
	if node.val < root.val{
		root.left = insert_with_node(root.left, node)
	}else if node.val > root.val{
		root.right = insert_with_node(root.right, node)
	}

	return root
}




func (tree *Tree) find(val int) *Node{
	return find_with_node(tree.root, val)
}
func find_with_node(root *Node, val int) *Node{
	if root == nil || root.val == val{
		return root
	}
	if val > root.val{
		return find_with_node(root.right, val)
	}
	return find_with_node(root.left, val)
}

func (tree *Tree) contains(val int) bool{
	if tree.find(val) == nil{
		return false
	}
	return true
}

func main(){
	tree := Tree{}
	tree.preorder()
	tree.insert(3)
	tree.insert(2)
	tree.insert(1)
	tree.insert(5)
	tree.insert(4)
	tree.preorder()
	fmt.Printf("%v \n", tree.contains(5))
	fmt.Printf("%v \n", tree.contains(8))
	fmt.Printf("%v \n", tree.contains(1))
}
