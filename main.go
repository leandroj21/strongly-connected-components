package main

import (
	"fmt"
	"os"
	"strongly_connected_components/src"
	"time"
)

func main() {
	start := time.Now()
	fileName := "./data/SCC5.txt"
	// Get file name from arguments
	if len(os.Args) > 1 {
		fileName = "./data/" + os.Args[1]
	}

	length, nodesList := src.ReadFile(fileName, start)

	elapsed := time.Since(start)
	fmt.Printf("Reading took %s\n", elapsed)
	fmt.Printf("Input %10d  Edges %10d Nodes\n", len(nodesList), length)

	// Create graph G_rev first
	graph := new(src.Graph)
	graph.Nodes = make([]*src.Node, length+1)
	graph.CreateGraph(nodesList, true)
	elapsed = time.Since(start)
	fmt.Printf("Created %10d Nodes  %s\n", len(graph.Nodes), elapsed)

	graph.Dfs(false)
	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)

	// Create reversed graph G
	reversedGraph := new(src.Graph)
	reversedGraph.Nodes = make([]*src.Node, length+1)
	reversedGraph.CreateGraph(nodesList, false)
	elapsed = time.Since(start)
	fmt.Printf("Created %10d Nodes  %s\n", len(graph.Nodes), elapsed)

	// Print max SCCs
	reversedGraph.PrintMaxSCCs()

	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)
	elapsed = time.Since(start)
	fmt.Printf("Created %10d Nodes  %s\n", len(graph.Nodes), elapsed)
}
