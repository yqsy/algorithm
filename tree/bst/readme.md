<!-- TOC -->

- [1. 说明](#1-说明)
- [2. 取整的思考方法](#2-取整的思考方法)

<!-- /TOC -->

<a id="markdown-1-说明" name="1-说明"></a>
# 1. 说明

>> 一棵二叉查找树(BST)是一棵二叉树,其中每个结点的键都`大于`其`左子树中`的任意结点的键而`小于` `右子树`的任意结点的键


实现
* size 某个结点的子树结点总数 / 树的所有结点数(包括根结点)
* get 查找
* put 插入
* min 最小键
* max 最大键
* floor 向下取整
* ceiling 向上取整
* select 查找排名为k的键
* rank 返回键的排名
* delete 删除键
* deleteMin 删除最小键 
* deleteMax 删除最大键
* keys 范围查找


<a id="markdown-2-取整的思考方法" name="2-取整的思考方法"></a>
# 2. 取整的思考方法

向下取整也就是寻找`<=key`的`最大值`

key在树中递归查找,3种情况:
1. node.key == key. 返回这个node,相等就是最大的了
2. node.key < key. 继续向node的右子树前进,寻找`小于等于key的最大值`
3. node.key > key. 向node的左子树前进,寻找小于key的情况

向上取整也就是寻找`>=key`的`最小值`

1. node.key == key. 返回这个node,相等就是最小的了
2. node.key < key. 向node的右子树前进,寻找大于key的情况
3. node.key > key. 继续向node的左子树前进,寻找`大于等于key的最小值`
