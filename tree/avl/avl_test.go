package avl

import (
	"testing"
	"fmt"
	"strconv"
)

func TestAvlSimplePrint(t *testing.T) {
	avl := &AVL{}

	for i := 0; i < 11; i ++ {
		avl.Put(i, strconv.Itoa(i))
	}

	fmt.Println(avl.Prettify())
}

// case1: 在30结点下插入的2种情况. + 70不存在的2种情况
//       100
//      / \
//     /   \
//     50   200
//    / \
//    30 70
//   /
//  10

//       50
//      / \
//     /   \
//     30   100
//    /   / \
//    10   70 200

func TestCase1(t *testing.T) {
	avl := &AVL{}
	avl.Put(100, "100")
	avl.Put(50, "50")
	avl.Put(200, "200")
	avl.Put(30, "30")
	avl.Put(70, "70")

	avl.Put(10, "10")

	if avl.head.key != 50 ||
		avl.head.left.key != 30 ||
		avl.head.right.key != 100 ||
		avl.head.left.left.key != 10 ||
		avl.head.right.left.key != 70 ||
		avl.head.right.right.key != 200 {
		t.Fatal("err")
	}
}

// case2: 在70结点下插入的2中情况. + 30不存在的2种情况
//       100
//      / \
//     /   \
//     50   200
//    / \
//    30 70
//       /
//      60

//      70
//      / \
//     /   \
//     50   100
//    / \   \
//    30 60   200
//

func TestCase2(t *testing.T) {
	avl := &AVL{}
	avl.Put(100, "100")
	avl.Put(50, "50")
	avl.Put(200, "200")
	avl.Put(30, "30")
	avl.Put(70, "70")

	avl.Put(60, "60")

	if avl.head.key != 70 ||
		avl.head.left.key != 50 ||
		avl.head.right.key != 100 ||
		avl.head.left.left.key != 30 ||
		avl.head.left.right.key != 60 ||
		avl.head.right.right.key != 200 {
		t.Fatal("err")
	}
}

// case4: 在250结点下插入的2种情况. + 150不存在的2种情况
//      100
//     / \
//    /   \
//    50   200
//       / \
//       150 250
//           /
//          230

//       200
//      /   \
//     /     \
//     100   250
//    / \    /
//    50 150 230

func TestCase4(t *testing.T) {
	avl := &AVL{}
	avl.Put(100, "100")
	avl.Put(50, "50")
	avl.Put(200, "200")
	avl.Put(150, "150")
	avl.Put(250, "250")

	avl.Put(230, "230")

	if avl.head.key != 200 ||
		avl.head.left.key != 100 ||
		avl.head.right.key != 250 ||
		avl.head.left.left.key != 50 ||
		avl.head.left.right.key != 150 ||
		avl.head.right.left.key != 230 {
		t.Fatal("err")
	}
}

// case3: 在150结点下插入的2种情况. + 250不存在的2种情况
//       100
//      / \
//     /   \
//     50   200
//        / \
//        150 250
//        /
//       120

//       150
//      /   \
//     /     \
//    100   200
//    /  \      \
//    50 120   250

func TestCase3(t *testing.T) {
	avl := &AVL{}
	avl.Put(100, "100")
	avl.Put(50, "50")
	avl.Put(200, "200")
	avl.Put(150, "150")
	avl.Put(250, "250")

	avl.Put(120, "120")

	if avl.head.key != 150 ||
		avl.head.left.key != 100 ||
		avl.head.right.key != 200 ||
		avl.head.left.left.key != 50 ||
		avl.head.left.right.key != 120 ||
		avl.head.right.right.key != 250 {
		t.Fatal("err")
	}
}

// 删除50 导致不平衡
//       100
//      / \
//     /   \
//     50   150
//        / \
//        120 200

//       150
//      / \
//     /   \
//     100   200
//      \
//      120

func TestDel1(t *testing.T) {
	avl := &AVL{}
	avl.Put(100, "100")
	avl.Put(50, "50")
	avl.Put(150, "150")
	avl.Put(120, "120")
	avl.Put(200, "200")

	avl.Delete(50)

	if avl.head.key != 150 ||
		avl.head.left.key != 100 ||
		avl.head.right.key != 200 ||
		avl.head.left.right.key != 120 {
		t.Fatal("err")
	}
}

// 删除150 导致不平衡
//       100
//      / \
//     /   \
//     50   150
//    / \
//    20 60

//       50
//      / \
//     /   \
//     20   100
//        /
//        60

func TestDel2(t *testing.T) {
	avl := &AVL{}
	avl.Put(100, "100")
	avl.Put(50, "50")
	avl.Put(150, "150")
	avl.Put(20, "20")
	avl.Put(60, "60")

	avl.Delete(150)

	if avl.head.key != 50 ||
		avl.head.left.key != 20 ||
		avl.head.right.key != 100 ||
		avl.head.right.left.key != 60 {
		t.Fatal("err")
	}
}


