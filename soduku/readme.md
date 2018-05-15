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

本实例三种方式都会做一遍