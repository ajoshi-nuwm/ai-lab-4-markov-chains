package main

import (
	"github.com/ajoshi-nuwm/ai-lab-4-markov-chains/node"
	"github.com/ajoshi-nuwm/ai-lab-4-markov-chains/utils"
	strings2 "strings"
	"fmt"
	"math/rand"
)

var (
	nodes []*node.Node
)

func main() {
	strings, _ := util.ReadFromFile("C:\\workspace\\bin\\text.txt")
	result := ""
	for _, s := range strings {
		result += s
	}

	splitted := strings2.Split(result, " ")
	for i := 0; i < len(splitted)-1; i++ {
		firstNode := findNode(splitted[i])
		secondNode := findNode(splitted[i+1])
		firstNode.AddNode(secondNode)
	}
	for _, n := range nodes {
		fmt.Println(n)
	}

	currentNode := nodes[rand.Intn(len(nodes))]
	generated := ""
	for i := 0; i < 100; i++ {
		currentNode = currentNode.GetNextNode()
		generated += fmt.Sprintf("%v ", currentNode.GetWord())
	}
	fmt.Println(generated)
}

func findNode(word string) *node.Node {
	for _, n := range nodes {
		if n.GetWord() == word {
			return n
		}
	}
	newNode := node.NewNode(word)
	nodes = append(nodes, newNode)
	return newNode
}
