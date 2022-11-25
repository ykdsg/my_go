package _3_array_list

type LinkedNode struct {
	element interface{}
	pre     *LinkedNode
	post    *LinkedNode
}

func NewNode(v interface{}) *LinkedNode {
	return &LinkedNode{
		element: v,
	}
}

// 头尾节点固定，里面值为nil
type DoubleLinkedList struct {
	length int
	head   *LinkedNode
	tail   *LinkedNode
}

func NewDLinkList() *DoubleLinkedList {
	h := NewNode(nil)
	t := NewNode(nil)
	h.post = t
	t.pre = h
	return &DoubleLinkedList{
		length: 0,
		head:   h,
		tail:   t,
	}
}

// 插入操作，在尾部插入
func (this *DoubleLinkedList) insert(node *LinkedNode) {
	if node == nil {
		return
	}
	tail := this.tail
	preLast := tail.pre
	preLast.post = node
	node.pre = preLast
	node.post = tail
	tail.pre = node
	this.length += 1
}
func (this *DoubleLinkedList) delete(node *LinkedNode) bool {
	if node == nil {
		return false
	}
	if this.head == node || this.tail == node {
		return false
	}

	preNode := node.pre
	postNode := node.post
	preNode.post = postNode
	postNode.pre = preNode
	this.length -= 1
	return true
}

type KV struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	queue    *DoubleLinkedList
	hashMap  map[int]*LinkedNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		queue:    NewDLinkList(),
		hashMap:  make(map[int]*LinkedNode),
	}

}

func (this *LRUCache) Get(key int) int {
	node, ok := this.hashMap[key]
	//没有key对应的情况
	if !ok {
		return -1
	}
	this.makeRecently(node)
	return (node.element).(KV).value
}

func (this *LRUCache) makeRecently(node *LinkedNode) {
	queue := this.queue
	queue.delete(node)
	queue.insert(node)
}

func (this *LRUCache) Put(key int, value int) {
	queue := this.queue
	hashMap := this.hashMap

	oldNode, ok := hashMap[key]
	//需要区分2种情况:1.已存在的key，2.新加入的key
	if ok {
		queue.delete(oldNode)
	} else { //新加入的key，需要考虑容量是否上限的情况
		//容量过大
		if queue.length >= this.capacity {
			this.removeLeast()
		}
	}
	this.addRecently(key, value)
}

// 删除最久没使用的元素
func (this *LRUCache) removeLeast() {
	queue := this.queue
	hashMap := this.hashMap
	firstNode := queue.head.post
	queue.delete(firstNode)
	delete(hashMap, (firstNode.element).(KV).key)
}

// 添加最近使用元素
func (this *LRUCache) addRecently(key int, value int) {
	queue := this.queue
	hashMap := this.hashMap
	newNode := NewNode(KV{
		key:   key,
		value: value,
	})
	hashMap[key] = newNode
	queue.insert(newNode)
}
