package max_flow

// 最小费用最大流 MCMF（即满流时的费用）
//https://oi-wiki.org/graph/flow/min-cost/

//下面的代码 minCostFlowSPFA 来自：
//https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/graph.go#L4308
//框架是 Edmonds-Karp 算法， 把 Edmonds-Karp 中发现增广路的 bfs() 改成 队列优化的Bellman-Ford 也就是 SPFA 算法。因为边权可能是负值，
//所以这里没有使用 Dijkstra 而使用 队列优化的 Belleman-Ford。
//mcmf 整个算法的时间复杂度是 O(nmf) (注意哈， Edmonds-Karp 计算最大流的算法的时间复杂度是 O（V*E^2) 的，没有depend on flow 的值域，
//但是基于 Edmonds-Karp 的 mcmf 的算法的时间复杂度是有f, flow 的值域项的）
//我们还知道 Dinic 再求最大流的时候时间复杂度是O（V^2*E) 如果用 Dinic 求 mcmf 时间复杂度还有 f, flow 的值域项吗？我猜测还是有的，如何证明？
//因为 Edmonds-Karp 改进了 Ford-Fulkerson 算法(Ford-Fulkerson 用 DFS去找增广路的话）让它不依赖 f 值域。 但是没办法让 mcmf 不依赖 f 的值域。 Dinic应该也不可以。
//另外 mcmf 是有快速算法的，网络单纯型法，（解决 linear programing 的 simplex 吗？） 但是太复杂了，留给以后吧。
