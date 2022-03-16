package unit

func unique1(elements []int) []int {
	keys := map[int]bool{}
	list := []int{}

	for i := range elements {
		if keys[elements[i]] == false {
			keys[elements[i]] = true
			list = append(list, elements[i])
		}
	}
	return list
}

func unique2(elements []int) []int {
	var list []int

	for i := 0; i < len(elements); i++ {
		exists := false
		for v := 0; v < i; v++ {
			if elements[v] == elements[i] {
				exists = true
				break
			}
		}
		if !exists {
			list = append(list, elements[i])
		}
	}
	return list
}

func unique3(elements []int) []int {
	keys := make(map[int]bool)
	var list []int
	for _, entry := range elements {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
