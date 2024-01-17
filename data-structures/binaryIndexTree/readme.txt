
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



找一个 BIT 解决逆序对的模板题目做一下。  BIT和逆序对是什么关系？ (在值域上创建 BIT, 然后统计逆序对的数量）
逆序对的计算，可以在merge sort 中 merge 的过程中作统计。 时间复杂度O(n*logn)
逆序对也可以通过 binary index tree 来作， BIT用来统计的是数组的值域。时间复杂度是 O(n*logm) m 是值域的最大值。需要查询和更新n次，每次 logm的时间复杂度。
