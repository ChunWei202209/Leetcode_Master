import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 1. 先排序
	res := [][]int{}
	n := len(nums)
	
	// 外層去重：固定的第一個數不能重複
	for i := 0; i < n-2; i++ {
		// 如果目前的數已經大於 0，後面的數也一定大於 0，不可能相加等於 0
		if nums[i] > 0 {
			break
		}

		// 去重：如果這個數字跟前一個數字一樣，跳過它以免結果重複
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 固定 nums[i]，在後面的範圍內找 two sum
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				
				// 去重：跳過相同數值，避免產生重複的三元組
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 去重：跳過相同數值，避免產生重複的三元組
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				
				// 移動指針（換位置）
				left++
				right--
			} else if sum < 0 {
				left++ // 和太小，往右移找大一點的數
			} else {
				right-- // 和太大，往左移找小一點的數
			}
		}
	}
	return res
}

// https://leetcode.com/problems/3sum/description/

// 核心技巧：
// 一指標固定，左右指標往中間掃，求和等於目標

// 題目要求：
// 找出陣列中所有「不重複」的三元組 [nums[i], nums[j], nums[k]]，使得它們的和為 0。
// 注意： 結果集中不能包含重複的三元組。
// 解法：排序 + 雙指標 (Sorting + Two Pointers)
// 要解決這題，最關鍵的步驟是 先排序。排序後，我們固定一個數，然後將剩下的問題轉化為 Two Sum II（雙指標法）。

// 演算法步驟：
// 1. 排序：將陣列由小到大排序。
// 2. 遍歷固定第一個數 i：
//     a. 如果 nums[i] > 0，因為陣列已排序，後面的數也都大於 0，不可能湊成 0，直接結束。
//     b. 去重：如果 nums[i] 與前一個數字相同，跳過（避免重複解）。
// 3. 雙指標找剩下兩個數：
//     a. 設 left = i + 1, right = len - 1。
//     b. 當 sum == 0：存入結果，並將 left 和 right 往中間移動，移動過程中要跳過重複的數字。
	   c. 必須去重（值層級） 
//     d. 當 sum < 0：將 left 右移（增加總和）。
//     e. 當 sum > 0：將 right 左移（減少總和）。

// 關鍵細節：
// 為什麼要去重？ 
// 題目要求「不重複的三元組」。
// 例如 nums = [-1, -1, 0, 1]： 當 i=0 時，我們找到了 [-1, 0, 1]。
// 當 i=1 時，如果我們不執行 if nums[i] == nums[i-1] 的判斷，會再次找到 [-1, 0, 1]。
// 同理，內層的 left 和 right 也需要去重，否則會收錄重複的組合，去的是「值」，不是「index」。 
// 為何 sum < 0 或 sum > 0 不去重？ 
// 若 sum < 0 或 sum > 0，代表這一組不符合要求，所以直接移動 left 或 right；
// 即便數字跟之前相同，因為不符合要求，就不用去重。

// 複雜度分析 ：
// 時間複雜度：
// O(n^2)。
// 排序需要 O(nlog n)，外層迴圈遍歷 (n) 次，內層雙指標遍歷 (n) 次，總計 O(n^2)。
// 空間複雜度：
// O(1) 或 O(log n)。
// 不計入回傳結果空間的話，僅取決於排序算法的額外空間需求。
