/*
https://webcache.googleusercontent.com/search?q=cache:kSe9PVclIDMJ:https://longrm.com/2018/05/05/2018-05-05-KM/&cd=8&hl=zh-CN&ct=clnk&gl=hk

https://longrm.com/2018/05/05/2018-05-05-KM/
*/

#include <bits/stdc++.h>
using namespace std;

template <typename T>
struct hungarian {  // km
  int n;
  vector<int> matchx;  // 左集合对应的匹配点
  vector<int> matchy;  // 右集合对应的匹配点
  vector<int> pre;     // 连接右集合的左点
  vector<bool> visx;   // 拜访数组 左
  vector<bool> visy;   // 拜访数组 右
  vector<T> lx;
  vector<T> ly;
  vector<vector<T> > g;
  T inf;
  T res;
  vector<int> queue;
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
    pre = vector<int>(n);
    visx = vector<bool>(n);
    visy = vector<bool>(n);
    lx = vector<T>(n, -inf);
    ly = vector<T>(n);
  }

  void addEdge(int u, int v, int w) {
    g[u][v] = max(w, 0);  // 负值还不如不匹配 因此设为0不影响
  }

  void bfs(int startX) {
    bool find = false;
    int endY = -1;

    queue = vector<int>(n);
    int qs =0, qe= 0;
    T a = inf;
    queue[qe++] = startX;

    while (true) {
      while (qs < qe && !find) {
        int x = queue[qs++];
        visx[x] = true;
        for (int y = 0; y < n; y++) {
          T tmp = lx[x] + ly[y] - g[x][y];
          if (tmp == 0) {
            if (visy[y]) {
                continue;
            }
            visy[y] = true;
            pre[y] = x;

            if (matchy[y] == -1) {
                endY = y;
                find = true;
                break;
            } else {
                queue[qe++] = matchy[y];
            }
          } else {
            a = min(a, tmp);
          }
        }
      }

      if (find) {
        break;
      }
      qs = qe = 0;
      // 没有增广路 修改顶标
      for (int j = 0; j < n; j++) {
        if (visx[j]) {  // S
          lx[j] -= a;
          queue[qe++] = j;
        }
        if (visy[j]) {  // T
          ly[j] += a;
        }
      }
      a = inf;
    }

    while (endY != -1) {
        int preX = pre[endY], preY =matchx[preX];
        matchx[preX] = endY;
        matchy[endY] = preX;
        endY = preY;
    }

  }

  void solve() {
    // 初始顶标
    for (int i = 0; i < n; i++) {
      for (int j = 0; j < n; j++) {
        lx[i] = max(lx[i], g[i][j]);
      }
    }

    for (int i = 0; i < n; i++) {
      fill(visx.begin(), visx.end(), false);
      fill(visy.begin(), visy.end(), false);
      bfs(i);
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
