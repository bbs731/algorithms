package stack

import "bytes"

func removeDuplicates(s string, k int) string {
	type pair struct {
		c byte
		cnts int
	}
	st := []pair{}
	for i := range s {
		x := s[i]
		if len(st) >0 && st[len(st)-1].c == x {
			if st[len(st)-1].cnts == k -1 {
				// pop stack
				st = st[:len(st)-k+1]
			} else {
				st = append(st, pair{x, st[len(st)-1].cnts+1})
			}
		} else {
			st = append(st, pair{x, 1})
		}
	}

	buf := bytes.Buffer{}
	for i:=0; i<len(st); i++ {
		buf.Write([]byte{st[i].c})
	}
	return buf.String()
}
