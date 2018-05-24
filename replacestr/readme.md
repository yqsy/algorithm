<!-- TOC -->

- [1. 说明](#1-说明)
- [2. 编译](#2-编译)
- [3. 测试用例](#3-测试用例)

<!-- /TOC -->

<a id="markdown-1-说明" name="1-说明"></a>
# 1. 说明

参考:  
剑指offer - 面试题5 替换空格


<a id="markdown-2-编译" name="2-编译"></a>
# 2. 编译

```bash
g++ replacestr.cpp replacestr_test.cpp -g -o replacestr_test
./replacestr_test
```

<a id="markdown-3-测试用例" name="3-测试用例"></a>
# 3. 测试用例

* 空格在最前面/中间/最后面/连续多个空格
* 没有空格
* 特殊: 空指针,空字符串,只有一个空格
* 极值
