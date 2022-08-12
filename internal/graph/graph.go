package graph

import (
	"fmt"
)

type Vertex[T comparable ] struct{
	key T
	weight int32 
	adjacent []*Vertex[T]
}

//type Int32Graph Graph[int32]
//type Int32Vertex Vertex[int32]

func  contains[T comparable ] (g []*Vertex[T],k T) bool{
	for _,v := range g{
		if(k == v.key){
			return true
		}
	}
	return false
}

type Graph[T comparable ] struct {
	vertices []*Vertex[T]
}

func   AddVertex[T comparable](g *Graph[T],k T){
	if contains(g.vertices,k){
		err := fmt.Errorf("Vertex %v not added because its an existing key",k)
		fmt.Println(err.Error())
	}else{
		g.vertices = append(g.vertices, &Vertex[T]{key:k ,weight: int32(1)} )
	}
}

func  GetVertex[T comparable](g *Graph[T],k T) *Vertex[T]{
	for i,v := range g.vertices{
		if k == v.key {
			return g.vertices[i]
		}
	}
	return nil
}

func  AddEdge[T comparable](g *Graph[T], from,to T){
	// get vertex
	fromVertex := GetVertex(g,from)
	toVertex := GetVertex(g,to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid Edge (%v-->%v)",from,to)
		fmt.Println(err.Error())
		return
	}

	if contains(fromVertex.adjacent,to)   {
		err := fmt.Errorf("Existing edge (%v-->%v)",from,to)
		fmt.Println(err.Error())
		return
	}
	//Double Link Vertexes
	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	toVertex.adjacent = append(toVertex.adjacent, fromVertex)
}

func  Print[T comparable](g *Graph[T]){
	for _,v := range g.vertices{
		fmt.Printf("\n Vertex %v : ",v.key)
		for _,a := range v.adjacent{
			fmt.Printf("%v",a.key)
		}
	}
}

func  DFS[T comparable](v *Vertex[T] , visited map[T]bool ) {

	
	visited[v.key] = true
	fmt.Printf("Visited %v",v.key)
	for _,a := range v.adjacent{
		if visited[a.key]{
			continue
		}else{
			
			DFS(a,visited)
		}
			
		 
	}
	
}

func  BFS[T comparable](starting *Vertex[T] ,  ) {

	queue := make([]*Vertex[T],0)
	visited := make(map[T]bool)
	visited[starting.key] = true
	queue = append(queue, starting)
	
	for len(queue) > 0  {
		currentVertex := queue[0]
		queue = queue[1:]	
		fmt.Printf("\n Visited %v",currentVertex.key)
		for _,a  := range currentVertex.adjacent {
			if !visited[a.key] {
				visited[a.key] = true
				queue = append(queue, a)
			}
		}
	}
	
}

func Dijkstra[T comparable](start,end *Vertex[T] ) []T{
	bestPriceTable := make(map[T]int32)
	bestPreviousVertex := make(map[T]T)
	
	
	visited := make(map[T]bool)
	unvisited := make([]*Vertex[T],0)
	
	currentCity := start
	// Iterate over all vertex and get
	// 1- min
	bestRouteCost[start.key] = 0

	for currentCity != nil  {
		
		// pop next vertex
		visited[currentCity.key] = true
		visit_queue = visit_queue[1:]

		// iterate over adjacent vertix
		for _,nextCity := range currentCity.adjacent {
			
			if !visited[nextCity.key]{
				unvisited = append(unvisited, nextCity)
			}
			
			priceThroughCurrent  := bestPriceTable[currentCity.key] + nextCity.weight
			priceToCurrent :=  currentCost + nextCity.weight
			if !bestPriceTable[nextCity.key] || priceThroughCurrent < bestPriceTable[nextCity.key]{
				bestPriceTable[nextCity.key] = priceThroughCurrent
				bestPreviousVertex[nextCity.key] = currentCity.key
			}  
		}

		// Just visit the one with 
		if len(unvisited) > 0{
			
			for i,v := range unvisited{
				if i == 0 || bestPriceTable[v.key] <
			}
			
		}else{
			currentCity = nil
		}
			
	}


	return []
}



