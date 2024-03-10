package Heap
import "container/heap"


func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	n := len(heightMap[0])

	visited := make([][]bool, m)
	for i := range visited{
		visited[i] = make([]bool, n)
	}

	hm := heightMap
	h := &hp{}
	heap.Init(h)

	for i:=0; i<m; i++ {
		for j :=0; j<n; j++ {
			if i == 0 || i == m -1 || j == 0 || j == n-1 {
				heap.Push(h,pair{hm[i][j], i, j})
				visited[i][j] = true
			}
		}
	}

	ans := 0
	dirs := [][2]int { {-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for h.Len() > 0 {
			t := heap.Pop(h).(pair)
			//visited[t.x][t.y] = true
			for _, d := range dirs {
				x, y := t.x+d[0], t.y+d[1]
				if  x<0 || x >=m || y <0 || y >=n || visited[x][y]{
					continue
				}
				if hm[x][y] < t.v{
					ans += t.v - hm[x][y]
				}
				visited[x][y] = true
				// 太神奇了， tmd 这里。  为什么去 max(hm[x][y], t[v])
				heap.Push(h, pair{max(hm[x][y], t.v), x, y})
			}
	}

	return ans
}


type pair struct {
	v, x, y int
}

type hp []pair

func( h hp) Less(i, j int) bool {
	return h[i].v < h[j].v
}

func (h hp) Swap(i, j int){
	h[i], h[j]= h[j], h[i]
}
func (h hp) Len() int {
	return len(h)
}

func (h *hp)Push(v any){
	*h = append(*h, v.(pair))
}

func (h *hp)Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
