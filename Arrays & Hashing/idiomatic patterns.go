✅ Hash Map 查找模板（Two Sum 型）💯

m := make(map[int]int)
for i, v := range arr {
    if j, ok := m[target-v]; ok {
        return ...
    }
    m[v] = i
}
