package src

import (
	"fmt"
)

var (
	t     int
	tree  []int
	stack Stack
)

type intTuple [2]int

type Node struct {
	Label     int
	Visited   bool
	Previous  int
	Neighbors []int
}

type Graph struct {
	Nodes        []*Node
	nodesVisited int
}

// createNode create a node in the graph
func (g *Graph) createNode(position int) {
	g.Nodes[position] = new(Node)
	// Same label as its position
	g.Nodes[position].Label = position
}

// CreateGraph create the graph from a list of pairs of nodes
func (g *Graph) CreateGraph(pairOfNodesList []intTuple, reverse bool) {
	var nodeFrom *Node
	for _, pair := range pairOfNodesList {
		from, to := pair[0], pair[1]
		if reverse {
			from, to = pair[1], pair[0]
		}

		// Create the nodes in the graph if they do not exist
		if g.Nodes[from] == nil {
			g.createNode(from)
		}
		nodeFrom = g.Nodes[from]
		if g.Nodes[to] == nil {
			g.createNode(to)
		}

		// Append neighbor
		nodeFrom.Neighbors = append(nodeFrom.Neighbors, to)
	}
}

// Display the graph in a user-friendly way
func (g *Graph) Display() {
	str, separator := "", ","
	maxAmountOfDigits := CountDigits(len(g.Nodes))
	for _, node := range g.Nodes {
		if node == nil {
			continue
		}

		str += fmt.Sprintf("%*d --> ", maxAmountOfDigits, node.Label)
		for index, neighbor := range node.Neighbors {
			// To not print the final comma
			if index == len(node.Neighbors)-1 {
				separator = ""
			}
			str += fmt.Sprintf("%*d%s ", maxAmountOfDigits, g.Nodes[neighbor].Label, separator)
		}
		fmt.Println(str)
		str, separator = "", ","
	}
}

// Todo: implement reverse DFS
func (g *Graph) depthSearch(index int, rollback, reverse bool) {
	if index == 0 {
		return
	}

	node := g.Nodes[index]
	if !rollback {
		g.Nodes[index].Visited = true
		g.nodesVisited++

		// Push to stack
		stack.Push(index)
	}

	// Look for the next node
	allVisited := true
	for _, neighbor := range node.Neighbors {
		if !g.Nodes[neighbor].Visited {
			allVisited = false

			// Continue to the next node
			g.Nodes[neighbor].Previous = index
			g.depthSearch(neighbor, false, reverse)
		}
	}

	// Go back to the previous node
	if allVisited {
		g.depthSearch(node.Previous, true, reverse)
	}
	return
}

// Dfs Do a DFS of the graph
func (g *Graph) Dfs(reverse bool) {
	for idx, node := range g.Nodes {
		if node == nil {
			continue
		}

		if !node.Visited {
			g.depthSearch(idx, false, reverse)
		}
	}

	// Todo: is it right to clear the stack?
	stack.Clear()
}

// Operations

func ListOrder(pg *Graph) {

}

func FindLeader(g *Graph) {

}

func ListTree(pg *Graph) {

}
