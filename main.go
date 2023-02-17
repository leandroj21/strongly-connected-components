package main

import (
	"fmt"
	"os"
	"strongly_connected_components/src"
	"time"
)

func main() {
	start := time.Now()
	fileName := "./data/SCC4.txt"
	// Get file name from arguments
	if len(os.Args) > 1 {
		fileName = "./data/" + os.Args[1]
	}

	lenght, nodesList := src.ReadFile(fileName, start)

	elapsed := time.Since(start)
	fmt.Printf("Reading took %s\n", elapsed)
	fmt.Printf("Input %10d  Edges %10d Nodes\n", len(nodesList), lenght)

	// Create graph G
	graph := new(src.Graph)
	graph.Nodes = make([]*src.Node, lenght+1)
	graph.CreateGraph(nodesList, false)
	elapsed = time.Since(start)
	fmt.Printf("Created %10d Nodes  %s\n", len(graph.Nodes), elapsed)

	graph.Dfs(false)
	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)

	// Create reversed graph G_rev
	reversedGraph := new(src.Graph)
	reversedGraph.Nodes = make([]*src.Node, lenght+1)
	reversedGraph.CreateGraph(nodesList, true)
	elapsed = time.Since(start)
	fmt.Printf("Created %10d Nodes  %s\n", len(graph.Nodes), elapsed)

	// Print SCCs
	reversedGraph.PrintSCC()

	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)
	elapsed = time.Since(start)
	fmt.Printf("Created %10d Nodes  %s\n", len(graph.Nodes), elapsed)
}
