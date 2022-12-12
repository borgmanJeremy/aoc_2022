package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type KeyValue struct {
	Key   string
	Value int
}

func sortMap(myMap map[string]int) map[string]int {
	pairs := make([]KeyValue, len(myMap))
	i := 0
	for k, v := range myMap {
		pairs[i] = KeyValue{k, v}
		i++
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	sortedMap := make(map[string]int)
	for _, pair := range pairs {
		sortedMap[pair.Key] = pair.Value
	}
	return sortedMap
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type Node struct {
	Name string
	Id   int
	Size int
	Type string
}

func nodeHash(n Node) int {
	return n.Id
}

func findParent(g graph.Graph[string, Node], currentNode string) string {
	am, _ := g.AdjacencyMap()
	for _, node := range am[currentNode] {
		prop := node.Properties.Attributes
		if prop["type"] == "parent" {
			return node.Target
		}
	}
	file, _ := os.Create("./mygraph.gv")
	_ = draw.DOT(g, file)

	log.Fatal("Could not find Parent")
	return ""
}

func directorySize(g graph.Graph[string, Node], currentNode string) int {
	am, _ := g.AdjacencyMap()
	size := 0
	for _, node := range am[currentNode] {
		vertex, _ := g.Vertex(node.Target)
		prop := node.Properties.Attributes

		if prop["type"] == "child" {
			if vertex.Type == "folder" {
				size += directorySize(g, node.Target)
			} else {
				size += vertex.Size
			}
		}
	}
	return size
}

func main() {
	instructions, err := os.Open("input/input.txt")
	checkError(err)
	defer instructions.Close()

	id := 0
	g := graph.New(nodeHash, graph.Directed(), graph.Acyclic())
	g.AddVertex(Node{"/", id, 0, "folder"})
	id += 1

	currentNodeId := 0
	dirList := []string{}
	dirList = append(dirList, "/")
	scanner := bufio.NewScanner(instructions)

	count := 0
	// Buliding Tree
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(count, ":", line)

		count += 1
		if line[0] == '$' {
			tokens := strings.Split(line, " ")
			if tokens[1] == "ls" {

			} else if tokens[1] == "cd" {
				if tokens[2] == ".." {
					parentId := findParent(g, currentNodeId)
					currentNodeId = parentId
				} else {
					if tokens[2] != "/" {
						currentNode = tokens[2]
					}
				}
			} else {
				log.Fatal("Invalid command")
			}
		} else {
			// handle files and dirs
			tokens := strings.Split(line, " ")
			if tokens[0] == "dir" {
				if !contains(dirList, tokens[1]) {
					dirList = append(dirList, tokens[1])

					g.AddVertex(Node{tokens[1], id, 0, "folder"})
					id += 1
				}

				g.AddEdge(currentNode, tokens[1], graph.EdgeAttribute("type", "child"))
				g.AddEdge(tokens[1], currentNode, graph.EdgeAttribute("type", "parent"))
			} else {
				weight, _ := strconv.Atoi(tokens[0])
				g.AddVertex(Node{tokens[1], weight, "file"})
				g.AddEdge(currentNode, tokens[1], graph.EdgeAttribute("type", "child"))
				g.AddEdge(tokens[1], currentNode, graph.EdgeAttribute("type", "parent"))
			}
		}
	}

	dirSize := make(map[string]int)
	for _, dir := range dirList {
		dirSize[dir] = directorySize(g, dir)
	}
	part_1 := 0
	for _, dir := range dirList {
		if dirSize[dir] <= 100000 {
			fmt.Println(dir, dirSize[dir])
			part_1 += dirSize[dir]
		}
	}

	fmt.Println("Part 1: ", part_1)

	//file, _ := os.Create("./mygraph.gv")
	//_ = draw.DOT(g, file)

}
