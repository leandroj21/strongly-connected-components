package main

import (
	"fmt"
	"os"
	"strongly_connected_components/src"
	"time"
)

func main() {
	start := time.Now()
	name := "./data/SCC2.txt"
	if len(os.Args) > 1 {
		name = "./data/" + os.Args[1]
	}

	nr, lisnod := src.ReadFile(name, start)

	elapsed := time.Since(start)
	fmt.Printf("Reading took %s\n", elapsed)
	fmt.Printf("Input %10d  Edges %10d Nodes\n", len(lisnod), nr)

	pg := new(src.Graph)
	pg.Nodes = make([]*src.Node, nr+1)
	pg.CreateGraph(lisnod, true)
	elapsed = time.Since(start)
	fmt.Printf("Created %10d Nodes  %s\n", len(pg.Nodes), elapsed)

	//pg.Display()

	pg.Dfs(false)
	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)

	src.ListOrder(pg)

	pgr := new(src.Graph)
	pgr.Nodes = make([]*src.Node, nr+1)
	pgr.CreateGraph(lisnod, false)
	elapsed = time.Since(start)
	fmt.Printf("Created %10d Nodes  %s\n", len(pg.Nodes), elapsed)

	//pgr.Display()

	pgr.Dfs(true)
	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)
	elapsed = time.Since(start)
	//    findLeader(pgr)
	src.ListTree(pgr)
	fmt.Printf("Created %10d Nodes  %s\n", len(pg.Nodes), elapsed)
}
