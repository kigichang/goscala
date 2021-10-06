package maps

type node[K comparable, V any] struct {
	key  K
	val  V
	prev *node[K, V]
	next *node[K, V]
}

func (n *node[K, V]) before(that *node[K, V]) {
	tmp := that.prev
	that.prev = n
	n.next = that
	n.prev = tmp
}

func (n *node[K, V]) after(that *node[K, V]) {
	tmp := that.next
	that.next = n
	n.prev = that
	n.next = tmp
}

type nodes[K comparable, V any] struct {
	compare func(K, K) int
	size int
	root *node[K, V]
	head *node[K, V]
}

func (n *nodes[K, V]) compareAndSet(a, b *node[K, V]) (ret int) {
	ret = n.compare(a.key, b.key)
	if ret == 0 {
		a.val = b.val
	}
	return
}

func newNodes[K comparable, V any](compare func(K, K) int, n ...node[K, V]) *nodes[K, V] {
	ns := &nodes[K, V] {
		compare: compare,
		size: 0,
		root: nil
		head: nil
	}

	return ns
}


func newNode[K comparable, V any](k K, v V) *node[K, V] {
	return &node[K, V]{
		key:  k,
		val:  v,
		prev: nil,
		next: nil,
	}
}
