<!-- TOC -->

- [1. 说明](#1-说明)
- [2. 解题思路](#2-解题思路)

<!-- /TOC -->


<a id="markdown-1-说明" name="1-说明"></a>
# 1. 说明

参考:  
算法的乐趣 - 19章节

相关网站:  
* https://www.websudoku.com/ (在线做题)

<a id="markdown-2-解题思路" name="2-解题思路"></a>
# 2. 解题思路

1. 最简单-穷举算法对每个单元格进行深度尝试
2. 为每个单元格建立候选列表,再穷举
3. 动态维护相关单元格的候选数列表,相关20格

---
* extra: 穷举时IsValids时间复杂度改为O(n),并且包含了上面的第2个

extra主要运用了`把数字放到下标内判断有无`的方式,使得原来每次移动一步要`对比24次`,现在变成`3次`.并且不用为每一个cell制造一个候选数列表,而是针对`行`,`列`,`格`建立一个全局的候选数字列表.