package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/golang/glog"
)

var part2Flag = flag.Bool("part2", false, "Activate part 2 functionality.")

func main() {
	flag.Parse()
	if err := doMain(os.Stdin, os.Stdout); err != nil {
		glog.Errorf("error: %v", err)
	}
}

func doMain(r io.Reader, w io.Writer) error {
	var hi int
	br := bufio.NewScanner(r)

	var ids []int
	total := 0
	for br.Scan() {
		t := br.Text()
		row, col, err := parseRowCol(t)
		if err != nil {
			return fmt.Errorf("failed to parse %q: %v", t, err)
		}
		total++

		glog.Infof("parsed %q into row=%d col=%d", t, row, col)

		n := bpID(row, col)

		if n > hi {
			hi = n
		}

		ids = append(ids, n)
	}
	if err := br.Err(); err != nil {
		return fmt.Errorf("failed to scan: %w", err)
	}

	glog.Infof("got %d total entries", total)

	fmt.Fprintf(w, "MAXIMUM = %d\n", hi)

	if !*part2Flag {
		return nil
	}

	var mySeat int
	sort.Ints(ids)
	for i := range ids {
		if i == 0 || i == len(ids)-1 {
			continue
		}

		if ids[i]+1 != ids[i+1] {
			mySeat = ids[i] + 1
			break
		}
	}

	fmt.Fprintf(w, "MY_SEAT = %d\n", mySeat)
	return nil
}

func parseRowCol(s string) (int, int, error) {
	if len(s) != 10 {
		return 0, 0, fmt.Errorf("got bad length of %d", len(s))
	}

	rs := s[:7]
	row := 0
	for i, c := range rs {
		if c == 'B' {
			row += 1 << (len(rs) - i - 1)
		} else if c != 'F' {
			return 0, 0, fmt.Errorf("bad row letter: %c", c)
		}
	}

	cs := s[7:]
	col := 0
	for i, c := range cs {
		if c == 'R' {
			col += 1 << (len(cs) - i - 1)
		} else if c != 'L' {
			return 0, 0, fmt.Errorf("bad col letter: %c", c)
		}
	}

	return row, col, nil
}

func bpID(row, col int) int {
	return 8*row + col
}
