package functions

//https://www.interviewbit.com/problems/subarray-with-given-xor/

func SubArraysWithXor(A []int, B int) int {
	count := 0
	mpp := map[int]int{}
	xr := 0
	mpp[0] = 1
	for i := 0; i < len(A); i++ {
		xr = xr ^ A[i]
		p := xr ^ B
		occu, ok := mpp[p]
		if ok {
			count += occu
			mpp[xr]++
		} else {
			mpp[xr] += 1
		}

	}
	return count
}
