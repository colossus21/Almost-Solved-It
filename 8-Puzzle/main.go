package main

import (
"errors"
"fmt"
"reflect"
"strconv"
)

const SHOW_STEPS = true

type Node struct {
	parent *Node
	children []*Node
	level int

	value [][]int
	goal [][]int

	pointX int
	pointY int
}
func (g *Node) Print(msg ...string) {
	if msg != nil {
		fmt.Println(msg)
	}
	for _,v := range g.value {
		fmt.Println(v)
	}
	fmt.Println()
}
func (g *Node) SetBox(n int) {
	for i:=0; i< len(g.value); i++ {
		for j:=0; j<len(g.value[i]); j++ {
			if n == g.value[i][j] {
				g.pointX = i
				g.pointY = j
			}
		}
	}
}
func (g *Node) moveLeft() error {
	if (g.pointY > 0) {
		x := g.pointX
		y := g.pointY
		grid := g.value
		grid[x][y], grid[x][y-1] = grid[x][y-1], grid[x][y]
		g.pointY=y-1
		return nil
	}
	return errors.New("cannot move left")
}
func (g *Node) moveRight() error {
	if (g.pointY < len(g.value[0])-1) {
		x := g.pointX
		y := g.pointY
		grid := g.value
		grid[x][y], grid[x][y+1] = grid[x][y+1], grid[x][y]
		g.pointY=y+1
		return nil
	}
	return errors.New("cannot move right")
}
func (g *Node) moveUp() error {
	if (g.pointX > 0) {
		x := g.pointX
		y := g.pointY
		grid := g.value
		grid[x-1][y], grid[x][y] = grid[x][y], grid[x-1][y]
		g.pointX = x-1
		return nil
	}
	return errors.New("cannot move up")
}
func (g *Node) moveDown() error {
	if (g.pointX < len(g.value)-1) {
		x := g.pointX
		y := g.pointY
		grid := g.value
		grid[x+1][y], grid[x][y] = grid[x][y], grid[x+1][y]
		g.pointX = x+1
		return nil
	}
	return errors.New("cannot move down")
}
func (g *Node) getClone() *Node {
	node := new(Node)

	node.parent = g.parent
	node.goal = g.goal
	node.pointX = g.pointX
	node.pointY = g.pointY

	clone := [][]int {
		{0,0,0},
		{0,0,0},
		{0,0,0},
	}
	for i:=0; i<len(g.value); i++ {
		for j:=0;j< len(g.value[i]);j++ {
			clone[i][j] = g.value[i][j]
		}
	}

	node.value = clone
	return node
}
func (g *Node) getChildren() []*Node {
	if SHOW_STEPS {
		g.Print("PARENT NODE:")
	}
	childNodes := []*Node{}
	//Left
	if (g.pointY > 0) {
		NodeLeft := g.getClone()
		NodeLeft.moveLeft()
		if SHOW_STEPS {
			NodeLeft.Print("Left")
		}
		NodeLeft.parent = g
		childNodes = append(childNodes, NodeLeft)
	}
	//Right
	if (g.pointY < len(g.value[0])-1) {
		NodeRight := g.getClone()
		NodeRight.moveRight()
		if SHOW_STEPS {
			NodeRight.Print("Right")
		}
		NodeRight.parent = g
		childNodes = append(childNodes, NodeRight)
	}
	//Up
	if (g.pointX > 0) {
		NodeUp := g.getClone()
		NodeUp.moveUp()
		if SHOW_STEPS {
			NodeUp.Print("Up")
		}
		NodeUp.parent = g
		childNodes = append(childNodes, NodeUp)
	}
	//Down
	if (g.pointX < len(g.value)-1) {
		NodeDown := g.getClone()
		NodeDown.moveDown()
		if SHOW_STEPS {
			NodeDown.Print("Down")
		}
		NodeDown.parent = g
		childNodes = append(childNodes, NodeDown)
	}

	return childNodes
}
func (g *Node) isGoalReached() bool {
	return reflect.DeepEqual(g.value, g.goal)
}
func (g *Node) setStates(startState [][]int, goalState [][]int) {
	g.value = startState
	g.goal = goalState
}

func Solve (start *Node, depth int) {
	level := depth
	//Store goal nodes
	var goalNodes = []*Node{}
	var stack = []*Node {start}
	for depth > 0 {
		children := []*Node {}
		//Get all children of all nodes in the stack
		for i, v := range stack {
			children = append(children, v.getChildren()...)
			stack = stack[0:i+1]
		}
		//Reset the stack
		stack = stack[0:0]
		//Adds the children to stack which are not the goal state
		for _, v := range children {
			if v.isGoalReached() {
				v.level = (level-depth)+1
				goalNodes = append(goalNodes, v)
			} else {
				stack = append(stack, v)
			}
		}
		//Continue for [depth] level(s)
		depth--
	}
	for i, v := range goalNodes {
		v.Print("Solution ", strconv.Itoa(i+1), "found in Level", strconv.Itoa(v.level))
		parentNode := v.parent
		for parentNode!=nil {
			parentNode.Print("Backtracking Solution",strconv.Itoa(i+1))
			parentNode = parentNode.parent
		}
	}
}

func main() {
	startState := [][]int {
		{3,0,7},
		{2,8,1},
		{6,4,5},
	}

	goalState 	:= [][]int {
		{3,8,7},
		{2,0,1},
		{6,4,5},
	}
	//Initiate Parent Node
	Puzzle := new(Node)
	//Start State
	Puzzle.setStates(startState, goalState)
	//SetBox is used to indicate integer used for the blank box. Here 0 is used as the blank box.
	Puzzle.SetBox(0)
	//Depth indicates number of levels the graph should traverse to find the solution
	Solve(Puzzle, 3)
}
