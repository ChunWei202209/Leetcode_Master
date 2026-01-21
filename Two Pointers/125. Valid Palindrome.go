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

// 解題步驟
// 1. 轉換為 Rune 切片：
//   a. 將 string 轉為 []rune。這是為了正確處理多位元組字元（如中文或特殊符號），避免直接使用索引 s[i] 存取到破碎的 UTF-8 位元組。
// 2. 初始化雙指標：
//   b. 設定 left 指標指向字串開頭（索引 0），right 指標指向字串結尾（索引 len-1）。
// 3. 過濾無效字元（向內移動指標）：
//   a. 左指標：若遇到非字母且非數字的字元（如空格、標點），則持續向右移動 left++。
//   b. 右指標：若遇到非字母且非數字的字元，則持續向左移動 right--。
//   c. 此步驟確保接下來比較的對象都是有效的「字母或數字」。
// 4. 忽略大小寫並比較：
//   a. 使用 unicode.ToLower 將左右字元統一轉為小寫後進行比較：
//   b. 不相等：直接回傳 false，表示這不是回文字串。
//   c. 相等：代表這組字元符合回文規則，將 left 右移一格、right 左移一格，繼續下一輪檢查。

// 複雜度分析：
// 1. 時間複雜度：O(n) 字串轉換：
// 將字串轉為 []rune 需要遍歷一次字串，耗時 (O(n))。
// 雙指標遍歷：雖然有巢狀的 for 迴圈，但 left 指標只會從左向右走，right 只會從右向左走，每個字元最多只會被存取一次。
// 2. 空間複雜度：O(n) Rune 切片：
// 程式碼中宣告了 runes := []rune(s)，這會建立一個與原字串長度相等的新切片來存放 Unicode 字元，因此需要 O(n) 的額外空間。
