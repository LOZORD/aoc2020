package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

var (
	part2Flag = flag.Bool("part2", false, "Activate part 2 logic.")
)

func main() {
	flag.Parse()
	if err := doMain(os.Stdin, os.Stdout); err != nil {
		glog.Errorf("error: %v", err)
	}
}

const (
	byr = "byr"
	iyr = "iyr"
	eyr = "eyr"
	hgt = "hgt"
	hcl = "hcl"
	ecl = "ecl"
	pid = "pid"
	cid = "cid"
)

/*
byr (Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid
*/

var (
	required = map[string]bool{
		byr: true,
		iyr: true,
		eyr: true,
		hgt: true,
		hcl: true,
		ecl: true,
		pid: true,
		// Not cid.
	}

	eclSet = map[string]bool{
		"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true,
	}

	rgx = regexp.MustCompile("(([a-z]+:.+\n?)*\n)")
	spc = regexp.MustCompile("\\s+")
	hex = regexp.MustCompile("^#[0-9a-f]{6}$")
	// prx = regexp.MustCompile("0*\\d{9}")
	prx = regexp.MustCompile("^\\d{9}$")
)

// ((([a-z]+):.+\n?)*\n)

// (([a-z]+:.+\n?)*\n)

func doMain(r io.Reader, w io.Writer) error {
	all, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read all: %w", err)
	}

	// requires (pbpaste; echo) | ./day4
	// i.e. this parse requires a file that ends in \n\n I think :/
	entries := rgx.FindAllString(string(all), -1)
	if len(entries) < 1 {
		return fmt.Errorf("bad parse on %v: %v", rgx, entries)
	}

	fmt.Fprintf(w, "TOTAL: %d\n", len(entries))

	valid := 0
	for i, e := range entries {
		clean := strings.TrimSpace(spc.ReplaceAllString(e, " "))
		// glog.Infof("got entry %d: %q", i, clean)

		fields := strings.Split(clean, " ")
		got := map[string]string{}
		for _, f := range fields {
			p := strings.Split(f, ":")
			if len(p) != 2 {
				return fmt.Errorf("weird pair from %q: %v", f, p)
			}
			got[p[0]] = p[1]
		}

		glog.V(2).Infof("got %v from %d -> %q", got, i, clean)

		if isValid(got) {
			valid++
		}
	}

	fmt.Fprintf(w, "VALID: %d\n", valid)
	return nil
}

func isValid(m map[string]string) bool {
	for r := range required {
		if _, ok := m[r]; !ok {
			return false
		}
	}

	if !*part2Flag {
		return true
	}

	if n, _ := strconv.Atoi(m[byr]); !(1920 <= n && n <= 2002) {
		glog.Warningf("invalid byr: %d", n)
		return false
	}
	if n, _ := strconv.Atoi(m[iyr]); !(2010 <= n && n <= 2020) {
		glog.Warningf("invalid iyr: %d", n)
		return false
	}
	if n, _ := strconv.Atoi(m[eyr]); !(2020 <= n && n <= 2030) {
		glog.Warningf("invalid eyr: %d", n)
		return false
	}
	if !strings.HasSuffix(m[hgt], "in") && !strings.HasSuffix(m[hgt], "cm") {
		glog.Warningf("invalid height missing in or cm %q", m[hgt])
		return false
	}
	if cm := strings.TrimSuffix(m[hgt], "cm"); cm != m[hgt] {
		if n, _ := strconv.Atoi(cm); !(150 <= n && n <= 193) {
			glog.Warningf("invalid cms: %d", n)
			return false
		}
	}
	if in := strings.TrimSuffix(m[hgt], "in"); in != m[hgt] {
		if n, _ := strconv.Atoi(in); !(59 <= n && n <= 76) {
			glog.Warningf("invalid inches: %d", n)
			return false
		}
	}
	if !hex.MatchString(m[hcl]) {
		glog.Warningf("bad hex: %q", m[hcl])
		return false
	}
	if _, ok := eclSet[m[ecl]]; !ok {
		glog.Warningf("bad eye color: %q", m[ecl])
		return false // account for possible overwrites within the same passport?
	}
	if !prx.MatchString(m[pid]) {
		glog.Warningf("bad passport id: %q", m[pid])
		return false
	}

	return true
}

/*
func doMain(r io.Reader, w io.Writer) error {
	br := bufio.NewScanner(r)

	cur := map[string]bool{}
	valid := 0
	for br.Scan() {
		t := br.Text()
		if t == "" {
			if isValid(cur) {
				valid++
			}
			cur = map[string]bool{}
			continue
		}

		fields := strings.Split(t, " ")

		glog.Infof("got fields %v", fields)

		for _, f := range fields {
			p := strings.Split(f, ":")
			if len(p) != 2 {
				glog.Warningf("weird pair: %q", f)
			}

			glog.Infof("got pair: %v", p)

			cur[p[0]] = true
		}
	}
	if err := br.Err(); err != nil {
		return err
	}

	fmt.Fprintf(w, "VALID: %d\n", valid)
	return nil
}

func isValid(c map[string]bool) bool {
	glog.Infof("checking cur map: %v", c)
	for k := range required {
		if _, ok := c[k]; !ok {
			return false
		}
	}

	return true
}
*/
