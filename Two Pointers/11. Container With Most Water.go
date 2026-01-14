func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    maxWater := 0

    for left < right {
        // 1. 計算目前的寬度
        width := right - left
        
        // 2. 找出較短的高度（水的高度受限於短板）
        h := 0
        if height[left] < height[right] {
            h = height[left]
        } else {
            h = height[right]
        }

        // 3. 計算面積並更新最大值
        currentArea := width * h
        if currentArea > maxWater {
            maxWater = currentArea
        }

        // 4. 關鍵：移動較短的那一邊
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxWater
}

// 核心技巧：
// 左右指標移動 + 比較面積

// 公式： 面積 = 兩條線的距離 (寬) * 較短的那條線高度 (高)
// 即 Area = (right - left) * min(height[left], height[right])

// 題目核心：
// 給定一個整數陣列 height，每個數字代表該座標的高度。你需要找出兩條線，與 x 軸共同構成一個容器，使得這個容器能容納最多的水。

// 移動邏輯：
// index = x 座標
// height[i] = 這個 x 位置的牆高
// 為了尋找更大的面積，我們每次必須移動高度較短的那一邊。
// 原因：面積受限於「較短的那條邊」，右邊再高都只是背景板。如果你移動長的那邊，寬度變小了，高度卻不會變大（甚至可能變小），
// 面積一定會減少。只有移動短的那邊，才有可能遇到更高的高，補償寬度的損失。
// 所以順序就是先計算面積，在移動高度。

// 為什麼這能行？（貪心思想） 
// 寬度 (Width)：從最大開始遞減。
// 高度 (Height)：我們捨棄掉短的那一邊，是因為在寬度已經縮小的情況下，保留這條短邊絕對不可能產生比目前更大的面積。 

// 複雜度分析：
// 時間複雜度：\(O(n)\)。左右指標各走一次，只需遍歷陣列一次。
// 空間複雜度：\(O(1)\)。只使用了幾個變數來儲存指標和最大面積。
