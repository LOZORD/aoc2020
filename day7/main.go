package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	if err := doMain(); err != nil {
		glog.Errorf("error: %v", err)
	}
}

func doMain() error {
	br := bufio.NewScanner(os.Stdin)

	// A map of rule color => (map of contents => content count).
	// (rules["red"]["blue"] = 3) => "red bags contain 3 blue bags".
	rules := map[string]map[string]int{}

	for br.Scan() {
		t := br.Text()
		glog.V(2).Infof("scanned %q", t)
		// I probably could have used a fancy regex here, but splitting is simple enough to get what we want.
		s := strings.Split(t, " bags contain ")
		if len(s) != 2 {
			return fmt.Errorf("expected to be able split %q in half with ` bags contain `", t)
		}
		ruleColor := s[0]
		rules[ruleColor] = map[string]int{}
		// s[1] =~ "n x bags, 1 y bag, m z bags.".
		// or s[1] =~ "no other bags.".
		contents := strings.Split(strings.TrimSuffix(s[1], "."), ", ")
		if len(contents) < 1 {
			return fmt.Errorf("expected contents half %q to be non-empty", s[1])
		}
		if contents[0] == "no other bags" {
			continue // rules[ruleColor] should be empty.
		}
		for _, content := range contents {
			n, c, err := parseContent(content)
			if err != nil {
				return fmt.Errorf("failed to parse contents for %q => %q: %w", ruleColor, content, err)
			}
			rules[ruleColor][c] = n
		}
	}
	if err := br.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	glog.V(2).Infof("rules: %v", rules)

	// Part 1.
	n := findContainingCount(rules, map[string]bool{}, "shiny gold")
	fmt.Fprintf(os.Stdout, "CONTAINS_SHINY_GOLD = %d\n", n)

	// Part 2.
	m := findRequiredCount(rules, "shiny gold")
	fmt.Fprintf(os.Stdout, "CONTAINED_BY_SHINY_GOLD = %d\n", m)

	return nil
}

var contentRgx = regexp.MustCompile("^(?P<amount>[1-9]+) (?P<color>[a-z]+ [a-z]+) bags?$")

// parseContent parses something like "1 dark blue bag" => (1, "dark blue").
// This code is shamelessly stolen from https://stackoverflow.com/a/20751656.
func parseContent(s string) (int, string, error) {
	m := map[string]string{}
	match := contentRgx.FindStringSubmatch(s)
	if len(match) == 0 {
		return 0, "", fmt.Errorf("no match for %q", s)
	}
	for i, name := range contentRgx.SubexpNames() {
		if i != 0 && name != "" {
			m[name] = match[i]
		}
	}

	a, c := m["amount"], m["color"]
	if a == "" || c == "" {
		return 0, "", fmt.Errorf("no named match: amount=%q, color=%q", a, c)
	}

	n, err := strconv.Atoi(a)
	if err != nil {
		return 0, "", fmt.Errorf("failed to convert %q to int: %w", a, err)
	}

	return n, c, nil
}

func findContainingCount(rules map[string]map[string]int, seen map[string]bool, findColor string) int {
	sum := 0
	for ruleColor, contents := range rules {
		if _, ok := contents[findColor]; !ok {
			// Ignore this rule (for now) if it doesn't contain what we're looking for.
			continue
		}
		if seen[ruleColor] {
			// If we've already visited this rule, we don't need to change the sum.
			continue
		}
		// We've got a new bag (with color $ruleColor) that can contain at least one $findColor bag.
		seen[ruleColor] = true // We don't need to revisit this rule later on.
		sum = sum + 1 + findContainingCount(rules, seen, ruleColor)
	}
	return sum
}

func findRequiredCount(rules map[string]map[string]int, findColor string) int {
	sum := 0
	for color, amount := range rules[findColor] {
		// sum += amount + amount * find(...)
		sum += amount * (1 + findRequiredCount(rules, color))
	}
	return sum
}
