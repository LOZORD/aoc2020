package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

// var slopesFlag = flag.String("slopes", "", "Comma-separated pairs of right-down numbers to use for slopes.")
var slopesFlag = pflag.IntSlice("slopes", nil, "List of pair-wise right-down numbers to use for slopes.")

// pbpaste | ./day3 | awk 'BEGIN { total = 1 } { total *= $3 } END { print total }'
// pbpaste | ./day3 --slopes=1,1,3,1,5,1,7,1,1,2 | awk 'BEGIN { total = 1 } { total *= $3 } END { printf "%f\n", total }'

func main() {
	pflag.Parse()
	if err := doMain(os.Stdin, os.Stdout, *slopesFlag); err != nil {
		glog.Errorf("error: %v", err)
	}
}

func doMain(r io.Reader, w io.Writer, slopeNums []int) error {
	var grid [][]bool // true => no tree; false => tree

	br := bufio.NewScanner(r)

	for br.Scan() {
		line := br.Text()
		row := make([]bool, len(line))
		for i, c := range line {
			row[i] = (c == '.')
		}
		grid = append(grid, row)
	}
	if err := br.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	// var slopes [][2]int
	// for _, s := strings.Split(sf, ",") {
	// 	var right, down int

	// 	slopes = append(slopes)
	// }
	if len(slopeNums)%2 != 0 {
		return fmt.Errorf("uneven slope pairs: %d", len(slopeNums))
	}

	total := 1
	for i := 0; i < len(slopeNums); i += 2 {
		right := slopeNums[i]
		down := slopeNums[i+1]
		trees := treesOnSlope(grid, right, down)
		fmt.Fprintf(w, "R%dD%d = %d\n", right, down, trees)
		total *= trees
	}
	fmt.Fprintf(w, "TOTAL = %d", total)

	return nil
}

func treesOnSlope(grid [][]bool, right, down int) int {
	width := len(grid[0])
	height := len(grid)
	x, y, trees := 0, 0, 0
	for y < height {
		if grid[y][x] == false {
			trees++
		}

		y = (y + down)
		x = (x + right) % width
	}
	return trees
}
