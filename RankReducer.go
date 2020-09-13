package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	debug   = flag.Bool("debug", false, "")
	prob    = flag.Float64("prob", 0.85, "")
	total   = flag.Int("total", 10000, "")
	thresh  = flag.Float64("thresh", 0.0001, "")
	sumRank = 0.0
)

type Msg struct {
	numCon int
	rank   float64
}

type Node struct {
	key     string
	rank    float64
	connect string
}

func output(node Node, weights []Msg) {
	if *debug {
		log.Printf("[%v] -- [%v]", node, weights)
	}
	newRank := 0.0
	for _, msg := range weights {
		r := (*prob / float64(msg.numCon)) * msg.rank
		newRank += r
	}
	newRank += (1 - *prob) * sumRank / float64(*total)
	if math.Abs(newRank-node.rank) < *thresh { // rank doesn't change much
		fmt.Printf("%s\tdeact\t%.7f\t%s\n", node.key, node.rank, node.connect)
	} else { //update
		fmt.Printf("%s\tactive\t%.7f\t%s\n", node.key, newRank, node.connect)
	}

}

func main() {
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	curKey := ""
	curNode := Node{}
	weights := []Msg{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1]
		ts := strings.Split(line, "\t")
		if len(ts) < 3 {
			continue
		}
		if ts[0] != curKey && curKey != "" {
			output(curNode, weights)
			weights = []Msg{}
			curNode = Node{}
		}
		if ts[1] == "msg" {
			nCon, _ := strconv.Atoi(ts[3])
			rank, _ := strconv.ParseFloat(ts[2], 64)
			c := Msg{nCon, rank}
			weights = append(weights, c)
		} else if ts[1] == "tot" {
			rank, _ := strconv.ParseFloat(ts[2], 64)
			if rank != sumRank && sumRank != 0.0 {
				log.Fatal(fmt.Sprintf("2 different sumRank values : %g and %g", rank, sumRank))
			}
			sumRank = rank
		} else {
			rank, _ := strconv.ParseFloat(ts[2], 64)
			curNode = Node{ts[0], rank, ts[3]}
		}
		curKey = ts[0]
	}
	output(curNode, weights)
}
