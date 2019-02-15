package main

import (
	"fmt"
	"strings"
)

type Node struct {
	parent *Node
	children []*Node
	name string
}

func BFS(start string, goal string, nodeMap map[string]*Node) *Node {
	fmt.Println("BFS:")
	startNode := nodeMap[start]
	goalNode := nodeMap[goal]

	queue := []*Node {startNode}

	for len(queue)>0 {
		for _, v := range queue {
			fmt.Print(v.name," ")
			if v == goalNode {
				fmt.Println()
				return v
			}
			queue = queue[1:]
			queue = append(queue, v.children...)
		}
	}
	return new(Node)
}

func DFS(start string, goal string, nodeMap map[string]*Node) *Node {
	fmt.Println()
	fmt.Println("DFS:")
	startNode := nodeMap[start]
	goalNode := nodeMap[goal]

	stack := []*Node {startNode}

	for len(stack)>0 {

		for i:=len(stack)-1; i>=0; i-- {
			currentNode := stack[i]
			fmt.Print(currentNode.name, " ")
			if currentNode == goalNode {
				fmt.Println()
				return currentNode
			}
			stack = stack[:len(stack)-1]
			stack = append(stack, currentNode.children...)
		}
	}
	return new(Node)
}

func NewNode(name string) *Node {
	node := new(Node)
	node.name = name
	node.children = []*Node{}
	node.parent = nil
	return node
}

func MakeGraph(input string, connections string)  map[string]*Node{
	nodeMap := map[string]*Node {}
	fields := strings.Fields(input)
	for i:=0; i<len(fields); i++ {
		nodeName := string(fields[i])
		nodeMap[nodeName] = NewNode(nodeName)
	}
	for i:=0; i<=len(connections)-2; i+=4 {
		node1, ok1 := nodeMap[string(connections[i])]
		node2, ok2 := nodeMap[string(connections[i+2])]

		if(ok1 && ok2) {
			node2.parent = node1
			node1.children = append(node1.children, node2)
		}
	}

	return nodeMap
}

func main() {
	input := "S A B C D E G" //Nodes separated by space
	connections := "S A S B S C A D A E A G B G C G" //Connection pairs separated by space
	graph := MakeGraph(input, connections) //Generate graph based on the given inputs
	BFS("S", "G", graph) //Perform BFS on the given graph
	DFS("S", "G", graph) //Perform DFS on the given graph
}