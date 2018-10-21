<!-- TOC -->

- [1. 说明](#1-说明)
- [2. TODO](#2-todo)

<!-- /TOC -->


<a id="markdown-1-说明" name="1-说明"></a>
# 1. 说明

* https://flaviocopes.com/golang-data-structure-graph/
* https://wuyin.io/2018/06/22/golang-data-structure-graph/ (中文)
* https://www.zhihu.com/question/28549888/answer/41229881 (DFS和BFS的取舍)

G(V,E)

* V (vertex): 顶点集
* E (edge): 边集, 元素是一个二元组数组对,用(x,y)表示

实现:


* 添加顶点
* 添加边
* BFS 队列
* DFS 栈
* 打印 顶点 - 相连接的边

<a id="markdown-2-todo" name="2-todo"></a>
# 2. TODO

```go

// 1. 为什么edges 的 key是Node而不是*Node
type Graph struct {
	nodes []*Node          // 顶点集
	edges map[Node][]*Node // 边集
}

// 2. 为什么容器类的items使用Node而不是*Node

type NodeStack struct {
	items []Node
}

type NodeQueue struct {
	items []Node
}

```

