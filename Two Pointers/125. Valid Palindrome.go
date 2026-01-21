import (
    "fmt"
    "unicode"
)

func isPalindrome(s string) bool {
    left, right := 0, len(s)-1

    // Go 的 string 是 UTF-8
    // 如果直接用 s[i]，遇到中文或特殊字元會出錯
    // 所以先轉成 rune slice，才能安全處理「一個字元」
    runes := []rune(s)

    // 外層 for：控制整個「左右夾擊」的流程
    // 當 left >= right，代表已經比完或交錯了
    for left < right {

        // ===== 處理左邊指標 =====
        // 不斷往右移，直到找到「字母或數字」
        for left < right {
            // 如果是字母或數字，就停下來
            if unicode.IsLetter(runes[left]) || unicode.IsDigit(runes[left]) {
                break
            }
            // 否則跳過（例如空白、逗號、標點符號）
            left++
        }

        // ===== 處理右邊指標 =====
        // 不斷往左移，直到找到「字母或數字」
        for left < right {
            if unicode.IsLetter(runes[right]) || unicode.IsDigit(runes[right]) {
                break
            }
            right--
        }

        // ===== 比較左右兩邊的字元 =====
        // ToLower：忽略大小寫
        if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
            return false
        }

        // 如果相等，代表這一組字元 OK
        // 指標往中間移，準備比較下一組
        left++
        right--
    }
    return true
}


// A string is a palindrome if, 
// after removing all non-alphanumeric characters and ignoring case, 
// it reads the same forwards and backwards.

// 用 []rune(s) 可以安全處理中文或 emoji，不會出現亂碼
// unicode.IsLetter 判斷是不是字母
// unicode.IsDigit 判斷是不是數字
// unicode.ToLower 直接大小寫轉換，支援 Unicode
