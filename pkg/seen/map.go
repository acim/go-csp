package seen

type Map map[int]bool

func CreateMap(n int) Map {
	m := map[int]bool{}
	for i := 1; i <= n; i++ {
		m[i] = false
	}
	return m
}

func (m Map) Seen(n int) bool {
	return m[n]
}

func (m Map) SetSeen(n int) {
	m[n] = true
}

func (m Map) AllSeen() bool {
	for _, v := range m {
		if !v {
			return false
		}
	}
	return true
}
