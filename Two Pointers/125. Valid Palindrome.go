import (
    "fmt"
    "unicode"
)

func isPalindrome(s string) bool {
    left, right := 0, len(s)-1
    runes := []rune(s) // 轉成 rune 才能安全處理非 ASCII

    for left < right {
        // 找左邊的有效字元
        for left < right {
            if unicode.IsLetter(runes[left]) || unicode.IsDigit(runes[left]) {
                break // 找到字母或數字就停
            }
            left++ // 不是的話就往右移
        }

        // 找右邊的有效字元
        for left < right {
            if unicode.IsLetter(runes[right]) || unicode.IsDigit(runes[right]) {
                break // 找到字母或數字就停
            }
            right-- // 不是的話就往左移
        }
      
        // 比較兩個字元（忽略大小寫）
        if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
            return false // 不一樣就不是回文
        }
      
        // 移動指標，繼續下一輪比較
        left++
        right--
    }

    return true
}

// A string is a palindrome if, 
// after removing all non-alphanumeric characters and ignoring case, 
// it reads the same forwards and backwards.

// unicode.IsLetter 判斷是不是字母
// unicode.IsDigit 判斷是不是數字
// unicode.ToLower 直接大小寫轉換，支援 Unicode
// 用 []rune(s) 可以安全處理中文或 emoji，不會出現亂碼
