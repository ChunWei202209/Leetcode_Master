func containsDuplicate(nums []int) bool {
    seen := make(map[int]bool)
  
    for _, num := range nums {
        if seen[num] {
            return true
        }
        seen[num] = true
    }
    return false
}

// Go map 的特性:
// 讀取一個不存在的 key → 會得到對應 type 的「零值」
// map[int]bool → 沒出現過的 key 讀出來是 false
// 已經出現過的 key → 我們存 true → 判斷時就可以直接用 if seen[num]
// 所以這裡用 true 是最簡單、最直觀的選擇。

// 複雜度:
// 空間: O(n)
// 時間: O(n)
