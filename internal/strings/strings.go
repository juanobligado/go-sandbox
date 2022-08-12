package main

func makeAnagram(a string, b string) int32 {
	// Write your code here
	aMap := make(map[rune]int32)
	bMap := make(map[rune]int32)
	keys := make(map[rune]bool)
	for _, v := range a {
		if !keys[v] {
			keys[v] = true
		}
		aMap[v] = aMap[v] + 1
	}

	for _, v := range b {

		if !keys[v] {
			keys[v] = true
		}
		bMap[v] = bMap[v] + 1
	}

	deltaTotal := int32(0)
	for k := range keys {

		aV, _ := aMap[k]
		bV, _ := bMap[k]
		if aV > bV {
			deltaTotal = deltaTotal + (aV - bV)
		} else if aV < bV {
			deltaTotal = deltaTotal + (bV - aV)
		}

	}
	return deltaTotal

}