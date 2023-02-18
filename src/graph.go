package src

import (
	"fmt"
)

type intTuple [2]int

const maxAmountOfSCCs = 5

var (
	count      int
	maxFiveSCC [maxAmountOfSCCs]int
	// [index, value]
	minimumSCC = [2]int{0, 0}
	stack      Stack
)

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
	maxAmountOfDigits := CountDigits(len(g.Nodes) - 1)
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

// insertPathLength in the five-elements array to save the max SCC
func insertPathLength(count int) {
	if count > minimumSCC[1] {
		maxFiveSCC[minimumSCC[0]] = count

		// Update minimum
		minimumSCC = [2]int{0, maxFiveSCC[0]}
		for idx, v := range maxFiveSCC {
			if v < minimumSCC[1] {
				minimumSCC = [2]int{idx, v}
			}
		}
	}
}

func (g *Graph) dfsVisit(index int, rollback, reverse bool) {
	if index == 0 {
		if reverse {
			// Insert path length into the max SCCs array
			insertPathLength(count)
		}
		return
	}

	node := g.Nodes[index]
	if !rollback {
		g.Nodes[index].Visited = true
		g.nodesVisited++

		if reverse {
			count++
		}
	}

	// Look for the next node
	allVisited := true
	for _, neighbor := range node.Neighbors {
		if !g.Nodes[neighbor].Visited {
			allVisited = false

			// Continue to the next node
			g.Nodes[neighbor].Previous = index
			g.dfsVisit(neighbor, false, reverse)
		}
	}

	// Go back to the previous node
	if allVisited {
		if !reverse {
			// Push to stack since all neighbors were visited
			stack.Push(index)
		}
		g.dfsVisit(node.Previous, true, reverse)
	}
	return
}

// Dfs of the graph
func (g *Graph) Dfs(reverse bool) {
	for idx, node := range g.Nodes {
		if node == nil {
			continue
		}

		if !node.Visited {
			g.dfsVisit(idx, false, reverse)
		}
	}
}

func (g *Graph) GetMaxSCCs(edgesList []intTuple, amountOfNodes int) [maxAmountOfSCCs]int {
	// Run DFS
	g.Dfs(false)

	// Create reversed graph G
	reversedGraph := new(Graph)
	reversedGraph.Nodes = make([]*Node, amountOfNodes+1)
	reversedGraph.CreateGraph(edgesList, false)

	// Pop one by one to make DFS to the top of stack
	for !stack.IsEmpty() {
		v, _ := stack.Pop()
		if !reversedGraph.Nodes[v].Visited {
			count = 0
			reversedGraph.dfsVisit(v, false, true)
		}
	}

	return maxFiveSCC
}
