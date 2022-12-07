package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	root := parse(os.Stdin)
	fmt.Println(root)
}

type CMD struct {
	Instruction string
	Results     []string
}

func parse(r io.Reader) Node {
	s := bufio.NewScanner(r)

	var cmds []CMD

	cmd := CMD{}

	for s.Scan() {
		t := strings.TrimSpace(s.Text())

		if strings.HasPrefix(t, "$ ") {
			if cmd.Instruction != "" {
				cmds = append(cmds, cmd)
			}
			cmd = CMD{}

			cmd.Instruction = strings.TrimPrefix(t, "$ ")
		} else {
			cmd.Results = append(cmd.Results, t)
		}
	}

	cmds = append(cmds, cmd)

	tree := buildFileTree(cmds)

	return tree
}

type NodeType string

const (
	FileType      NodeType = "file"
	DirectoryType NodeType = "dir"
)

type Node interface {
	Name() string
	Children() []Node
	Size() int
	Type() NodeType
	Parent() Node
}

type File struct {
	name   string
	size   int
	parent Node
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Children() []Node {
	return nil
}

func (f *File) Size() int {
	return f.size
}

func (f *File) Type() NodeType {
	return FileType
}

func (f *File) Parent() Node {
	return f.parent
}

type Directory struct {
	name     string
	children []Node
	parent   Node
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Children() []Node {
	return d.children
}

func (d *Directory) Size() int {
	sum := 0
	for _, c := range d.Children() {
		sum += c.Size()
	}
	return sum
}

func (d *Directory) Type() NodeType {
	return DirectoryType
}

func (f *Directory) Parent() Node {
	return f.parent
}

func buildFileTree(cmds []CMD) Node {
	root := &Directory{
		name:   "/",
		parent: nil,
	}

	var context Node

	context = root

	for _, cmd := range cmds {
		instructionParts := splitInstruction(cmd.Instruction)
		directive := instructionParts[0]

		switch directive {
		case "cd":
			location := instructionParts[1]

			switch location {
			case "/":
				context = root
			case "..":
				context = context.Parent()
			default:
				for _, child := range context.Children() {
					if child.Name() == location {
						context = child
					}
				}
			}

		case "ls":
			dir, ok := context.(*Directory)
			if !ok {
				panic("type cast to dir failed")
			}
			dir.children = nodesFromResults(context, cmd.Results)

		default:
			panic(fmt.Sprintf("unknown directive for instruction", cmd.Instruction))
		}
	}

	return root

}

func splitInstruction(s string) []string {
	return strings.Split(s, " ")
}

func nodesFromResults(ctx Node, results []string) []Node {
	var nodes []Node
	for _, r := range results {
		if strings.HasPrefix(r, "dir ") {
			nodes = append(nodes, &Directory{
				name:     strings.TrimPrefix(r, "dir "),
				children: nil,
				parent:   ctx,
			})
		} else {
			parts := strings.Split(r, " ")
			size, file := mustParseInt(parts[0]), parts[1]

			nodes = append(nodes, &File{
				name:   file,
				size:   size,
				parent: ctx,
			})
		}
	}
	return nodes
}

func mustParseInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse %s", n))
	}

	return int(n)
}
