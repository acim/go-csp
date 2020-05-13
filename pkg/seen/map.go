package seen

// Map to keep status of seen integers.
type Map map[int]bool

// CreateMap creates map of seen positive numbers.
func CreateMap(n int) Map {
	m := map[int]bool{}
	for i := 1; i <= n; i++ {
		m[i] = false
	}
	return m
}

// Seen returns whether an integer is seen.
// Take care than n should be a positive number.
func (m Map) Seen(n int) bool {
	return m[n]
}

// SetSeen sets that an integer argumen is seen.
// Take care than n should be a positive number.
func (m Map) SetSeen(n int) {
	m[n] = true
}

// AllSeen return whether all ingeters in the map are seen.
func (m Map) AllSeen() bool {
	for _, v := range m {
		if !v {
			return false
		}
	}
	return true
}
