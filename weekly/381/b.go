package weekly

/***
    for (k = 0; k < V; k++) {
        // Pick all vertices as source one by one
        for (i = 0; i < V; i++) {
            // Pick all vertices as destination for the
            // above picked source
            for (j = 0; j < V; j++) {
                // If vertex k is on the shortest path from
                // i to j, then update the value of
                // dist[i][j]
                if (dist[i][j] > (dist[i][k] + dist[k][j])
                    && (dist[k][j] != INF
                        && dist[i][k] != INF))
                    dist[i][j] = dist[i][k] + dist[k][j];
            }
        }
    }
 */

func countOfPairs(n int, x int, y int) []int {
	inf := 3*n + 12
	dist := make([][]int, n)
	for i := range dist{
		dist[i]	 = make([]int, n)
		for j:=0; j<n; j++ {
			dist[i][j] = inf
		}
		dist[i][i] = 0
	}

	for i:=0; i<n-1; i++ {
		dist[i][i+1] = 1
		dist[i+1][i] = 1
	}
	dist[x-1][y-1] = 1
	dist[y-1][x-1] = 1

	for k:=0; k< n; k++ {
		for i:=0; i<n; i++{
			for j:=0; j<n; j++{
				if dist[i][j] > dist[i][k] + dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	ans := make([]int, n)
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			if i != j {
				ans[dist[i][j]-1]++
			}
		}
	}
	return ans
}