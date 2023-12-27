
/*
https://zhuanlan.zhihu.com/p/113042043

树上的倍增， 用空间换时间。

时间复杂度：
预处理 O(n*logn)  查询： O(logn)

*/


int Log2[MAXN], fa[MAXN][20], dep[MAXN]; // fa的第二维大小不应小于log2(MAXN)
bool vis[MAXN];
void dfs(int cur, int fath = 0) {
    if (vis[cur])
        return;
    vis[cur] = true;
    dep[cur] = dep[fath] + 1;
    fa[cur][0] = fath;
    for (int i = 1; i <= Log2[dep[cur]]; ++i)
        fa[cur][i] = fa[fa[cur][i - 1]][i - 1];
    for (int eg = head[cur]; eg != 0; eg = edges[eg].next)
        dfs(edges[eg].to, cur);
}
int lca(int a, int b) {
    if (dep[a] > dep[b])
        swap(a, b);
    while (dep[a] != dep[b])
        b = fa[b][Log2[dep[b] - dep[a]]];
    if (a == b)
        return a;
    for (int k = Log2[dep[a]]; k >= 0; k--)
        if (fa[a][k] != fa[b][k])
            a = fa[a][k], b = fa[b][k];
    return fa[a][0];
}

int main() {
    // ...
    for (int i = 2; i <= n; ++i)
        Log2[i] = Log2[i / 2] + 1;
    // ...
    dfs(s); // 无根树可以随意选一点为根
    // ...
    return 0;
}
