package main

import (
	"fmt"
	"testing"
)

/*
 * Complete the 'extraLongFactorials' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func TestExtraLongFactorials(m *testing.T) {
//    extraLongFactorials(1)
//    extraLongFactorials(2)
    extraLongFactorials(3)

    extraLongFactorials(25)
}



func TestBribe(m *testing.T) {
    //    extraLongFactorials(1)
    //    extraLongFactorials(2)
        q := []int32{1, 2, 5, 3 ,7, 8, 6, 4}
        // 1 + 2 + 4
        minimumBribes(q)
    
    }


// func TestSwaps(m *testing.T) {
//         //    extraLongFactorials(1)
//         //    extraLongFactorials(2)
//             q := []int32{1, 3, 5, 2 ,4, 6, 7}
//             // 1 + 2 + 4
// //      swaps :=      minimumSwaps(q)
// }
    

func TestHourGlass(m *testing.T) {
    //    extraLongFactorials(1)
    //    extraLongFactorials(2)
        q := [][]int32{
            {1, 1, 1, 0,0,0},
            {0, 1, 0, 0,0,0},
            {1, 1, 1, 0,0,0},
            {0, 0, 2, 4,4,0},
            {0, 0, 0, 2,0,0},
            {0, 0, 1, 2,4,0}}
        // 1 + 2 + 4
        max := hourglassSum(q)
        fmt.Print(max)
    }
