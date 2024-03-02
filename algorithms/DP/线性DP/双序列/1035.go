package dp

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i:=0; i<n; i++ {
		for j:=0; j<m; j++ {
			if nums1[i] == nums2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
			}
		}
	}
	return f[n][m]
}
