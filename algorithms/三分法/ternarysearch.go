package ternary

/*
https://oi-wiki.org/basic/binary/

用来找到： 单峰函数的极值
 */

func ternarySearch() {

	var f func(int) int
	var r, l, eps int

	for r-l > eps {
		mid := (l + r) >> 1
		lmid := mid - eps
		rmid := mid + eps
		if f(lmid) < f(rmid) {
			r = mid
		} else {
			l = mid
		}
	}

}
