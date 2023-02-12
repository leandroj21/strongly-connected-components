package src

import (
	"fmt"
)

var (
	t     int
	stack []int
	arbol []int
)

type intTuple [2]int

type Node struct {
	Label     int
	Visited   int
	Previous  int
	Neighbors []int
}

type Graph struct {
	Nodes []*Node
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
			g.Nodes[from] = new(Node)
			g.Nodes[from].Label = from
			g.Nodes[from].Previous = -1
		}
		nodeFrom = g.Nodes[from]
		if g.Nodes[to] == nil {
			g.Nodes[to] = new(Node)
			g.Nodes[to].Label = to
			g.Nodes[to].Previous = -1
		}

		// Append neighbor
		nodeFrom.Neighbors = append(nodeFrom.Neighbors, to)
	}
}

// Display display the graph in an user-friendly way
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

func (g *Graph) Dfs(rev bool) {

}

// Operations

func DfsVisit(g *Graph, i int, u *Node, rev bool) {

}

func ListOrder(pg *Graph) {

}

func FindLeader(g *Graph) {

}

func ListTree(pg *Graph) {

}
