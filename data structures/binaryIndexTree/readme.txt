
https://oi-wiki.org/ds/fenwick/

Binary Index Tree （也成为 fenwick tree) :  树状数组
解决：
1. 前缀和 getsum(x) = a[1] + a[2] + ....+ a[x]
2. 单点修改的问题。  add(x, k)


getsum 的复杂度： O(lgn)
add 的复杂度： O（lgn)
创建BIT数的复杂度： O（n.lgn),  有 O（n)建树的技巧， 参看 https://oi-wiki.org/ds/fenwick/

但是， range update (l, r, +v) 的修改，就不如线段树， 但是也有相应的技巧， 譬如维护数组的 a[] 的差分数组， 这样， range update 时间上就可以变成，2个常数的 update： add(l, v)  add(r+1, -v)
具体看 https://oi-wiki.org/ds/fenwick/



2 Dimentions BIT : see Letcode 308