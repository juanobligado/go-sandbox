package graph

import (
	"testing"
)

func TestDFS(m *testing.T){
	testGraph := &Graph[string]{}
	
	AddVertex(testGraph,"Alice")
	AddVertex(testGraph,"Bob")
	AddVertex(testGraph,"Fred")
	AddVertex(testGraph,"Helen")
	AddVertex(testGraph,"Candy")
	AddVertex(testGraph,"Derek")
	AddVertex(testGraph,"Elaine")
	AddVertex(testGraph,"Gina")
	AddVertex(testGraph,"Irena")
	AddEdge(testGraph,"Alice","Bob")
	AddEdge(testGraph,"Alice","Candy")
	AddEdge(testGraph,"Alice","Derek")
	AddEdge(testGraph,"Alice","Elaine")
	AddEdge(testGraph,"Bob","Fred")
	AddEdge(testGraph,"Fred","Helen")
	AddEdge(testGraph,"Candy","Helen")
	AddEdge(testGraph,"Derek","Gina")
	AddEdge(testGraph,"Derek","Elaine")
	AddEdge(testGraph,"Gina","Irena")


	Print(testGraph)
	visited := make(map[string]bool)
	DFS(GetVertex(testGraph,"Alice"),visited)
}

func TestBFS(m *testing.T){
	testGraph := &Graph[string]{}
	
	AddVertex(testGraph,"Alice")
	AddVertex(testGraph,"Bob")
	AddVertex(testGraph,"Fred")
	AddVertex(testGraph,"Helen")
	AddVertex(testGraph,"Candy")
	AddVertex(testGraph,"Derek")
	AddVertex(testGraph,"Elaine")
	AddVertex(testGraph,"Gina")
	AddVertex(testGraph,"Irena")

	AddEdge(testGraph,"Alice","Bob")
	AddEdge(testGraph,"Alice","Candy")
	AddEdge(testGraph,"Alice","Derek")
	AddEdge(testGraph,"Alice","Elaine")

	AddEdge(testGraph,"Bob","Fred")
	AddEdge(testGraph,"Fred","Helen")
	AddEdge(testGraph,"Candy","Helen")
	AddEdge(testGraph,"Derek","Gina")
	AddEdge(testGraph,"Derek","Elaine")
	AddEdge(testGraph,"Gina","Irena")


	Print(testGraph)

	BFS(GetVertex(testGraph,"Alice"))
}


func TestClouds(m *testing.T){
	testGraph := &Graph[int32]{}
	

	c := []int32{0,1,0,0,0,1,0}

	for i,v := range c{
		AddVertex(testGraph,i)
	}
	lastIndex = len(c) -1
	for i:= 0;i< lastIndex ; i++{
		if c[i +1] == 0{
			AddEdge(testGraph,i,i+1)
		}
		if i+2 < lastIndex && c[i +2] == 0{
			AddEdge(testGraph,i,i+2)
		}
	}


	Print(testGraph)

}

