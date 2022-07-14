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

// each service must have it Url for internet access. The service Home is its
// facebook clone for further information e.g Activities, Relationships, Places
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

// Relationship is a directed graph from one Vertex to another Vertex
// It can be a Vertex for organization structure and for relationship
// cultivation and/or grooming to a certain position
type Relationship struct { // can be a node to cultivate
    Id          string `json:  "id"`
    Name        string
    CurrentPost string
    TargetPost  string
    Props       map[string]string
}
type Organization struct { // can be a node and have many services
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
// ML vertices from the grass roots for shaping persona or organization culture
type Fibonacci struct { // can be a patterned node from persona
    Id          string `json:  "id"`
    Name        string
    Props       map[string]string
}
// Family cultivation and connecting tree
type Family struct { // is anode
    Id          string `json:  "id"`
    Name        string
    Props       map[string]string
}
// CoOp node for community GsLp ThankYou club and its vibrant economics
type GsLp struct { // is a node
    Id          string `json:  "id"`
    Name        string
    Props       map[string]string
}

// Graph represents an adjacency list graph
type Graph struct {
    vertices []*Vertex
    lock        sync.RWMutex
}

// Vertex represents a graph vertex, run Golang 1.8 for generic Vertex.
// specialized properties and processes applied, improved searches
// type Vertex struct { value GenericVertex }
type GenericVertex interface {
    Persona | Relationship | Organization | Code | Fibonacci | 
    Service | Family | GsLp
}
type Vertex struct {
    key int // unique key from Pdb
    adjacent    []*Vertex
    owners      []*Vertex
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(k int) {
    g.lock.Lock()
    g.vertices = append(g.vertices, &Vertex{key: k})
    g.lock.Unlock()
}    
 

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
    g.lock.Lock()
    // get Vertex
    fromVertex := g.getVertex(from)
    toVertex := g.getVertex(to)
    // check error
    // add edge
    fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
    g.lock.Unlock()
}


// getVertex returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
    for i, v := range g.vertices {
        if v.key == k {
            return g.vertices[i]
        }
    }
    return nil
}

// Print will print the adjacent list for each vertiex of the graph
func (g *Graph) Print() {
    for _,v := range g.vertices {
        fmt.Printf("\nVertex %v: ", v.key)
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
        test.AddVertex(i)
    }

    test.AddEdge(1, 2)
    test.Print()
}
