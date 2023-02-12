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

func (g *Graph) Display() {
	s := ""
	for i := 0; i < len(g.Nodes); i++ {
		near := g.Nodes[i]
		if near == nil {
			continue
		}
		s += fmt.Sprintf(" %8d --> ", near.Label)
		for _, j := range near.Neighbors {
			s += fmt.Sprintf(" %8d| ", g.Nodes[j].Label)
		}
		fmt.Println(s)
		s = ""
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
