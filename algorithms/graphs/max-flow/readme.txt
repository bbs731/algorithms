总纲： Ford-Fulkerson Algorithm to find augment path to increase the max-flow until no augment path can be found. (include Residue network)

如果 max-flow 的值域很大， find augment path 太慢的话, 达到终值太慢了。 时间复杂度是 O(f*E)
根据如何快速的找到 Augment Path, 有多种 algorithm.
1. Edmonds-Karp: O(E^2*V)  use BFS to find augment path
2. Dinic: O(V^2*E) use combination of BFS + DFS to find augment path
3. Capacity scaling:  O(E^2*Log(f)) adding a heuristic on top of Ford-Fulkerson to pick larger paths first
4. Push Relabel: O(V^2*E) use concept of maintaining a preflow instead of finding augment path.

最大流最小割定理的证明
https://oi-wiki.org/graph/flow/max-flow/#%E6%9C%80%E5%A4%A7%E6%B5%81%E6%9C%80%E5%B0%8F%E5%89%B2%E5%AE%9A%E7%90%86


Ford-Fulkerson 实现：
https://github.com/williamfiset/Algorithms/blob/master/src/main/java/com/williamfiset/algorithms/graphtheory/networkflow/FordFulkersonDfsSolverAdjacencyList.java