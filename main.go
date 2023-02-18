package main

import (
	"fmt"
	"os"
	"strongly_connected_components/src"
	"time"
)

func printTiming(edges, nodes int, start time.Time, timeReading time.Duration) {
	fmt.Printf("With an input of %d nodes and %d edges:\n", nodes, edges)
	fmt.Printf("\tReading file took %s\n", timeReading)
	fmt.Printf("\tTotal finish time %s\n", time.Since(start))
}

func main() {
	start := time.Now()
	fileName := "./data/SCC5.txt"
	// Get file name from arguments
	if len(os.Args) > 1 {
		fileName = "./data/" + os.Args[1]
	}

	length, nodesList := src.ReadFile(fileName)
	elapsedReading := time.Since(start)

	// Create graph G_rev first
	graph := new(src.Graph)
	graph.Nodes = make([]*src.Node, length+1)
	graph.CreateGraph(nodesList, true)

	graph.Dfs(false)

	// Create reversed graph G
	reversedGraph := new(src.Graph)
	reversedGraph.Nodes = make([]*src.Node, length+1)
	reversedGraph.CreateGraph(nodesList, false)

	// Print max SCCs
	reversedGraph.PrintMaxSCCs()

	printTiming(len(nodesList), length, start, elapsedReading)
}
