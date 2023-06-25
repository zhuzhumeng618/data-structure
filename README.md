
# 冒泡排序

通过每一次遍历获取最大值，将最大值放入尾部，然后除开最大值，剩下的数据再进行遍历获取最大值。

```go
// Package main 冒泡排序
package main

import (
	"fmt"
)

// BubbleASort 升序
func BubbleASort(values []int) {
	l := len(values)
	// 有几个数就比较几次
	for i := 0; i < l-1; i++ {
		// 比较总长度，每次减一
		for j := i + 1; j < l; j++ {
			// 如果后面的比前面的小就交换，所以是升序
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}

// BubbleZSort 升序
func BubbleZSort(values []int) {
	l := len(values)
	// 有几个数就比较几次
	for i := 0; i < l-1; i++ {
		// 比较总长度，每次减一
		for j := i + 1; j < l; j++ {
			// 如果后面的比前面的小就交换，所以是升序
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}

func main() {
	// sli := []int{5, 4, 3, 2, 1}
	// BubbleASort(sli)
	// for i, v := range sli {
	// 	fmt.Printf("i: %v, v: %v\n", i, v)
	// }
	// fmt.Println("升序排序后: ", sli)

	sli := []int{1, 2, 3, 4, 5, 6}
	BubbleZSort(sli)
	for i, v := range sli {
		fmt.Printf("i: %v, v: %v\n", i, v)
	}
	fmt.Println("降序序排序后: ", sli)
}

```

# 选择排序

```go
// Package main 选择排序
package main

import (
	"fmt"
)

func SelectSort(values []int) {
	l := len(values)
	for i := 0; i < l-1; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if values[j] < values[min] {
				min = j
			}
		}

		// 交换值
		tmp := values[i]
		values[i] = values[min]
		values[min] = tmp
	}
}

func main() {
	sli := []int{5, 4, 3, 2, 1}
	SelectSort(sli)
	fmt.Println("升序排序后:", sli)
}
```

# 插入排序

```go
// Package main 插入排序
package main

import (
	"fmt"
)

func InsertSort(values []int) {
	l := len(values)
	// 总循环次数，元素个数
	for i := 0; i < l; i++ {
		// 从后面第二个开始，取出，找到合适的位置插入
		for j := i; j > 0; j-- {
			// 比较
			if values[j] > values[j-1] {
				// 交换
				tmp := values[j]
				values[j] = values[j-1]
				values[j-1] = tmp
			}
		}
	}
}

func main() {
	sli := []int{1, 2, 3, 4, 5, 6}
	InsertSort(sli)
	fmt.Println("降序排序后:", sli)
}

```

# 快速排序

```go
// Package main 快速排序
package main

import (
	"fmt"
)

func QuickSort(values []int) []int {
	l := len(values)

	if l < 1 {
		return values
	}

	// 假设第一个元素是中间值
	middle := values[0]

	// 左边元素
	var left []int

	// 右边元素
	var right []int

	for i := 1; i < l; i++ {
		if middle < values[i] {
			right = append(right, []int{values[i]}...)
		} else {
			left = append(left, []int{values[i]}...)
		}
	}

	middle_s := []int{middle}
	// 递归排左边
	left = QuickSort(left)

	// 递归排右边
	right = QuickSort(right)

	result := append(append(left, middle_s...), right...)

	return result
}

func main() {
	sli := []int{5, 4, 3, 2, 1}
	sli = QuickSort(sli)
	fmt.Println("升序排序后:", sli)
}

```

# 归并排序

```go
// Package main 归并排序
package main

import (
	"fmt"
)

func MergeSort(values []int, a int, b int) {
	// 长度小于等于 1，不用排序
	if b-a <= 1 {
		return
	}

	// 取中间值
	c := (a + b) / 2

	// 递归调用
	MergeSort(values, a, c) // 左边部分排序
	MergeSort(values, c, b) // 右边部分排序

	leftSli := make([]int, c-a)
	rightSli := make([]int, b-c)

	copy(leftSli, values[a:c])
	copy(rightSli, values[c:b])

	i := 0
	j := 0
	for k := a; k < b; k++ {
		switch {
		case i >= c-a:
			values[k] = rightSli[j]
			j++
		case j >= b-c:
			values[k] = leftSli[i]
			i++
		case leftSli[i] < rightSli[i]:
			values[k] = leftSli[i]
			i++
		default:
			values[k] = rightSli[j]
			j++
		}
	}
}

func main() {
	sli := []int{5, 4, 3, 2, 1}
	MergeSort(sli, 0, len(sli))
	fmt.Println("升序排序后:", sli)
}

```

# 二分查找

```go
// Package main 二分查找
package main

import (
	"fmt"
	"sort"
)

// BinarySearchFor 二分查找循环实现
func BinarySearchFor(sli []int, k int) int {
	// 低位高位 长度
	lo, hi := 0, len(sli)-1
	// 如果低位小于等于高位一直循环
	for lo <= hi {
		// 取中间位，向右移动一位等于除以 2
		m := (lo + hi) >> 1 // 等同于 (lo + hi) / 2
		if sli[m] < k {
			lo = m + 1
		} else if sli[m] > k {
			hi = m - 1
		} else {
			return m
		}
	}

	return -1 // 没找到是 -1
}

// BinarySearchRecursion 二分查找递归实现
func BinarySearchRecursion(sli *[]int, leftIndex int, rightIndex int, findVal int) {
	// 退出条件，递归必须写退出条件，不然栈会压爆
	if leftIndex > rightIndex {
		fmt.Println("没找到")
		return
	}

	middleIndex := (leftIndex + rightIndex) / 2
	// 如果值大于中间索引处的值
	if findVal > (*sli)[middleIndex] {
		// 将索引 +1
		BinarySearchRecursion(sli, middleIndex+1, rightIndex, findVal)
	} else if findVal < (*sli)[middleIndex] {
		BinarySearchRecursion(sli, leftIndex, middleIndex-1, findVal)
	} else {
		// 如果相等，则就是这个中间索引
		fmt.Printf("找到了，下标为: %v\n", middleIndex)
	}
}

func main() {
	// 二分查找前提：必须是有序
	sli := []int{8, 7, 6, 5, 4, 3, 2, 1}
	// 先排序
	sort.Ints(sli)
	fmt.Printf("sli: %v\n", sli) // 1, 2, 3, 4, 5, 6, 7, 8

	result := BinarySearchFor(sli, 5) // 返回下标数
	fmt.Println(result)

	BinarySearchRecursion(&sli, 0, len(sli), 8)
}

```

# 链表

```go
// Package main 链表
package main

import (
	"fmt"
)

type Object interface{}

type Node struct {
	Data Object // 定义数据域
	Next *Node  // 定义地址域（指向下一个表的地址）
}

type List struct {
	headNode *Node // 头节点
}

// IsEmpty 判断链表是否为空，如果头节点为空，则列表为空
func (l *List) IsEmpty() bool {
	if l.headNode == nil {
		return true
	}
	return false
}

// GetLength 获取列表长度
func (l *List) GetLength() int {
	current := l.headNode
	len := 0
	for current != nil {
		len++
		current = current.Next
	}
	return len
}

// Add 从链表头部添加元素
func (l *List) Add(obj Object) *Node {
	node := &Node{
		Data: obj,
	}
	node.Next = l.headNode
	l.headNode = node
	return node
}

// Append 从链表尾部添加元素
func (l *List) Append(obj Object) {
	node := &Node{
		Data: obj,
	}
	if l.IsEmpty() {
		l.headNode = node
	} else {
		current := l.headNode
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

// Insert 在链表指定位置添加元素
func (l *List) Insert(index int, obj Object) {
	if index < 0 {
		l.Add(obj)
	} else if index > l.GetLength() {
		l.Append(obj)
	} else {
		pre := l.headNode
		len := 0
		for len < (index - 1) {
			pre = pre.Next
			len++
		}
		node := &Node{Data: obj}
		node.Next = pre.Next
		pre.Next = node
	}
}

// Remove 删除链表指定值的元素
func (l *List) Remove(obj Object) {
	pre := l.headNode
	if pre.Data == obj {
		l.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == obj {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

// RemoveAtIndex 删除指定位置的元素
func (l *List) RemoveAtIndex(index int) {
	pre := l.headNode
	if index <= 0 {
		l.headNode = pre.Next
	} else if index > l.GetLength() {
		fmt.Println("超长了")
		return
	} else {
		len := 0
		for len != (index-1) && pre.Next != nil {
			len++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

// Contains 查看链表是否包含某个元素
func (l *List) Contains(obj Object) bool {
	current := l.headNode
	for current != nil {
		if current.Data == obj {
			return true
		}
		current = current.Next
	}
	return false
}

// LoopList 遍历链表所有节点
func (l *List) LoopList() {
	if !l.IsEmpty() {
		current := l.headNode
		for {
			fmt.Printf("current.Data: %v\n", current.Data)
			if current.Next != nil {
				current = current.Next
			} else {
				break
			}
		}
	}
}

func main() {
	list := List{}
	list.Append("Java")
	list.Append("Python")
	list.Append("Golang")
	fmt.Printf("list.GetLength(): %v\n", list.GetLength())
	list.LoopList()

	fmt.Println("------------------------------")
	list.Add("before")
	list.LoopList()

	fmt.Println("------------------------------")
	list.Remove("Java")
	list.LoopList()

	fmt.Println("------------------------------")
	list.RemoveAtIndex(2)
	list.LoopList()
}

```

# 栈

```go
// Package main 栈 后进先出
package main

import (
	"fmt"
)

// Stack 使用字符串切片存储数据
type Stack []string

// IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push 入栈、压栈
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Pop 出栈、弹栈
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index] // 获取栈顶元素
		*s = (*s)[:index]      // 删除元素
		return element, true   // 返回
	}
}

func main() {
	var stack Stack // 创建一个栈

	stack.Push("java")
	stack.Push("python")
	stack.Push("golang")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}

}

```

# 树

```go
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

```