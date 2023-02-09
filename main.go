package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type twoInt [2]int

type Node struct {
	label     int // nombre del nodo
	visited   int // 0 not visited, -1 gray, 1 visited
	previo    int // indice a nodo anterior en recorrido
	neighbors []int
}

type Graph struct {
	nodes []*Node
}

var (
	t     int
	pila  []int
	arbol []int
)

func (g *Graph) display() {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		near := g.nodes[i]
		if near == nil {
			continue
		}
		s += fmt.Sprintf(" %8d --> ", near.label)
		for _, j := range near.neighbors {
			s += fmt.Sprintf(" %8d| ", g.nodes[j].label)
		}
		fmt.Println(s)
		s = ""
	}
}

func (pg *Graph) creatGraph(lsls []twoInt, rev bool, start time.Time) {

}

func (g *Graph) dfs(rev bool) {

}

func dfsVisit(g *Graph, i int, u *Node, rev bool) {

}

func listOrder(pg *Graph) {

}

func findLeader(g *Graph) {

}

func listTree(pg *Graph) {

}

func readFile(name string, start time.Time) (mxval int, lisnod []twoInt) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var anod twoInt
	var i int
	var elapsed time.Duration
	for scanner.Scan() {
		lineStr := scanner.Text()
		fmt.Sscanf(lineStr, "%d %d", &anod[0], &anod[1])
		lisnod = append(lisnod, anod)
		if mxval < anod[0] {
			mxval = anod[0]
		}
		if mxval < anod[1] {
			mxval = anod[1]
		}
		i++
		if i%100000 == 0 {
			elapsed = time.Since(start)
			fmt.Printf("%8d %v %8d %8d \n", i, elapsed, anod[0], anod[1])
		}
	}
	return
}

func main() {
	start := time.Now()
	name := "./data/SCC.txt"
	if len(os.Args) > 1 {
		name = "./data/" + os.Args[1]
	}

	nr, lisnod := readFile(name, start)

	elapsed := time.Since(start)
	fmt.Printf("Reading took %s\n", elapsed)
	fmt.Printf("Input %10d  Edges %10d Nodes\n", len(lisnod), nr)

	pg := new(Graph)
	pg.nodes = make([]*Node, nr+1)
	pg.creatGraph(lisnod, true, start)
	elapsed = time.Since(start)
	fmt.Printf("Created %10d Nodes  %s\n", len(pg.nodes), elapsed)

	pg.display()

	pg.dfs(false)
	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)

	listOrder(pg)

	pgr := new(Graph)
	pgr.nodes = make([]*Node, nr+1)
	pgr.creatGraph(lisnod, false, start)
	elapsed = time.Since(start)
	fmt.Printf("Created %10d Nodes  %s\n", len(pg.nodes), elapsed)

	pgr.display()

	pgr.dfs(true)
	elapsed = time.Since(start)
	fmt.Printf("Finish time %s \n", elapsed)
	elapsed = time.Since(start)
	//    findLeader(pgr)
	listTree(pgr)
	fmt.Printf("Created %10d Nodes  %s\n", len(pg.nodes), elapsed)
}
