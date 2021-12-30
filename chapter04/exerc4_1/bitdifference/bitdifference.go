package bitdifference

func BitDifference(a, b [32]byte) int {
	if len(a) != len(b) {
		return -1
	}

	count := 0
	for j := 0; j < len(a); j++ {
		x := a[j]
		y := b[j]
		for i := 0; i < 8; i++ {
			if ((x >> i) & 1) != ((y >> i) & 1) {
				count++
			}
		}
	}
	return count
}
