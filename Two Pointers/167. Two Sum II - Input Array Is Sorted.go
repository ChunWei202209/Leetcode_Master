func twoSum(numbers []int, target int) []int {
    // 初始化左右兩個指標
    left := 0
    right := len(numbers) - 1

    for left < right {
        currentSum := numbers[left] + numbers[right]

        if currentSum == target {
            // 題目要求回傳 1-based index，所以索引都要加 1
            return []int{left + 1, right + 1}
        } else if currentSum < target {
            // 如果和小於目標值，將左指標右移以增加總和
            left++
        } else {
            // 如果和大於目標值，將右指標左移以減少總和
            right--
        }
    }

    // 根據題目保證一定有一組解，此處通常不會執行到
    return []int{}
}

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

// 解題步驟： 
// 1. 初始化：我們將 left 指標指向陣列的第一個元素（索引 0），right 指標指向最後一個元素。
// 2. 移動規則：如果 numbers[left] + numbers[right] 剛好等於 target，
//     恭喜找到答案，回傳 [left + 1, right + 1] (題目要求將 index+1)。
// 3. 如果 總和小於 target：
//     因為陣列是排序過的，為了讓總和變大，我們必須把 left 往右移（指向更大的數）。
//     如果 總和大於 target：為了讓總和變小，我們必須把 right 往左移（指向更小的數）。

// 複雜度分析：時間複雜度：\(O(n)\)。在最壞的情況下，兩個指標會遍歷整個陣列一次。
// 空間複雜度：\(O(1)\)。我們只額外使用了兩個整數變數（left 與 right）。
