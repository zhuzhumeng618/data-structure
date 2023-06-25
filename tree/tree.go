// Package main 树
package main

import (
	"fmt"
	"sync"
	"testing"
)

type Node struct {
	key   int
	value interface{}
	left  *Node // 左子树
	right *Node // 右子树
}

// ItemBinarySearchTree 二叉搜索树item
type ItemBinarySearchTree struct {
	root *Node
	lock sync.RWMutex // 读写锁
}

// Insert 插入数据
func (bst *ItemBinarySearchTree) Insert(key int, value interface{}) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &Node{key, value, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

// 找到合适位置插入
func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// InOrderTraverse 以中序遍历访问所有节点
func (bst *ItemBinarySearchTree) InOrderTraverse(f func(value interface{})) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverse(bst.root, f)
}

// 中序遍历递归函数
func inOrderTraverse(n *Node, f func(value interface{})) {
	if n != nil {
		inOrderTraverse(n.left, f)
		f(n.value)
		inOrderTraverse(n.right, f)
	}
}

// PreOrderTraverse 前序遍历
func (bst *ItemBinarySearchTree) PreOrderTraverse(f func(value interface{})) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	preOrderTraverse(bst.root, f)
}

// 前序遍历递归函数
func preOrderTraverse(n *Node, f func(value interface{})) {
	if n != nil {
		f(n.value)
		preOrderTraverse(n.left, f)
		preOrderTraverse(n.right, f)
	}
}

// PostOrderTraverse 后序遍历
func (bst *ItemBinarySearchTree) PostOrderTraverse(f func(value interface{})) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	postOrderTraverse(bst.root, f)
}

// 后序遍历递归函数
func postOrderTraverse(n *Node, f func(value interface{})) {
	if n != nil {
		postOrderTraverse(n.left, f)
		postOrderTraverse(n.right, f)
		f(n.value)
	}
}

// Min 最小值
func (bst *ItemBinarySearchTree) Min() *interface{} {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	n := bst.root
	if n == nil {
		return nil
	}
	for {
		if n.left == nil {
			return &n.value
		}
		n = n.left
	}
}

// Max 最大值
func (bst *ItemBinarySearchTree) Max() *interface{} {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	n := bst.root
	if n == nil {
		return nil
	}
	for {
		if n.right == nil {
			return &n.value
		}
		n = n.right
	}
}

// Search 搜索
func (bst *ItemBinarySearchTree) Search(key int) bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return search(bst.root, key)
}

// 搜索实现
func search(n *Node, key int) bool {
	if n == nil {
		return false
	}
	if key < n.key {
		return search(n.left, key)
	}
	if key > n.key {
		return search(n.right, key)
	}
	return true
}

// Remove 删除
func (bst *ItemBinarySearchTree) Remove(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	remove(bst.root, key)
}

// 删除实现
func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}
	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}
	// key == node.key
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	leftmostrightside := node.right
	for {
		// find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	node.key, node.value = leftmostrightside.key, leftmostrightside.value
	node.right = remove(node.right, node.key)
	return node
}

// 字符串输出显示
func (bst *ItemBinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

// 打印树实现
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.right, level)
	}
}

// 测试部分
var bst ItemBinarySearchTree

// 填充树
func fillTree(bst *ItemBinarySearchTree) {
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")
	bst.Insert(6, "6")
	bst.Insert(1, "1")
	bst.Insert(3, "3")
	bst.Insert(5, "5")
	bst.Insert(7, "7")
	bst.Insert(9, "9")
}

// 判断两个切片是否相同
func isSameSlice(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// 测试插入
	fillTree(&bst)
	bst.String()

	bst.Insert(11, "11")
	bst.String()

	fmt.Println("------------------")

}

func TestInOrderTraverse(t *testing.T) {
	var result []string
	bst.InOrderTraverse(func(i interface{}) {
		result = append(result, fmt.Sprintf("%s", i))
	})
	if !isSameSlice(result, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}) {
		t.Errorf("遍历错误 %v", result)
	}
}

func TestPreOrderTraverse(t *testing.T) {
	var result []string
	bst.PreOrderTraverse(func(i interface{}) {
		result = append(result, fmt.Sprintf("%s", i))
	})
	if !isSameSlice(result, []string{"8", "4", "2", "1", "3", "6", "5", "7", "10", "9", "11"}) {
		t.Errorf("遍历错误 %v  %v", result, []string{"8", "4", "2", "1", "3", "6", "5", "7", "10", "9", "11"})
	}
}

func TestPostOrderTraverse(t *testing.T) {
	var result []string
	bst.PostOrderTraverse(func(i interface{}) {
		result = append(result, fmt.Sprintf("%s", i))
	})
	if !isSameSlice(result, []string{"1", "3", "2", "5", "7", "6", "4", "9", "11", "10", "8"}) {
		t.Errorf("遍历错误 %v  %v", result, []string{"1", "3", "2", "5", "7", "6", "4", "9", "11", "10", "8"})
	}
}

func TestMin(t *testing.T) {
	if fmt.Sprintf("%s", *bst.Min()) != "1" {
		t.Errorf("最小值应该是 1")
	}
}

func TestMax(t *testing.T) {
	if fmt.Sprintf("%s", *bst.Max()) != "11" {
		t.Errorf("最大值是 11")
	}
}

func TestSearch(t *testing.T) {
	if !bst.Search(1) || !bst.Search(8) || !bst.Search(11) {
		t.Errorf("搜索失败")
	}
}

func TestRemove(t *testing.T) {
	bst.Remove(1)
	if fmt.Sprintf("%s", *bst.Min()) != "2" {
		t.Errorf("最小值应该是 2")
	}
}
