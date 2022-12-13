package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"example.com/day_7/pathMap"
	"github.com/dominikbraun/graph"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type NodeType int

const (
	Folder NodeType = iota
	File
)

type Node struct {
	Id   int
	Size int
	Type NodeType
}

func nodeHash(n Node) int {
	return n.Id
}

func pathStackToString(pathStack []string) string {
	path := ""
	if len(pathStack) == 1 {
		return "/"
	}
	for idx, dir := range pathStack {
		if idx != 0 {
			path += "/" + dir
		}
	}
	return path
}

func directorySize(g graph.Graph[int, Node], currentNode int) int {
	am, _ := g.AdjacencyMap()
	size := 0
	for _, node := range am[currentNode] {
		vertex, _ := g.Vertex(node.Target)
		prop := node.Properties.Attributes

		if prop["type"] == "child" {
			if vertex.Type == Folder {
				size += directorySize(g, node.Target)
			} else {
				size += vertex.Size
			}
		}
	}
	return size
}

func part_1(g graph.Graph[int, Node], pathMap pathMap.PathMap) {
	totalSize := 0
	for _, id := range pathMap.PMap {
		vertex, _ := g.Vertex(id)
		if vertex.Type == Folder {
			size := directorySize(g, id)
			if size <= 100000 {
				totalSize += size
			}
		}
	}
	fmt.Println("Part 1: ", totalSize)
}

func main() {
	instructions, err := os.Open("input/input.txt")
	checkError(err)
	defer instructions.Close()

	scanner := bufio.NewScanner(instructions)
	pathStack := []string{}
	pathMap := pathMap.New()
	pathMap.AddKey("/")
	g := graph.New(nodeHash, graph.Directed(), graph.Acyclic())
	g.AddVertex(Node{
		Id:   0,
		Size: 0,
		Type: Folder,
	})

	currentId := 0
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(count, ": ", line)
		count++
		if line[0] == '$' {
			tokens := strings.Split(line, " ")
			if tokens[1] == "ls" {
				// Do nothing
			} else if tokens[1] == "cd" {
				if tokens[2] == ".." {
					pathStack = pathStack[:len(pathStack)-1]
					currentId = pathMap.PMap[pathStackToString(pathStack)]
				} else {
					pathStack = append(pathStack, tokens[2])
					path := pathStackToString(pathStack)
					pathMap.AddKey(path)
					currentId = pathMap.PMap[pathStackToString(pathStack)]
				}
			} else {
				log.Fatal("Invalid command")
			}
		} else {
			// handle files and dirs
			tokens := strings.Split(line, " ")
			path := pathStackToString(append(pathStack, tokens[1]))

			if tokens[0] == "dir" {
				if _, ok := pathMap.PMap[path]; !ok {
					pathMap.AddKey(path)
					g.AddVertex(Node{
						Id:   pathMap.PMap[path],
						Size: 0,
						Type: Folder,
					})
					g.AddEdge(currentId, pathMap.PMap[path], graph.EdgeAttribute("type", "child"))
					g.AddEdge(pathMap.PMap[path], currentId, graph.EdgeAttribute("type", "parent"))
				}
			} else {
				// file
				size, _ := strconv.Atoi(tokens[0])
				pathMap.AddKey(path)
				g.AddVertex(Node{
					Id:   pathMap.PMap[path],
					Size: size,
					Type: File,
				})
				g.AddEdge(currentId, pathMap.PMap[path], graph.EdgeAttribute("type", "child"))
				g.AddEdge(pathMap.PMap[path], currentId, graph.EdgeAttribute("type", "parent"))
			}
		}
	}

	// part 1
	part_1(g, pathMap)
}
