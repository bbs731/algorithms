分组循环  (解决简单题的神器)
适用场景：按照题目要求，数组会被分割成若干组，且每一组的判断/处理逻辑是一样的。

核心思想：
外层循环负责遍历组之前的准备工作（记录开始位置），和遍历组之后的统计工作（更新答案最大值）。
内层循环负责遍历组，找出这一组最远在哪结束。

https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solutions/2528771/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/

模板题目 2760

n = len(nums)
i = 0
while i < n:
    start = i
    while i < n and ...:
        i += 1

    # 从 start 到 i-1 是一组
    # 下一组从 i 开始，无需 i += 1




https://leetcode.cn/problems/longest-alternating-subarray/solutions/2615916/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-r57bz/

1446. 连续字符 1165
1869. 哪种连续子字符串更长 1205
1957. 删除字符使字符串变好 1358
2038. 如果相邻两个颜色均相同则删除当前颜色
2110. 股票平滑下跌阶段的数目 1408
1759. 统计同质子字符串的数目
228. 汇总区间
2760. 最长奇偶子数组 1420
1887. 使数组元素相等的减少操作次数 1428
2038. 如果相邻两个颜色均相同则删除当前颜色 1468
1759. 统计同质子字符串的数目 1491
1578. 使绳子变成彩色的最短时间 1574
1839. 所有元音按顺序排布的最长子字符串 1580
2765. 最长交替子序列


