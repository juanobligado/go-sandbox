package main

func sort(buffer []int32) []int32 {
	for l := 0; l < len(buffer)-1; l++ {
		for r := l; r < len(buffer); r++ {
			if buffer[l] > buffer[r] {
				temp := buffer[l]
				buffer[l] = buffer[r]
				buffer[r] = temp
			}
		}
	}
	return buffer
}

func activityNotifications(expenditure []int32, d int32) int32 {
	// Write your code here
	totalExp := int32(len(expenditure))
	if totalExp < d {
		return int32(0)
	}
	notices := int32(0)
	isOdd := d%2 == 1

	//sort buffer
	buffer := make([]int32, 0)
	buffer = append(buffer, expenditure[0:d]...)
	buffer = sort(buffer)

	for i := int32(0); i < (totalExp - d); i++ {
		mid := int32(len(buffer) / 2)
		median := buffer[mid]
		if !isOdd {
			median = (median + buffer[mid+1]) / 2
		}
		nextItem := expenditure[d+i]
		if nextItem >= 2*median {
			notices++
		}

		inserted := false
		for ib := 1; ib < len(buffer); ib++ {
			buffer[ib-1] = buffer[ib]
			if !inserted {
				if nextItem < buffer[ib] {
					buffer[ib-1] = nextItem
					break
				} else if ib == len(buffer)-1 {
					buffer[ib] = nextItem
				}

			}

		}
	}

	return notices

}