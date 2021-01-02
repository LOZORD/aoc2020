package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	if err := doMain(); err != nil {
	}
}

type cmd int

const (
	nop cmd = iota + 1
	acc
	jmp
)

type instruction struct {
	cmd cmd
	amt int
}

type program struct {
	instructions   []instruction
	accumulator    int
	programCounter int

	seen map[int]bool
}

func (p *program) execute() {
	for {
		if p.programCounter == len(p.instructions) {
			glog.Infof("program successfully terminated at pc=%d, acc=%d", p.programCounter, p.accumulator)
			break
		}

		if p.seen[p.programCounter] {
			glog.Infof("found cycle at instruction %d (acc=%d)", p.programCounter, p.accumulator)
			break
		}
		p.seen[p.programCounter] = true

		inst := p.instructions[p.programCounter]
		switch inst.cmd {
		case nop:
			p.programCounter++
			continue
		case acc:
			p.accumulator += inst.amt
			p.programCounter++
			continue
		case jmp:
			p.programCounter += inst.amt
			continue
		default:
			panic(fmt.Errorf("unknown command: %v", inst.cmd))
		}
	}
}

func doMain() error {
	var is []instruction
	br := bufio.NewScanner(os.Stdin)
	for br.Scan() {
		t := br.Text()
		i, err := parseInstruction(t)
		if err != nil {
			return fmt.Errorf("failed to parse instruction %q: %v", t, err)
		}
		is = append(is, i)
	}

	// Part 1.
	{
		p := &program{
			instructions: is,
			seen:         map[int]bool{},
		}

		p.execute()
		fmt.Printf("LOOP_DETECTED: PCTR = %d ; ACCM = %d\n", p.programCounter, p.accumulator)
	}

	// Part 2.
	{
		var tacc int
		icpy := make([]instruction, len(is))
		for ind, ist := range is {
			var newCmd cmd
			switch ist.cmd {
			case acc:
				continue
			case nop:
				newCmd = jmp
			case jmp:
				newCmd = nop
			}

			copy(icpy, is)
			icpy[ind] = instruction{cmd: newCmd, amt: ist.amt}
			p := &program{
				instructions: icpy,
				seen:         map[int]bool{},
			}
			p.execute()
			if p.programCounter == len(p.instructions) {
				tacc = p.accumulator
				break
			}
		}
		fmt.Printf("SUCCESSFUL_TERMINATION_ACC = %d\n", tacc)
	}

	return nil
}

func parseInstruction(s string) (instruction, error) {
	spl := strings.Split(s, " ")
	if len(spl) != 2 {
		return instruction{}, fmt.Errorf("got %d-part split on %q, wanted 2", len(spl), s)
	}

	var c cmd
	switch spl[0] {
	case "nop":
		c = nop
	case "acc":
		c = acc
	case "jmp":
		c = jmp
	default:
		return instruction{}, fmt.Errorf("unknown command %q", spl[0])
	}

	a, err := strconv.Atoi(spl[1])
	if err != nil {
		return instruction{}, fmt.Errorf("failed to parse amount from %q", spl[1])
	}

	return instruction{cmd: c, amt: a}, nil
}
