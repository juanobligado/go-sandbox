package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'whatFlavors' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY cost
 *  2. INTEGER money
 */

func whatFlavors(cost []int32, money int32) {
    // Write your code here
    CostToIndex := make(map[int32][]int)
     
    for k,v := range cost{
        CostToIndex[v] = append(CostToIndex[v],k+1)          
    }
    maxCost := int32(0)
    index1,index2 :=int(-1) ,int(-1)
    currentCost := int32(0)
    for v1,v1Items := range CostToIndex {
        if len(v1Items) >1{           
           currentCost = 2*v1
           if currentCost <= money && currentCost > maxCost{
               maxCost = currentCost
               index1 = v1Items[0]
               index2 = v1Items[1]               
           }   
        }
        if maxCost == money{
            break
        }
                
        for v2,v2Items := range CostToIndex {
            if(v1 == v2){
                continue                    
            }else{
                currentCost = v1 + v2
                if currentCost <= money && currentCost > maxCost{
                    maxCost = currentCost
                    index1 = v1Items[0]
                    index2 = v2Items[0] 
                }
            }
            if maxCost == money{
                break
            }
            
        }
    }
    if index1 < index2{
        fmt.Printf("%v %v\n",index1,index2)        
    }else{
        fmt.Printf("%v %v\n",index2,index1)       
    }
     
    
     

}



/*
 * Complete the 'minimumAbsoluteDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func minimumAbsoluteDifference(arr []int32) int32 {
    // Write your code here
    min := int32(math.MaxInt32)
    hashMap := make(map[int32]int32)    
    sorted := make([]int32,0)
    for i := 0 ; i < len(arr)  ; i++{        
        _,exists := hashMap[arr[i]]
        if exists{
            return 0
        }
        sorted = appendWithOrder(sorted,arr[i])
    }
    
    for i := 0; i< len(sorted) -1 ; i++ {
        c := abs(sorted[i],sorted[i+1])
        if c < min{
            min = c
        }
    }
    return min
}

func appendWithOrder(arr []int32,v int32) []int32{
    
    if len(arr) == 0 || v > arr[len(arr)-1] {
        return append(arr,v)
    }
    
    if v < arr[0] {
        return append( []int32{v} ,arr...)
    }
    
    // find first index such as arr[i] < v && arr[i+1] > v
    left := int32(0)
    right := int32(len(arr) -1)    
    index:= int32(-1)
    for index < 0 {
        mid :=  (left + right) / 2     
        if v > arr[mid]{
          left = mid
          // found it break the loop
          if v < arr[mid+1]{
              index = mid
          }   
        }else if v < arr[mid] {
            right = mid
            // found it break the loop
            if v > arr[mid -1]{
                index = mid
            } 
        }
    }
    // shift right from 
    arr = append(arr[:index+1],arr[index:]...)
    arr[index+1] = v
    return arr
}



func abs (a int32, b int32)int32{
    
    c := a-b
    if c < 0{
        c = -c
    }
    return c
}


func fizzBuzz(n int32) {
    // Write your code here
    for i:=int32(1) ; i <= n ; i++ {
        is3 := (i % 3 == 0) 
        is5 := (i % 5 == 0) 
        if !is3 && !is5 {
            fmt.Println(i)
            continue
        }
        
        if is3 && i >= 3{
            fmt.Print("Fizz")
        } 
        
        if is5 && i >= 5{
            fmt.Print("Buzz")
        }         
        fmt.Print("\n")
    }

    

}


func mcd( k int32)int32{
    n := int32(1)
    for mcd := int32(2) ; mcd <=k ; mcd++{
        if k % mcd == 0{
            return mcd           
        }
    }
    return n
}


func maxSetSize(riceBags []int32) int32 {
    // Write your code here
    bins := make(map[int32][]int32)
    for _,bagSize:=range riceBags{
        b := mcd(bagSize)
        bins[b] = append(bins[b],bagSize)            
    }
    max := int32(-1)
    for _,v := range bins {
        if int32(len(v)) > max{
            m := calcmax(v)
            if m > max && m > 1{
                max = m
            }
    
        }
    }
    return max
    

}



func calcmax(riceBags []int32) int32 {
    // Write your code here
    if len(riceBags) < 2{
        return int32(-1)
        
    }
    dict := make(map[int32]bool)
    for i:=0 ; i < len(riceBags) ;i++{
        dict[riceBags[i]]=true
    }

    base := mcd(riceBags[0])
    n :=base 
    count := int32(0)

    for i:=0;i<len(riceBags);i++{
        
        if !dict[n]{
            return count
        }else{
            count++
        }  
        n  = n*n           
    }
    return count
    

}