package 

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	initVal = flag.Float64("init", 0.15, "")
)

func main() {
	flag.Parse()
	curKey := ""
	dest := []string{}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1]
		ts := strings.Split(line, "\t")
		if len(ts) < 1 {
			continue
		}
		if ts[0] != curKey && curKey != "" {
			fmt.Printf("%s\tactive\t%g\t%s\n", curKey, *initVal, strings.Join(dest, "|"))
			dest = []string{}
		}
		for i := 1; i < len(ts)-1; i++ {
			if ts[i] != "" {
				dest = append(dest, ts[i])
			}
		}
		curKey = ts[0]

	}
	fmt.Printf("%s\tactive\t%g\t%s\n", curKey, *initVal, strings.Join(dest, "|"))
}
