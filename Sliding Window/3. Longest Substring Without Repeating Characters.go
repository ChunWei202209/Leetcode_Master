func lengthOfLongestSubstring(s string) int {
    // 記錄目前視窗內出現過的字元
    // Go 的 string 本質是 一串 byte，而他在電腦裡可以代表字元
    seen := make(map[byte]bool)

    left := 0 // 視窗左邊界
    maxLen := 0
  
    // 視窗右邊界一路往右走
    for right := 0; right < len(s); right++ { 
        // 如果右邊的字在 map 是否已經出現過，值是 true
        for seen[s[right]] {
            // 把左邊的字移出視窗，就是字元是 s[left]」這一筆 key
            delete(seen, s[left])
            left++
        }

        // 把目前字元加入視窗
        seen[s[right]] = true

        // 更新最大長度
        // 我們要 每次都檢查是否是目前最長，而不是在 loop 之後檢查，會漏掉中間最長的子串
        if right-left+1 > maxLen {
            maxLen = right - left + 1
        }
    }

    return maxLen
}

// https://neetcode.io/problems/longest-substring-without-duplicates/solution

// 題目要求：
// 給你一個字串，找出「不包含重複字元的最長連續子字串」的長度。

// 重點有三個：
// Substring：一定要「連續」
// Without Repeating：裡面的字不能重複
// Longest：要最長的那一段

// 核心技巧：
// 用 map 記字元位置 / 窗口左右界移動

// 解題步驟：
// 1. 用 map 記錄「目前視窗內有哪些字」，「我用 byte（0~255 的數字）來當 key，記錄哪些字出現過」。
      key：字元（像 'a', 'b'）。
      value：true 表示「在房間裡」。
// 2. 用 right 向右走，一個字一個字看。
// 3. 如果發現重複字，就移動 left 把重複字移掉。
      delete(map, key) 是 delete(誰, 刪誰) 的意思，跟 map 裡面的順序無關。
// 4. 每一步都更新「目前最長長度」。
      視窗長度是右邊界 (index) - 左邊界 (index)  + 1。
