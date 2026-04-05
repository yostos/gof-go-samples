// Interpreter pattern: builds an AST for a mini turtle-graphics language and prints parsed tree.
// Uses Go interfaces and structs to represent the grammar nodes.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Node represents a parsed AST node.
type Node interface {
	String() string
}

// Context holds the tokenized input and provides parsing helpers.
type Context struct {
	tokens []string
	index  int
	token  string
}

func NewContext(text string) *Context {
	c := &Context{
		tokens: strings.Fields(text),
		index:  0,
	}
	c.nextToken()
	return c
}

func (c *Context) nextToken() string {
	if c.index < len(c.tokens) {
		c.token = c.tokens[c.index]
		c.index++
	} else {
		c.token = ""
	}
	return c.token
}

func (c *Context) currentToken() string {
	return c.token
}

func (c *Context) skipToken(expected string) error {
	if c.token == "" {
		return fmt.Errorf("'%s' expected, but no more tokens", expected)
	}
	if c.token != expected {
		return fmt.Errorf("'%s' expected, but '%s' found", expected, c.token)
	}
	c.nextToken()
	return nil
}

func (c *Context) currentNumber() (int, error) {
	if c.token == "" {
		return 0, fmt.Errorf("number expected, but no more tokens")
	}
	n, err := strconv.Atoi(c.token)
	if err != nil {
		return 0, fmt.Errorf("number expected, but '%s' found", c.token)
	}
	return n, nil
}

// ProgramNode: <program> ::= program <command list>
type ProgramNode struct {
	cmdList Node
}

func (n *ProgramNode) String() string {
	return fmt.Sprintf("[program %s]", n.cmdList)
}

func parseProgram(ctx *Context) (Node, error) {
	if err := ctx.skipToken("program"); err != nil {
		return nil, err
	}
	cmdList, err := parseCommandList(ctx)
	if err != nil {
		return nil, err
	}
	return &ProgramNode{cmdList: cmdList}, nil
}

// CommandListNode: <command list> ::= <command>* end
type CommandListNode struct {
	commands []Node
}

func (n *CommandListNode) String() string {
	parts := make([]string, len(n.commands))
	for i, c := range n.commands {
		parts[i] = c.String()
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func parseCommandList(ctx *Context) (Node, error) {
	var commands []Node
	for {
		if ctx.currentToken() == "" {
			return nil, fmt.Errorf("missing 'end'")
		}
		if ctx.currentToken() == "end" {
			ctx.skipToken("end")
			return &CommandListNode{commands: commands}, nil
		}
		cmd, err := parseCommand(ctx)
		if err != nil {
			return nil, err
		}
		commands = append(commands, cmd)
	}
}

// parseCommand: <command> ::= <repeat command> | <primitive command>
func parseCommand(ctx *Context) (Node, error) {
	if ctx.currentToken() == "repeat" {
		return parseRepeat(ctx)
	}
	return parsePrimitive(ctx)
}

// RepeatNode: <repeat command> ::= repeat <number> <command list>
type RepeatNode struct {
	count   int
	cmdList Node
}

func (n *RepeatNode) String() string {
	return fmt.Sprintf("[repeat %d %s]", n.count, n.cmdList)
}

func parseRepeat(ctx *Context) (Node, error) {
	if err := ctx.skipToken("repeat"); err != nil {
		return nil, err
	}
	count, err := ctx.currentNumber()
	if err != nil {
		return nil, err
	}
	ctx.nextToken()
	cmdList, err := parseCommandList(ctx)
	if err != nil {
		return nil, err
	}
	return &RepeatNode{count: count, cmdList: cmdList}, nil
}

// PrimitiveNode: <primitive command> ::= go | right | left
type PrimitiveNode struct {
	name string
}

func (n *PrimitiveNode) String() string {
	return n.name
}

func parsePrimitive(ctx *Context) (Node, error) {
	token := ctx.currentToken()
	if token != "go" && token != "right" && token != "left" {
		return nil, fmt.Errorf("unknown primitive command: '%s'", token)
	}
	ctx.skipToken(token)
	return &PrimitiveNode{name: token}, nil
}

func main() {
	// Read from program.txt if it exists, otherwise use embedded programs
	programs := []string{
		"program end",
		"program go end",
		"program go right go right go right go right end",
		"program repeat 4 go right end end",
		"program repeat 4 repeat 3 go right go left end right end end",
	}

	if f, err := os.Open("program.txt"); err == nil {
		defer f.Close()
		programs = nil
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			programs = append(programs, scanner.Text())
		}
	}

	for _, text := range programs {
		fmt.Printf("text = \"%s\"\n", text)
		ctx := NewContext(text)
		node, err := parseProgram(ctx)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("node = %s\n", node)
	}
}
