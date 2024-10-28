package functions

// import "fmt"

func Subsets(nums []int) [][]int {
    var sol [][]int
    var temp []int
    generateSubsets(nums, 0, temp, &sol)
    return sol
}

func generateSubsets(nums []int, index int, temp []int, sol *[][]int) {
    if index == len(nums) {
        *sol = append(*sol, append([]int{}, temp...))  // Append a copy of temp
        return
    }
    // Exclude current element
    generateSubsets(nums, index+1, temp, sol)
    // Include current element
    temp = append(temp, nums[index])
    generateSubsets(nums, index+1, temp, sol)
    // Backtrack
    temp = temp[:len(temp)-1]
}
