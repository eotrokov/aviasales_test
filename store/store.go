package store

var arrays = []string{}

func AddToStore(arr []string) {
	arrays = union(arrays, arr)
}

func GetStore() []string {
	return arrays
}

func union(a, b []string) []string {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			a = append(a, item)
		}
	}
	return a
}
