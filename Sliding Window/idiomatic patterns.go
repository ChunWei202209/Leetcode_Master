✅ Sliding Window 模板（Longest Substring 型）💯

m := make(map[byte]bool)
left := 0
for right := 0; right < len(s); right++ {
    for m[s[right]] {
        delete(m, s[left])
        left++
    }
    m[s[right]] = true
    // 更新答案
}
