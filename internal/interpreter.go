package internal

import (
	"fmt"
	"lol/lib"
	"strconv"
)

func Interpret(tokenisedFile *Program) {
	stack := lib.NewStack[int]()
	programCounter := 0

	for tokenisedFile.Program[programCounter] != HALT {
		opcode := tokenisedFile.Program[programCounter]
		programCounter++

		if opcode == PUSH {
			num, err := strconv.Atoi(tokenisedFile.Program[programCounter])
			if err != nil {
				panic(err)
			}

			stack.Push(num)
			programCounter++

		} else if opcode == POP {
			stack.Pop()

		} else if opcode == ADD {
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(a + b)

		} else if opcode == SUB {
			a := stack.Pop()
			b := stack.Pop()
			stack.Push(b - a)

		} else if opcode == PRINT {
			str := tokenisedFile.Program[programCounter]
			programCounter++
			fmt.Println(str)

		} else if opcode == READ {
			var inp int
			_, err := fmt.Scan(&inp)
			if err != nil {
				panic(err)
			}
			stack.Push(inp)

		} else if opcode == JUMP_EQ_0 {
			if stack.Top() == 0 {
				label := tokenisedFile.Program[programCounter]
				programCounter = tokenisedFile.Labels[label]
			} else {
				programCounter++
			}

		} else if opcode == JUMP_GT_0 {
			if stack.Top() > 0 {
				label := tokenisedFile.Program[programCounter]
				programCounter = tokenisedFile.Labels[label]
			} else {
				programCounter++
			}

		} else {
			fmt.Println("opcode: ", opcode)
			panic("Unknown opcode")
		}
	}
}
