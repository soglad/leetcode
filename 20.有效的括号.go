/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
func isValid(s string) bool {
	stack := make([]byte, len(s))
	idx := 0
	for i:=0;i<len(s);i++{
		switch s[i] {
		case '(','[','{':
			stack[idx]=s[i]
			idx++
		case ')':
			if idx>0 && stack[idx-1]=='('{
				idx--
			}else {
				return false
			}
		case ']':
			if idx>0 && stack[idx-1]=='['{
				idx--
			}else {
				return false
			}
		case '}':
			if idx>0 && stack[idx-1]=='{'{
				idx--
			}else {
				return false
			}
		default:
			continue
		}
	}
	return idx==0
}

