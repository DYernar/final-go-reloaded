package student

func ActiveBits(n int) uint {
	var activeBits uint
	for n > 0 {
		reminder := n%2 
		if reminder == 1 {
			activeBits++
		}
		n = n/2
	}
	
	return activeBits
}