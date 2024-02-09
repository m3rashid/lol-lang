package internal

import (
	"errors"
	"strings"
)

const (
	PUSH      = "PUSH"
	POP       = "POP"
	ADD       = "ADD"
	SUB       = "SUB"
	PRINT     = "PRINT"
	READ      = "READ"
	JUMP_EQ_0 = "JUMP.EQ.0"
	JUMP_GT_0 = "JUMP.GT.0"
	HALT      = "HALT"
)

type Program struct {
	Labels  map[string]int
	Program []string
}

func newProgram() Program {
	return Program{
		Labels:  map[string]int{},
		Program: []string{},
	}
}

func TokenizeProgram(lines *[]string) (Program, error) {
	tokenCount := 0
	program := newProgram()

	for _, line := range *lines {
		parts := strings.Split(line, " ")
		opcode := parts[0]

		if opcode == "" {
			continue
		}

		if opcode[len(opcode)-1] == ':' {
			program.Labels[opcode[:len(opcode)-1]] = tokenCount
			continue
		}

		program.Program = append(program.Program, opcode)
		tokenCount++

		if opcode == PUSH || opcode == JUMP_EQ_0 || opcode == JUMP_GT_0 {
			program.Program = append(program.Program, parts[1])
			tokenCount++
		} else if opcode == PRINT {
			joinStr := strings.Join(parts[1:], " ")

			if joinStr[0] != '"' || joinStr[len(joinStr)-1] != '"' {
				return program, errors.New("strings must be enclosed in double quotes")
			}

			program.Program = append(program.Program, joinStr[1:len(joinStr)-1])

			tokenCount++
		}

	}

	return program, nil
}
