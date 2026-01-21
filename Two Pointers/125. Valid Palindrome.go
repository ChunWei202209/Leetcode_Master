import (
    "fmt"
    "unicode"
)

func isPalindrome(s string) bool {
    runes := []rune(s)
    left, right := 0, len(runes)-1

    for left < right {
        // 跳過左邊非字母或數字
        if !unicode.IsLetter(runes[left]) && !unicode.IsDigit(runes[left]) {
            left++
            continue
        }

        // 跳過右邊非字母或數字
        if !unicode.IsLetter(runes[right]) && !unicode.IsDigit(runes[right]) {
            right--
            continue
        }

        // 比較兩個字元（忽略大小寫）
        if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
            return false
        }

        // 指標移動
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
