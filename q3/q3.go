package q3

//Você está subindo uma escada. Leva n passos para chegar ao topo.
//
//A cada vez, você pode subir 1 ou 2 degraus. De quantas maneiras distintas você pode subir até o topo?

func fatorial(n int) (resultado int) {
	if n > 0 {
		resultado = n * fatorial(n-1)
		return resultado
	}
	return 1
}

func ClimbStairs(n int) int {
	pos := 0
	menos := 1
	if n%2 == 0 {
		pos++
		for i := 1; i < n; i++ {
			pos += fatorial(n-menos) / (fatorial(n-i-menos) * fatorial(i))
			menos++
		}
	} else {
		pos++
		for i := 1; i <= (n / 2); i++ {
			pos += fatorial(n-menos) / (fatorial(n-i-menos) * fatorial(i))
			menos++
		}
	}
	return pos
}
