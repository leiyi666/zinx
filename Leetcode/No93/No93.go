package No93

func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 || n > 12 {
		return make([]string, 0)
	}
	temp := ""
	var result []string = make([]string, 0)
	res := dfs(s, n, 0, temp, result)
	return res
}

func dfs(s string, length int, depth int, temp string, result []string) []string {
	if depth == 4 && length == 0 {
		result = append(result, temp)
		return result
	} else if depth == 4 && length != 0 {
		return result
	}
	cur := temp
	for i := 0; i < length; i++ {
		if i == 3 {
			break
		}
		ss := s[:i+1]
		if (i+1 == 3 && ss > "255") || (ss[0] == '0' && i+1 != 1) {
			continue
		}
		if temp == "" {
			temp = ss
		} else {
			temp = temp + "." + ss
		}
		result = dfs(s[i+1:], length-i-1, depth+1, temp, result)
		temp = cur
	}
	return result
}
