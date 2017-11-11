package node

import (
	"math/rand"
	"fmt"
)

type Node struct {
	word  string
	nodes map[*Node]int
}

type Segment struct {
	node  *Node
	start float64
	end   float64
}

func NewNode(word string) *Node {
	nodes := make(map[*Node]int)
	return &Node{word, nodes}
}

func (node *Node) GetWord() string {
	return node.word
}

func (node *Node) contains(word string) (bool, *Node) {
	for n := range node.nodes {
		if n.word == word {
			return true, n
		}
	}
	return false, nil
}

func (node *Node) AddNode(other *Node) {
	if contains, found := node.contains(other.word); contains {
		node.nodes[found] += 1
	} else {
		node.nodes[other] = 1
	}
}

func (segment *Segment) isInSegment(value float64) bool {
	return value >= segment.start && value <= segment.end
}

func (node *Node) GetNextNode() *Node {
	r := rand.Float64()
	for _, s := range node.getSegments() {
		if s.isInSegment(r) {
			return s.node
		}
	}
	return nil
}

func (node *Node) getSegments() []Segment {
	segments := make([]Segment, 0)
	current := 0.0
	for n := range node.nodes {
		nodeProb := node.getNodeProb(n)
		segments = append(segments, Segment{n, current, current + nodeProb})
		current += nodeProb
	}
	return segments
}

func (node *Node) getNodeProb(currentNode *Node) float64 {
	var sum int
	for _, v := range node.nodes {
		sum += v
	}
	return float64(node.nodes[currentNode]) / float64(sum)
}

func (node Node) String() string {
	result := fmt.Sprintf("[%v] ", node.word)
	for n, v := range node.nodes {
		result += fmt.Sprintf("(%v - %v), ", n.word, v)
	}
	return result
}
