package main

func lonelyinteger(a []int32) int32 {
	var found int32

	ints := make(map[int32]int32)
	for i := 0; i < len(a); i++ {
		ints[a[i]] += 1
	}

	for k := range ints {
		if ints[k] == 1 {
			found = k
		}
	}

	return found
}

func main() {
	lonelyinteger([]int32{1, 2, 3, 4, 3, 2, 1})
}
