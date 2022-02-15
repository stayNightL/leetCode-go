package offer


func replaceSpace(s string) string {

	//计算所需总长度
	var all int =0
	for _, i := range s {
		if i == ' ' {
			all=all+3
			continue
		}
		all+=1
	}
	runes := make([]rune, all)
	index := len(s)-1

	for i := index; i >=0; i-- {
		if s[i] == ' ' {
			runes[all-1]='0'
			runes[all-2]='2'
			runes[all-3]='%'
			all = all-3
		}else {

			runes[all-1]= rune(s[i])
			all=all-1
		}
	}
	return string(runes)
	// 从后向前
}