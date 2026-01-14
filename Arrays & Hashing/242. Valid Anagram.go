func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    count := [26]int{}
    for i := 0; i < len(s); i++ {
        count[s[i]-'a']++
        count[t[i]-'a']--
    }

    for _, val := range count {
        if val != 0 {
            return false
        }
    }
    return true
}

// 解法
// 1. 如果兩個字串長度不同，立即回傳 false。
// 2. 建立一個大小為 26、初始值為 0 的頻率陣列 count。
// 3. 同時遍歷兩個字串：
//     對應到 s[i] 的索引，計數器加 1。
//     對應到 t[i] 的索引，計數器減 1。
// 4. 處理完兩個字串後，掃描整個 count 陣列：
//     如果有任何值不是 0，回傳 false，表示字母頻率不同。
//     如果所有值都是 0，回傳 true，表示兩個字串是 anagram（字母異位詞）。

// 觀念
// 這題的重點就是 把字母座標系平移到陣列能用的起點
// 字元其實是數字，用的是 ASCII / Unicode 編碼。
// 'a' 的編碼是 97
// 'b' 是 98
// 'z' 是 122
// 所以
// 'a' - 'a' = 0
// 'b' - 'a' = 1
// 而只有 'a' 是那個能讓最小字母剛好對齊 0 的基準點。
// 'c' - 'a' = 2

// 複雜度
// 時間複雜度 = O(n)
// 空間複雜度 = O(1)（常數空間）
