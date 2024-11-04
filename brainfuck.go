package brainfuck

import (
	"fmt"
	"strings"
)

const (
	VERSION_MAJOR = 0
	VERSION_MINOR = 1
	VERSION_PATCH = 0
	CELL_SIZE     = 30000
)

type CellType = uint8

type Opcode uint8

const (
	OpNop   Opcode = iota
	OpPlus         // +
	OpMinus        // -
	OpGT           // >
	OpLT           // <
	OpLB           // [
	OpRB           // ]
	OpDot          // .
	OpComma        // ,
)

type Instruction struct {
	opcode         Opcode
	cellValue      CellType
	cellIndexValue int
	jumpTarget     *Instruction
	next           *Instruction
}

type Context struct {
	cell      []CellType
	cellIndex int
}

func compress(source string) (result string) {
	for pc := 0; pc < len(source); pc++ {
		c := source[pc]
		if strings.IndexByte("+-><[].,", c) != -1 {
			result += string(c)
		}
	}
	return
}

func prepare(source string) (*Instruction, error) {
	compressedSource := compress(source)
	stack := make([]*Instruction, 1024)
	sp := 0

	head := &Instruction{opcode: OpNop}
	current := head
	var next *Instruction
	length := len(compressedSource)

	for pc := 0; pc < length; {
		c := compressedSource[pc]
		switch c {
		case '+':
			next = &Instruction{
				opcode:    OpPlus,
				cellValue: 0,
			}
			for pc < length {
				if compressedSource[pc] == '+' {
					next.cellValue++
					pc++
				} else {
					break
				}
			}
			pc--
		case '-':
			next = &Instruction{
				opcode:    OpMinus,
				cellValue: 0,
			}
			for pc < length {
				if compressedSource[pc] == '-' {
					next.cellValue++
					pc++
				} else {
					break
				}
			}
			pc--
		case '>':
			next = &Instruction{
				opcode:         OpGT,
				cellIndexValue: 0,
			}
			for pc < length {
				if compressedSource[pc] == '>' {
					next.cellIndexValue++
					pc++
				} else {
					break
				}
			}
			pc--
		case '<':
			next = &Instruction{
				opcode:         OpLT,
				cellIndexValue: 0,
			}
			for pc < length {
				if compressedSource[pc] == '<' {
					next.cellIndexValue++
					pc++
				} else {
					break
				}
			}
			pc--
		case '[':
			next = &Instruction{
				opcode:     OpLB,
				jumpTarget: nil,
			}
			stack[sp] = next
			sp++
		case ']':
			if sp <= 0 {
				return nil, fmt.Errorf("unmatched ']'")
			}
			next = &Instruction{
				opcode:     OpRB,
				jumpTarget: nil,
			}
			sp--
			target := stack[sp]
			next.jumpTarget = target
			target.jumpTarget = next
		case '.':
			next = &Instruction{
				opcode: OpDot,
			}
		case ',':
			next = &Instruction{
				opcode: OpComma,
			}
		}
		pc++
		current.next = next
		current = current.next
	}
	if sp > 0 {
		return nil, fmt.Errorf("unmatched '['")
	}
	return head, nil
}

func NewContext() *Context {
	return &Context{
		cell:      make([]CellType, CELL_SIZE),
		cellIndex: 0,
	}
}

func (c *Context) Execute(source string) error {
	head, err := prepare(source)

	if err != nil {
		return err
	}

	for instruction := head; instruction != nil; {
		switch instruction.opcode {
		case OpNop:
		case OpPlus:
			c.cell[c.cellIndex] += instruction.cellValue
		case OpMinus:
			c.cell[c.cellIndex] -= instruction.cellValue
		case OpGT:
			c.cellIndex += instruction.cellIndexValue
		case OpLT:
			c.cellIndex -= instruction.cellIndexValue
		case OpLB:
			if c.cell[c.cellIndex] == 0 {
				instruction = instruction.jumpTarget
			}
		case OpRB:
			if c.cell[c.cellIndex] != 0 {
				instruction = instruction.jumpTarget
			}
		case OpDot:
			fmt.Printf("%c", c.cell[c.cellIndex])
		case OpComma:
			fmt.Scanf("%c", &(c.cell[c.cellIndex]))
		}
		if instruction != nil {
			instruction = instruction.next
		}
	}

	return nil
}
