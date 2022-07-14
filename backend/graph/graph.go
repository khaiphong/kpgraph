/*
 Ref: Graph data structure https://www.youtube.com/watch?v=JDP1OVgoa0Q
 https://www.youtube.com/watch?v=bSZ57h7GN2w put error checking
 Ref: 19 Graph Algorithms https://memgraph.com/blog/graph-algorithms-list
 Ref: Functional programming https://yourbasic.org/golang/your-basic-func/
*/
package main

import (
    "fmt"
    "sync"

//    "github.com/khaiphong/personadb/backend/pdb"
)

// GraphQL with EIP as entry point to Url of requested service. Persona Home
// is the user facebook clone
type Persona struct {
    Id      string `json:  "id"`
    Name    string
    HomePage    string
    Url         string
 //   Service     service.Service
 //   Services    [service.Service]
    Props       map[string]string
}

// each service must have its Url for internet access. The service Home is its
// facebook clone for further information e.g activities, relationships, places
type Service struct {
    Url         string
    ServiceName string
    Pdb         string
    HomePage    string
    Invoice     string
    Contract    string
 //   Users       [persona.Persona]
    Props       map[string]string
}

// Relationship is a directed graph from one Node to another Node
// It can be a Node for organization structure and for relationship
// cultivation and/or grooming to a certain position
type Relationship struct { // can be a Node to cultivate
    Id          string `json:  "id"`
    Name        string
    CurrentPost string
    TargetPost  string
    Props       map[string]string
}
type Organization struct { // can be a Node and have many services
    Id          string `json:  "id"`
    Name        string
    //   Services    [service.Service]
    Props       map[string]string
}
type Code struct {
    Id          string `json:  "id"`
    Name        string
    Props       map[string]string
}

// cultivating Inner Space of persona or organization properties Props. It is
// ML nodes from the grass roots for shaping persona or organization culture
type Fibonacci struct { // can be a patterned Node from persona
    Id          string `json:  "id"`
    Name        string
    Props       map[string]string
}
// Family cultivation and connecting tree
type Family struct { // is a Node
    Id          string `json:  "id"`
    Name        string
    Props       map[string]string
}
// CoOp Node for community GsLp ThankYou club and its vibrant economics
type GsLp struct { // is a node
    Id          string `json:  "id"`
    Name        string
    Props       map[string]string
}

// Graph represents an adjacency list graph
type Graph struct {
    nodes []*Node
    lock        sync.RWMutex
}

// Node represents a graph Node, run Golang 1.8 for generic Node.
// Specialized properties and processes applied, improved searches
// type Node struct { value GenericNode }
type GenericNode interface {
    Persona | Relationship | Organization | Code | Fibonacci | 
    Service | Family | GsLp
}
type Node struct {
    key int // unique key from Pdb
    adjacent    []*Node
    owners      []*Node
}

// AddNode adds a Node to the graph
func (g *Graph) AddNode(k int) {
    g.lock.Lock()
    g.nodes = append(g.nodes, &Node{key: k})
    g.lock.Unlock()
}    
 

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
    g.lock.Lock()
    // get Vertex
    fromNode := g.getNode(from)
    toNode := g.getNode(to)
    // check error
    // add edge
    fromNode.adjacent = append(fromNode.adjacent, toNode)
    g.lock.Unlock()
}


// getNode returns a pointer to the Node with a key integer
func (g *Graph) getNode(k int) *Node {
    for i, v := range g.nodes {
        if v.key == k {
            return g.nodes[i]
        }
    }
    return nil
}

// Print will print the adjacent list for each Node of the graph
func (g *Graph) Print() {
    for _,v := range g.nodes {
        fmt.Printf("\n Node %v: ", v.key)
        for _,v := range v.adjacent {
            fmt.Printf(" %v ", v.key)
        }
    }
}

func main(){

//    var personas map[string]Persona
//    var relationshipss map[string]Relationship
//    var organizations map[string]Organization
//    var codes map[string]Code
//    var fibonacciss map[string]Fibonacci
//    var services map[string]Service
//    var families map[string]Family
//    var gslps map[string]GsLp

    test := &Graph{}

    for i := 0; i < 5; i++ {
        test.AddNode(i)
    }

    test.AddEdge(1, 2)
    test.Print()
}
