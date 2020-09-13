package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	totRank = 0.0
)

// input : node <--> active|deact <--> rank <--> linked nodes
// output: node <--> msg <--> rank of src <--> # links of source
func main() {
	flag.Parse()
	nodeMap := map[string]float64{}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1]
		ts := strings.Split(line, "\t")
		if len(ts) < 4 {
			continue
		}
		rank, _ := strconv.ParseFloat(ts[2], 64)
		nodeMap[ts[0]] = rank
		fmt.Println(line)
		if ts[1] != "active" {
			continue
		}
		dest := strings.Split(ts[3], "|")

		for _, d := range dest {
			if d != "" {
				fmt.Printf("%s\tmsg\t%s\t%d\n", d, ts[2], len(dest))
			}
		}

	}
	for _, v := range nodeMap {
		totRank += v
	}
	log.Printf("Total rank : %g", totRank)
	for k := range nodeMap {
		fmt.Printf("%s\ttot\t%g\n", k, totRank)
	}

}
