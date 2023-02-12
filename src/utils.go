package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// ReadFile reads a text file to extract the nodes list of it
func ReadFile(name string, start time.Time) (maxvalue int, nodeList []intTuple) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Some error just happened.")
		}
	}(file)
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var anod intTuple
	var i int
	var elapsed time.Duration
	for scanner.Scan() {
		lineStr := scanner.Text()
		_, err2 := fmt.Sscanf(lineStr, "%d %d", &anod[0], &anod[1])
		if err2 != nil {
			return 0, nil
		}
		nodeList = append(nodeList, anod)
		if maxvalue < anod[0] {
			maxvalue = anod[0]
		}
		if maxvalue < anod[1] {
			maxvalue = anod[1]
		}
		i++
		if i%100000 == 0 {
			elapsed = time.Since(start)
			fmt.Printf("%8d %v %8d %8d \n", i, elapsed, anod[0], anod[1])
		}
	}
	return
}
