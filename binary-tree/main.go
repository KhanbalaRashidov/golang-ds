package main

import "fmt"

type Node struct {
	data  string
	left  *Node
	right *Node
}

func InorderPrint(root *Node) {
	if root == nil {
		return
	}
	if root.left != nil {
		InorderPrint(root.left)
	}
	fmt.Println(root.data)
	if root.right != nil {
		InorderPrint(root.right)
	}
}

func main() {
	A := Node{data: "A"}
	B := Node{data: "B"}
	C := Node{data: "C"}
	D := Node{data: "D"}
	E := Node{data: "E"}
	F := Node{data: "F"}
	G := Node{data: "G"}

	A.left = &B
	A.right = &C
	B.left = &D
	B.right = &E
	C.left = &F
	C.right = &G

}
