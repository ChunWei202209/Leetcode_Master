func twoSum(nums []int, target int) []int {
    prevMap := make(map[int]int)

    for i, n := range nums {
        diff := target - n
        
        // j 是「之前那個數字在陣列裡的位置（index）」
        if j, found := prevMap[diff]; found { 
            return []int{j, i}
        }
        
        // 將 n 存到 map 的 key 中，而 value 就是之前該數字在 nums 的 index
        prevMap[n] = i 
    }
    return []int{}
}

// https://leetcode.com/problems/two-sum/description/

// 目標: 還傳兩個不同數字的 index。
// i：目前這個數字的 index
// j：就是「之前某一次迴圈的 i，被你存進 prevMap 的那個值」

// 解題關鍵:
// 使用 先找，再存 的方法是避免使用同個元素兩次。

// 解法順序：
// 1. 建立一個雜湊表（hash map），用來儲存陣列中每個元素的值以及它對應的索引。
    // key：數字（nums 裡出現過的值）
    // value：那個數字的 index
// 2. 使用索引 i 迭代整個陣列，並計算目前元素所需的補數（complement），也就是 target - nums[i]。
// 3. 檢查這個補數是否存在於雜湊表的 key 中。
    a. Go 的 map 查詢語法：
    b. value, ok := myMap[key]
    c. value：key 對應的值
    d. ok：key 存不存在（bool）
// 4. 如果存在，回傳目前元素的索引以及該補數對應的索引 (j：之前出現過的那個數字的 index)。
// 5. 如果遍歷完整個陣列後仍然找不到符合條件的數字組合，則回傳一個空的陣列。

// 複雜度:
// 時間 O(n)，空間 O(n)
