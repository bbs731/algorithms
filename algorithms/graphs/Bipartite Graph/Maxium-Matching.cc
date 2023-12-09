/*

Augmenting Path Algorithm

定义：
交错路（alternating path) 始于非匹配点且由匹配边与非匹配边交错而成。
增广路（augmenting path）是始于非匹配点且终于非匹配点的交错路。


Berge's 定理：
states that a matching M in a graph G is maximum (contains the largest possible number of edges) if and only if there is no augmenting path (a path that starts and ends on free (unmatched) vertices, and alternates between edges in and not in the matching) with M.
当找不到增广路的时候，得到最大匹配。


Augmenting Path Algorithm 用来实现 Berge's theorem

时间复杂度： O(nm)
*/

/*
    找到一条增广路，  未匹配边 ， 匹配边，  未匹配边， 匹配边， .....,  为匹配边.
    把这条增广路（翻转，未匹配边，变成匹配的，原匹配的拆开匹配，变成不匹配） 会发现，这样匹配数会 +1 变大了。
    Berge's 定理说， 找不到这种增广路之后， 匹配数不能再变大了，则达到最大匹配
*/

/*

https://oi-wiki.org/graph/graph-matching/bigraph-match/
 因为增广路长度为奇数，路径起始点非左即右，所以我们先考虑从左边的未匹配点找增广路。 注意到因为交错路的关系，增广路上的第奇数条边都是非匹配边，第偶数条边都是匹配边，于是左到右都是非匹配边，右到左都是匹配边。 于是我们给二分图 定向，问题转换成，有向图中从给定起点找一条简单路径走到某个未匹配点，此问题等价给定起始点  能否走到终点 。 那么只要从起始点开始 DFS 遍历直到找到某个未匹配点

解释：
   我们用 pa, pb 数组来保存，目前找到的匹配的信息。 为了方便，用了两个数组， 一个描绘的是从左到右的匹配， 一个描绘的是从右向左的匹配.
   为了找增广路， 方便我们思考，我们从左边找一个未匹配的点开始， 并且认为，从左到右是未匹配边,从右往左，是匹配的边.

*/


#include <bits/stdc++.h>
using namespace std;

struct augment_path {
  vector<vector<int> > g;
  vector<int> pa;  // pa 作为 pb 的映射对称关系存在， 主要利用 pa[v] != -1 来判断是否是 未匹配点
  vector<int> pb;  // 记录的是真正的，匹配的信息。
  vector<int> vis;  // 访问, 因为，找到的 Augment Path 需要是简单路径，不能有环， 所以， 需要 vis 数组。
  int n, m;         // 顶点和边的数量
  int dfn;          // 时间戳记
  int res;          // 匹配数

  augment_path(int _n, int _m) : n(_n), m(_m) {
    assert(0 <= n && 0 <= m);
    pa = vector<int>(n, -1);
    pb = vector<int>(m, -1);
    vis = vector<int>(n);
    g.resize(n);
    res = 0;
    dfn = 0;
  }

  void add(int from, int to) {
    assert(0 <= from && from < n && 0 <= to && to < m);
    g[from].push_back(to);
  }

  bool dfs(int v) {
    vis[v] = dfn;
    for (int u : g[v]) {
      if (pb[u] == -1) {
        pb[u] = v;
        pa[v] = u;
        return true;
      }
    }
    for (int u : g[v]) {
      if (vis[pb[u]] != dfn && dfs(pb[u])) {
        pa[v] = u;
        pb[u] = v;
        return true;
      }
    }
    return false;
  }

  int solve() {
    while (true) {
      dfn++;
      int cnt = 0;
      for (int i = 0; i < n; i++) {
        if (pa[i] == -1 && dfs(i)) {
          cnt++;
        }
      }
      if (cnt == 0) {
        break;
      }
      res += cnt;
    }
    return res;
  }
};

int main() {
  int n, m, e;
  cin >> n >> m >> e;
  augment_path solver(n, m);
  int u, v;
  for (int i = 0; i < e; i++) {
    cin >> u >> v;
    u--, v--;
    solver.add(u, v);
  }
  cout << solver.solve() << "\n";
  for (int i = 0; i < n; i++) {
    cout << solver.pa[i] + 1 << " ";
  }
  cout << "\n";
}