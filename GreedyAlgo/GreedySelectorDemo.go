package GreedyAlgo

func GreedySelector(n int, s []int, f []int, A []bool){
	A[0] = true
	j := 0

	for i := 1; i < n; i++{
		if s[i] >= f[j]{
			A[i] = true
			j ++
		}else{
			A[i] = false
		}
	}
}
