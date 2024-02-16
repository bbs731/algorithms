package interview_questions

import "strings"

func simplifyPath(path string) string {

	st := []string{}
	start , end := -1, 0
	path += "/"  // 边界条件总是忘了啊， 依赖最后一个 / 触动 processing
	for i:=0; i<len(path); i++ {
		if path[i] == '/' {
			// processing the [start, end] substr
			if start == -1 {
				continue
			}
			if end - start + 1 == 2  && path[end] == '.'  && path[end-1]== '.'{
				if len(st)> 0 {
					// pop stack
					st = st[:len(st)-1]
				}
			} else if end -start +1 == 1 && path[end] == '.' {
				// do nothing
			} else {
				st = append(st, path[start:end+1])
			}
			start = -1
		} else {
			if start == -1 {
				start = i
			}
			end = i
		}
	}
	return "/" + strings.Join(st, "/")
}


// 牛叉啊， 逻辑好清晰啊
func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	stack := make([]string, 0)
	for _, v := range arr {
		if v == "" || v == "." {
			continue
		} else if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, v)
		}
	}
	var ret = "/"
	return ret + strings.Join(stack, "/")
}