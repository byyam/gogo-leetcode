package leetcode

func canMeasureWater(x int, y int, z int) bool {
	// 若z为0，空杯就可以满足
	if z == 0 {
		return true
	}
	// x+y 如果小于z，则不可能恰好为z
	if x+y < z {
		return false
	}
	// a*x + b*y = z, 是否可以取到整数a,b
	g := gcd(x, y)
	if g == 0 {
		return false
	}
	if z%g == 0 {
		return true
	}

	return false
}

func gcd(x, y int) int {
	max := x
	if x > y {
		max = y
	}
	gcd := 0
	for g := 1; g <= max; g++ {
		if x%g == 0 && y%g == 0 {
			gcd = g
		}
	}
	return gcd
}
