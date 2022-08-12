package main

import (
	"fmt"
	"math/big"
)

/*
 * Complete the 'extraLongFactorials' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func extraLongFactorials(n int32) {
    // Write your code here
    f := big.NewInt(int64(n))
    for i := int32(1) ; i < n ; i++ {
        prev := big.NewInt( int64(n - i) )
        f= f.Mul(f,prev)  
    }
    fmt.Printf("%v",f) 

}

//one person can bribe at most 2 but there is no limit in the amount ob bribes received
func minimumBribes(q []int32) {
    // Write your code here
    bribes := int32(0)
    tooChaotic := false 
    maxBribes :=int32(2)
    for currentIndex,currentPerson := range q{
        expectedPerson := int32(currentIndex+1)

          delta := currentPerson - expectedPerson
          if delta  >  maxBribes{
            tooChaotic = true
            break
          }

          // check how many people we do have
          beforeCurrentPersonIndex := currentPerson -1 -1 
          if beforeCurrentPersonIndex < 0{
            beforeCurrentPersonIndex = 0
          }

          for j:=beforeCurrentPersonIndex;j<int32(currentIndex);j++{
                if q[j] > currentPerson{
                        bribes++
                    }
           }
    
                  
    }
    if tooChaotic {
        fmt.Println("Too chaotic")
    }else{
        fmt.Println(bribes)
    }
}



// arr numbers are from 1 to n so 
func minimumSwaps(arr []int32) int32 {

    swaps := int32(0)
    // Holds Index +1  to Value (should be the same on sorted)
    a := make(map[int32]int32)
    // Holds Value to Current index  (should be the same on sorted)
    b := make(map[int32]int32)
    // just need to count how many are out of order from 1 to n-1
    for i:=0;i<len(arr) ; i++{
        nIndex := int32(i+1)
        a[nIndex] = arr[i]
        b[arr[i]] = nIndex
    }
    // check 
    for nIndex,nIndexValue := range a{
        
        if nIndex  != nIndexValue {
            // get the current index of the expected value
            y := b[nIndex]
            // swap current with that
            a[y] = nIndex
            b[nIndex] = y
            swaps++
        }
    }
    return swaps
}



/*
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */
func arrayManipulation(n int32, queries [][]int32) int64 {
    // Write your code here
    max:=int64(0)
    indexMap := make([]int64,n)
    for _,row := range queries {
        a:= row[0]
        b:= row[1]
        k:= int64(row[2]) 
        indexMap[a-1] += k
        if(b < n){
            indexMap[b] -= k            
        }
    }
    
    cum :=int64(0)
    for _,v :=range indexMap{
        cum = cum + v 
        if cum > max{
            max = cum
        } 
    }
    
    return max
}

func hourglassSum(arr [][]int32) int32 {
    // Write your code here
    max := int32(0)
    for iRowOffset:=0 ; iRowOffset < 4; iRowOffset++{
        for iColumnOffset:=0 ; iColumnOffset < 4; iColumnOffset++{
            sum := int32(0)
            for hourIndex:= 0 ; hourIndex < 3 ; hourIndex++{
                if hourIndex != 1 {
                    sum = sum + arr[iRowOffset + hourIndex][iColumnOffset] + arr[iRowOffset+ hourIndex][iColumnOffset+1] + arr[iRowOffset+ hourIndex][iColumnOffset+2] 

                }else{
                    sum = sum +  arr[iRowOffset+ hourIndex][iColumnOffset+1]                      
                }
                    
            }
            if sum > max {
                max = sum
            }
        }
    }
    return max
}


/*
 * Complete the 'countSwaps' function below.
 *
 * The function accepts INTEGER_ARRAY a as parameter.
 */

 func countBubbleSortSwaps(a []int32) {
    // Write your code here
    n := len(a)
    swaps := 0
    for i := 0; i < n; i++ {
        
        for  j := 0; j < n - 1; j++ {
            // Swap adjacent elements if they are in decreasing order
            if (a[j] > a[j + 1]) {
                temp := a[j + 1]
                a[j+1] = a[j]
                a[j] = temp
                swaps++
            }
        }        
    }
    fmt.Printf("Array is sorted in %v swaps.\n",swaps)
    fmt.Printf("First Element: %v\n",a[0])
    fmt.Printf("Last Element: %v\n",a[n-1])

}