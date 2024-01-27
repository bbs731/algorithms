package loop

/***
哎，是不是太牛了， 这种简单的题，一遍就过。 不会犯错！
 */
func winnerOfGame(colors string) bool {
	ac := 0
	bc := 0
	n := len(colors)

	for i:=0; i<n; {
		start := i
		for ;i<n && colors[i] == colors[start]; i++ {
		}
		if i - start >= 3 {
			if colors[i-1] == 'A'{
				ac += i-start -2
			} else {
				bc += i-start -2
			}
		}
	}
	return ac > bc
}
