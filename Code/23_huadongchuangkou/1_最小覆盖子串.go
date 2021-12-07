package main

func main()  {
	s := "ab"
	t := "eidboaoo"

	println(checkInclusion(s, t))
}

func minWindow(s string, t string) string {

	window := make(map[byte]int,0)
	need := make(map[byte]int,0)

	for i := range t {
		need[t[i]]++
	}

	left,right,match :=0,0,0
	// 记录最小覆盖子串的起始索引及长度
	start,end := 0,0
	min := len(s) + 1

	for right < len(s) {
		// 将s[right]加入，形成(left,right]
		ch1 := s[right]
		window[ch1]++

		// 右移窗口
		right++

		// 更新状态
		if window[ch1] == need[ch1] {
			match++
		}

		// 判断左侧窗口是否需要收缩
		for match == len(need) {
			// 更新最小覆盖子串
			if right -left < min {
				start,end = left,right
				min = right -left
			}

			// 获取将要移出窗口的字符
			ch2 := s[left]
			left++

			// 更新状态
			if window[ch2] == need[ch2] {
				match--
			}
			window[ch2]--
		}
	}

	return s[start:end]
}

func checkInclusion(s1 string, s2 string) bool {
	left,right,match := 0,0,0
	window := make(map[byte]int,0)
	need := make(map[byte]int,0)

	for i := range s1 {
		need[s1[i]]++
	}

	for right < len(s1) {
		ch1 := s2[right]
		right++
		window[ch1]++

		if window[ch1] == need[ch1] {
			match++
		}

		for right - left >= len(s2) {
			// 判断是否有合法的子串
			if match == len(need) {
				return true
			}

			ch2 := s2[left]
			left++

			if window[ch2] == need[ch2] {
				match--
			}
			window[ch2]--
		}
	}

	return false
}