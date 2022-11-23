package _3_array_list

type LinkedNode[T any] struct {
	element T
	pre     *LinkedNode[T]
	post    *LinkedNode[T]
}

func NewNode[T any](v T) *LinkedNode[T] {
	return &LinkedNode[T]{
		element: v,
	}
}

// 头尾节点固定，里面值为nil
type DoubleLinkedList[T any] struct {
	length int
	head   *LinkedNode[T]
	tail   *LinkedNode[T]
}

func NewDLinkList[T any]() *DoubleLinkedList[T] {
	var empty T
	h := NewNode[T](empty)
	t := NewNode[T](empty)
	h.post = t
	t.pre = h
	return &DoubleLinkedList[T]{
		length: 0,
		head:   h,
		tail:   t,
	}
}

// 插入操作，在尾部插入
func (this *DoubleLinkedList[T]) insert(node *LinkedNode[T]) {
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
func (this *DoubleLinkedList[T]) delete(node *LinkedNode[T]) bool {
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
	queue    *DoubleLinkedList[KV]
	hashMap  map[int]*LinkedNode[KV]
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		queue:    NewDLinkList[KV](),
		hashMap:  make(map[int]*LinkedNode[KV]),
	}

}

func (this *LRUCache) Get(key int) int {
	node, ok := this.hashMap[key]
	//没有key对应的情况
	if !ok {
		return -1
	}
	queue := this.queue
	queue.delete(node)
	queue.insert(node)
	return node.element.value
}

func (this *LRUCache) Put(key int, value int) {
	queue := this.queue
	//容量过大
	if queue.length >= this.capacity {
		head := queue.head
		firstNode := head.post
		queue.delete(firstNode)
		delete(this.hashMap, firstNode.element.key)

	}
	node := NewNode[KV](KV{
		key:   key,
		value: value,
	})
	this.hashMap[key] = node
	queue.insert(node)
}
