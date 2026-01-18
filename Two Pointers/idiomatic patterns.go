3️⃣ Two Pointers 模板（有序陣列）

l, r := 0, len(arr)-1
for l < r {
    sum := arr[l] + arr[r]
    if sum < target {
        l++
    } else if sum > target {
        r--
    } else {
        return ...
    }
}
