func twoSum(nums []int, target int) []int {
    prevMap := make(map[int]int)

    for i, n := range nums {
        diff := target - n
        if j, found := prevMap[diff]; found {
            return []int{j, i}
        }
        prevMap[n] = i
    }
    return []int{}
}

// 複雜度:
// 時間 O(n)，空間 O(n)
// 
// 使用先找，再存的方法是避免使用同個元素兩次
