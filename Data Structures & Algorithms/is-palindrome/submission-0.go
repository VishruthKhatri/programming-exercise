func isPalindrome(s string) bool {
	
	right := len(s) - 1
	left := 0
	
	for left < right {

	 for left < right && !alphacheck(s[left]){
		left++
	 }
	 for left < right && !alphacheck(s[right]){
		right--
	 }
	 if left<right {
		b1 := s[left]
		b2 := s[right]
		if b1 >= 'A' && b1 <= 'Z' { b1 += 32 }
			if b2 >= 'A' && b2 <= 'Z' { b2 += 32 }

			if b1 != b2 {
				return false
			}
			left++
			right--
		
	 }
		
	}
	return true
}
func alphacheck(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
	}