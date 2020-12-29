package main

import "testing"

func TestParseRowCol(t *testing.T) {
	for _, tc := range []struct {
		s        string
		row, col int
	}{
		{
			s:   "FBFBBFFRLR",
			row: 44,
			col: 5,
		}, {
			s:   "BFFFBBFRRR",
			row: 70,
			col: 7,
		}, {
			s:   "FFFBBBFRRR",
			row: 14,
			col: 7,
		}, {
			s:   "BBFFBBFRLL",
			row: 102,
			col: 4,
		}, {
			s:   "FFFFFFFRRR",
			row: 0,
			col: 7,
		}, {
			s:   "BBBBBBBRRR",
			row: 127,
			col: 7,
		}, {
			s:   "FFFFFFFLLL",
			row: 0,
			col: 0,
		}, {
			s:   "BBBBBBBLLL",
			row: 127,
			col: 0,
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			row, col, err := parseRowCol(tc.s)
			if err != nil {
				t.Errorf("bad parse: %v", err)
			}
			if row != tc.row {
				t.Errorf("got row %d, want row %d", row, tc.row)
			}
			if col != tc.col {
				t.Errorf("got col %d, want col %d", col, tc.col)
			}
		})
	}
}
