package others

func lastRemaining(n int, m int) int {
	// f(N, M) = (f(N-1, M) + M)%N
	end := 0
	for i := 2; i <= n; i++ {
		end = (end + m) % i
	}
	return end
}
