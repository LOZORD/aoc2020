package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/golang/glog"
)

var partFlag = flag.Uint("part", 1, "The problem part.")

func main() {
	flag.Parse()
	if err := doMain(os.Stdin, os.Stdout, *partFlag); err != nil {
		glog.Error("error: %v", err)
	}
}
func doMain(r io.Reader, w io.Writer, part uint) error {
	var total, valid uint

	br := bufio.NewScanner(r)
	for br.Scan() {
		total++

		var lo, hi int
		var r rune
		var pw string
		t := br.Text()
		if _, err := fmt.Sscanf(t, "%d-%d %c: %s", &lo, &hi, &r, &pw); err != nil {
			return fmt.Errorf("failed to scan %q: %v", t, err)
		}

		if part == 1 {
			if isValid1(lo, hi, r, pw) {
				valid++
			}
		} else {
			// Assume part 2.
			if isValid2(lo, hi, r, pw) {
				valid++
				glog.Infof("%d-%d %c: %q is VALID\n", lo, hi, r, pw)
			} else {
				glog.Infof("%d-%d %c: %q is invalid\n", lo, hi, r, pw)
			}
		}

	}
	if err := br.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	fmt.Fprintf(w, "TOTAL = %d\nVALID = %d\n", total, valid)

	return nil
}

func isValid1(lo, hi int, r rune, pw string) bool {
	n := 0
	for _, c := range pw {
		if c == r {
			n++
		}
	}
	return lo <= n && n <= hi
}

func isValid2(fst, snd int, r rune, pw string) bool {
	feq := rune(pw[fst-1]) == r
	seq := rune(pw[snd-1]) == r
	return feq != seq
}
