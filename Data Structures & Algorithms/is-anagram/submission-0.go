func isAnagram(s string, t string) bool {

    if len(s) != len(t) {
        return false
    }
    count := make(map[rune]int)
    for _, n := range s {
        count[n]++
    }
    for _, m := range t {
        count[m]--
    }
    for _, k := range count {
        if k != 0 {
            return false
        }
    }
    return true
}
