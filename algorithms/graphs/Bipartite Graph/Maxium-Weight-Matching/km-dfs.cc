
/*
KM 算法， DFS 版本的实现，

https://longrm.com/2018/05/05/2018-05-05-KM/

https://uoj.ac/submission/579212
时间复杂度是 O（n^4) , 测试用例会超时! 有  O(n^3) 的 BFS的解法
*/

#include <bits/stdc++.h>
using namespace std;

template <typename T>
struct hungarian {  // km
  int n;
  vector<int> matchx;
  vector<int> matchy;
  vector<int> pre;
  vector<bool> visx;
  vector<bool> visy;
  vector<T> lx;
  vector<T> ly;
  vector<vector<T> > g;
  vector<T> slack;
  T a;
  T inf;
  T res;
  queue<int> q;
  int org_n;
  int org_m;

  hungarian(int _n, int _m) {
    org_n = _n;
    org_m = _m;
    n = max(_n, _m);
    inf = numeric_limits<T>::max();
    res = 0;
    g = vector<vector<T> >(n, vector<T>(n));
    matchx = vector<int>(n, -1);
    matchy = vector<int>(n, -1);
    visx = vector<bool>(n);
    visy = vector<bool>(n);
    lx = vector<T>(n, -inf);
    ly = vector<T>(n);
  }

  void addEdge(int u, int v, int w) {
    g[u][v] = max(w, 0);  // 负值还不如不匹配 因此设为0不影响
  }

  bool dfs(int x) {
    visx[x] = true;

    for (int y=0; y <n; y++) {
        if (visy[y]) {
            continue;
        }
        T tmp = lx[x] + ly[y] - g[x][y];
        if (tmp == 0 ) {
            visy[y] = true; // 加入相等子图
            if (matchy[y] == -1 || dfs(matchy[y])) {
                matchx[x] = y;
                matchy[y] = x;
                return true;
            }
        } else {
            a =  min(tmp, a);
        }
    }
    return false;
  }

  void solve() {
    // 初始顶标
    for (int i = 0; i < n; i++) {
      for (int j = 0; j < n; j++) {
        lx[i] = max(lx[i], g[i][j]);
      }
    }

    for (int i = 0; i < n; i++) {
      a = inf;
      fill(visx.begin(), visx.end(), false);
      fill(visy.begin(), visy.end(), false);
      while(!dfs(i)) {
        for (int i=0; i<n; i++) {
            if (visx[i]) {
                lx[i] -= a;
            }
            if (visy[i]) {
                ly[i] += a;
            }
        }
        a = inf;
        fill(visx.begin(), visx.end(), false);
        fill(visy.begin(), visy.end(), false);
      }
    }

    // custom
    for (int i = 0; i < n; i++) {
      if (g[i][matchx[i]] > 0) {
        res += g[i][matchx[i]];
      } else {
        matchx[i] = -1;
      }
    }
    cout << res << "\n";
    for (int i = 0; i < org_n; i++) {
      cout << matchx[i] + 1 << " ";
    }
    cout << "\n";
  }
};

int main() {
  ios::sync_with_stdio(0), cin.tie(0);
  int n, m, e;
  cin >> n >> m >> e;

  hungarian<long long> solver(n, m);

  int u, v, w;
  for (int i = 0; i < e; i++) {
    cin >> u >> v >> w;
    u--, v--;
    solver.addEdge(u, v, w);
  }
  solver.solve();
}