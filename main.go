package main

import (
	"fmt"
	"os"
	"sort"
	"strongly_connected_components/src"
	"time"
)

func printTiming(edges, nodes int, start time.Time, timeReading, elapsedGettingSCC time.Duration) {
	fmt.Printf("With an input of %d nodes and %d edges:\n", nodes, edges)
	fmt.Printf("\tRead file took: %s\n", timeReading)
	fmt.Printf("\tGet SCCs took: %s\n", elapsedGettingSCC)
	fmt.Printf("\tTotal finish time: %s\n", time.Since(start))
}

func main() {
	fmt.Println("Get the last five SCCs")
	start := time.Now()
	fileName := "./data/SCC.txt"
	// Get file name from arguments
	if len(os.Args) > 1 {
		fileName = "./data/" + os.Args[1]
	}
	amountOfNodes, edgesList := src.ReadFile(fileName)
	elapsedReading := time.Since(start)

	// Create graph G_rev first
	graph := new(src.Graph)
	graph.Nodes = make([]*src.Node, amountOfNodes+1)
	graph.CreateGraph(edgesList, true)

	// Sort in decreasing order
	startSCC := time.Now()
	maxFiveSCC := graph.GetMaxSCCs(edgesList, amountOfNodes)
	sort.Sort(sort.Reverse(sort.IntSlice(maxFiveSCC[:])))
	elapsedGettingSCC := time.Since(startSCC)
	fmt.Printf("Last five SCCs: %v\n", maxFiveSCC)

	printTiming(len(edgesList), amountOfNodes, start, elapsedReading, elapsedGettingSCC)
}
