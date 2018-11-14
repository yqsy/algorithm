<!-- TOC -->

- [1. 说明](#1-说明)

<!-- /TOC -->



<a id="markdown-1-说明" name="1-说明"></a>
# 1. 说明

参考:  
* https://zhuanlan.zhihu.com/json-tutorial
* http://www.ecma-international.org/publications/files/ECMA-ST/ECMA-404.pdf
* http://www.json.org/
* https://github.com/tidwall/gjson (go实现)
* http://json.org/example.html
* https://jsonlint.com/

6种数据类型

数据类型|实例值|解析思路|存储
-|-|-|-
null|null|n开头|nil
boolean|true false| t f 开头|bool
string|""|"开头|string
number|浮点数|默认|float64
array|[]|[开头|[]interface{}
object|{...}|{开头|map[string]interface{}

空白
* 空格符 " "
* 制表符 "\t"
* 回车符 "\r"
* 换行符 "\n"

