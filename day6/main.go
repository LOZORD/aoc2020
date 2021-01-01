package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/glog"
)

var part2Flag = flag.Bool("part2", false, "Toggles between part 1 and part 2 functionality.")

func main() {
	flag.Parse()
	if err := doMain(os.Stdin, os.Stdout); err != nil {
		glog.Errorf("error: %v", err)
	}
}

func doMain(reader io.Reader, writer io.Writer) error {
	ab, err := ioutil.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read all: %w", err)
	}
	all := string(ab)

	// Part 1.
	if !*part2Flag {
		sum := 0
		groups := strings.Split(all, "\n\n")
		for i, g := range groups {
			glog.Infof("GROUP %d: %q", i, g)
			m := map[rune]bool{}
			for j, p := range strings.Split(g, "\n") {
				glog.Infof("\tPERSON %d: %q", j, p)
				for _, c := range p {
					m[c] = true
				}
			}
			sum += len(m)
		}
		fmt.Fprintf(writer, "SUM = %d\n", sum)
	} else {
		// Part 2.
		sum := 0
		groups := strings.Split(all, "\n\n")
		for _, g := range groups {
			persons := strings.Split(g, "\n")
			m := map[rune]int{}
			for _, p := range persons {
				for _, c := range p {
					m[c] = m[c] + 1
				}
			}

			for _, n := range m {
				if n == len(persons) {
					sum++
				}
			}
		}

		fmt.Fprintf(writer, "SUM = %d\n", sum)
	}

	return nil
}
