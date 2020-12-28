package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/golang/glog"
)

// Flags.
var (
	countFlag = flag.Uint("count", 2, "Count of numbers to search for to get to 2020.")
)

func main() {
	flag.Parse()
	if err := doMain(os.Stdin, os.Stdout, *countFlag); err != nil {
		glog.Exitf("error: %v", err)
	}
}

func doMain(r io.Reader, w io.Writer, count uint) error {
	s := map[int]bool{}

	br := bufio.NewScanner(r)
	for br.Scan() {
		t := br.Text()
		n, err := strconv.Atoi(t)
		if err != nil {
			return fmt.Errorf("failed to parse %q as a number: %w", t, err)
		}

		if count == 2 {
			m := 2020 - n
			if ok := s[m]; ok {
				if _, err := fmt.Fprintf(w, "%d * %d = %d\n", m, n, m*n); err != nil {
					return fmt.Errorf("failed to write output: %w", err)
				}
				return nil
			}

			s[n] = true
		} else if count == 3 {
			for m := range s {
				o := 2020 - m - n
				if ok := s[o]; ok {
					if _, err := fmt.Fprintf(w, "%d * %d * %d = %d\n", m, n, o, m*n*o); err != nil {
						return fmt.Errorf("failed to write output: %w", err)
					}
					return nil
				}
			}
			s[n] = true
		} else {
			return fmt.Errorf("unimplemented count %d", count)
		}

	}
	if err := br.Err(); err != nil {
		return fmt.Errorf("failed to scan: %w", err)
	}

	return fmt.Errorf("never found a pair in set %v", s)
}
