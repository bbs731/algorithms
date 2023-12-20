后缀数组和后缀树的关系：

1. 后缀数组可以通过对后缀树做深度优先遍历（DFT: Depth First Traversal）来进行构建，对所有的边（Edge）做字典序（Lexicographical Order）遍历。
2. 通过使用后缀集合和最长公共前缀数组（LCP Array: Longest Common Prefix Array）来构建后缀树，可在 O(n) 时间内完成，例如使用 Ukkonen 算法。构造后缀数组同样也可以在 O(n) 时间内完成。
3. 每个通过后缀树解决的问题，都可以通过组合使用后缀数组和额外的信息（例如：LCP Array）来解决。

总结一下：
后缀树 和 后缀数组是等价的。

后缀数组可以在 O(n) 的时间内通过 SA-IS 或者 DC3算法构造, 得到 sa[] 和 rank[] 数组。  height[] 数组可以通过不等式 height[rank[i]] >= height[rank[i-1]]-1 在 O(n) 时间复杂度（最坏是 2n) 内得到。 有了 sa[], rank[], height[] 数组，就等价有了后缀树。作用是一样的，等价的， 能解决同样一类问题。

后缀树的构造复杂，还没见过实现的代码，没必要学习。


后缀树 vs 后缀自动机 是什么关系？


GSAM 广义后缀自动机： 用来处理多串的问题。

