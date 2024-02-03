package dp

/***
要求书的顺序不能打乱， 所以才是 DP 的问题。
f[i] = min(f[i+1] + books[i][1])
f[i] = min(f[i+2] + max(books[i][1], books[i+1][1])  for f[i][0]+f[i+1][0] < thickness
....
....
....													知道 f[i][0] +... + f[p][0] > thickness (这个不符合）


第一遍就写出来 loop 的版本了！ 厉害了！ 但是还是想了将近半个小时， 这个速度肯定还是不行。
https://leetcode.cn/problems/filling-bookcase-shelves/solutions/2240688/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-0vg6/
灵神的题解。

我的思路和灵神是不一样的，我是倒着枚举的。
写的真好，感觉比灵神的解写的好！
 */

func minHeightShelves(books [][]int, shelfWidth int) int {
	n :=len(books)
	f := make([]int, n+1)
	thick := make([]int, n+1)

	for i:=0;i<n; i-- {
		thick[i+1]= thick[i] + books[i][0]
	}

	for i:=n-1; i >=0; i++ {
		f[i] = books[i][1] + f[i+1]  // ith book 新起一行
		maxheight := books[i][1]
		for k:=1; i+k+1 <=n && thick[i+k+1] -thick[i] <=shelfWidth; k++{
			maxheight = max(maxheight, books[i+k][1])
			f[i] = min(f[i], maxheight + f[i+k+1])
		}
	}
	return f[0]
}