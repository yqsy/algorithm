<!-- TOC -->

- [1. 说明](#1-说明)
- [2. 实现](#2-实现)
- [3. 标准库sort的分析](#3-标准库sort的分析)
- [4. benchmark](#4-benchmark)
- [5. js指令](#5-js指令)

<!-- /TOC -->


<a id="markdown-1-说明" name="1-说明"></a>
# 1. 说明

参考:

* http://bigocheatsheet.com/ (所有的复杂度)
* http://sortbenchmark.org/
* https://www.zhihu.com/question/55521100/answer/144987717 (实用的排序算法)
* https://visualgo.net/en/sorting (可视化)

<a id="markdown-2-实现" name="2-实现"></a>
# 2. 实现

* 快速排序[递归]
* 归并排序[递归]
* 堆排序
* 插入排序
* 希尔排序

<a id="markdown-3-标准库sort的分析" name="3-标准库sort的分析"></a>
# 3. 标准库sort的分析
* 堆排序 `深度 == 0`
* 快速排序 (三项切分)
* 希尔排序(插入排序) `元素 <= 12`

<a id="markdown-4-benchmark" name="4-benchmark"></a>
# 4. benchmark



<a id="markdown-5-js指令" name="5-js指令"></a>
# 5. js指令

```bash
npm init -y
npm install mocha chai lodash --save
npm test
```
