package coinchange

func Change(n int) []int {
	var change []int
	for n > 0 {
		switch {
		case n > 25:
			change = append(change, 25)
			n -= 25
		case n > 10:
			change = append(change, 10)
			n -= 10
		case n > 5:
			change = append(change, 5)
			n -= 5
		default:
			change = append(change, 1)
			n--
		}
	}
	return change
}
