package problem1095

type MountainArray []int

func (m *MountainArray) get(idx int) int { return (*m)[idx] }
func (m *MountainArray) length() int     { return len(*m) }
