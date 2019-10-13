package main

func hcf(big int, small int) int {
	var rem int = big % small
	if (rem == 0) {
		return small
	} else {
		return hcf(small, rem)
	}
}
