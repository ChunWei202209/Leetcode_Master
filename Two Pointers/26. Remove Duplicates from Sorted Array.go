func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    // slow 指向目前處理好、不重複序列的最後一個位置
    slow := 0

    // fast 從第二個元素開始遍歷
    for fast := 1; fast < len(nums); fast++ {
        // 如果快指標發現了一個與慢指標不同的新元素
        if nums[fast] != nums[slow] {
            // 慢指標往後移一格
            slow++
            // 把新的元素蓋掉原本重複的位置
            nums[slow] = nums[fast]
        }
    }

    // 回傳長度。因為 slow 是索引 (0-based)，所以長度要 +1
    return slow + 1
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/

// 核心技巧：
// 左指標追蹤位置，右指標遍歷

// 關鍵點解析： 
// 1. 為什麼要原地修改？
//     題目要求空間複雜度必須是 \(O(1)\)，所以我們不能建立一個新的切片 (Slice) 或 Map 來存放結果，
//     必須直接修改 nums。
// 2. 邏輯運作過程：
//     a. 假設 nums = [0, 0, 1, 1, 1, 2, 2]。
//     b. 一開始 slow = 0, nums[slow] = 0。
//     c. fast 跑到索引 2 時發現 nums[2] = 1（與 0 不同）。
//     d. slow 變成 1，並把 nums[1] 改成 1。
//     e. 陣列變成 [0, 1, ...]。以此類推，直到 fast 走完。
// 3.回傳值 slow + 1：
//     a. LeetCode 的評判系統會根據你回傳的長度 \(k\)，去檢查 nums 陣列的前 \(k\) 個元素是否正確。
//     b. 因為 slow 是從 0 開始的索引，所以最終不重複元素的個數就是 slow + 1。 

// 複雜度： 
// 時間複雜度：\(O(n)\)，其中 \(n\) 為陣列長度。只需遍歷一次陣列。
// 空間複雜度：\(O(1)\)，只使用了額外的指標變數。
