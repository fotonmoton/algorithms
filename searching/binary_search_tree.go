package searching

import "github.com/fotonmoton/algorithms/fundamentals/queue"

type bstNode[K any, V any] struct {
	left  *bstNode[K, V]
	right *bstNode[K, V]
	key   K
	val   V
	n     int64
}

// TODO: maybe pass pointers for recursive funcs?
type bst[K any, V any] struct {
	root *bstNode[K, V]
	cmp  func(*K, *K) int
}

func NewBST[K any, V any](cmp func(*K, *K) int) SymbolTable[K, V] {
	return &bst[K, V]{nil, cmp}
}

func (t *bst[K, V]) Put(key K, val V) {
	t.root = t.put(key, val, t.root)
}

func (t *bst[K, V]) put(key K, val V, node *bstNode[K, V]) *bstNode[K, V] {
	if node == nil {
		return &bstNode[K, V]{nil, nil, key, val, 1}
	}

	cmp := t.cmp(&key, &node.key)

	if cmp < 0 {
		node.left = t.put(key, val, node.left)
	}

	if cmp == 0 {
		node.val = val
	}

	if cmp > 0 {
		node.right = t.put(key, val, node.right)
	}

	node.n = t.size(node.left) + t.size(node.right) + 1
	return node
}

func (t *bst[K, V]) Get(key K) *V {
	return t.get(key, t.root)
}

func (t *bst[K, V]) get(key K, node *bstNode[K, V]) *V {
	if node == nil {
		return nil
	}

	cmp := t.cmp(&key, &node.key)

	if cmp < 0 {
		return t.get(key, node.left)
	}

	if cmp > 0 {
		return t.get(key, node.right)
	}

	return &node.val
}

func (t *bst[_, __]) Size() int64 {
	return t.size(t.root)
}

func (t *bst[K, V]) size(node *bstNode[K, V]) int64 {
	if node == nil {
		return 0
	}

	return node.n
}

func (t *bst[K, _]) Min() *K {
	if t.root == nil {
		return nil
	}

	return &t.min(t.root).key
}

func (t *bst[K, V]) min(node *bstNode[K, V]) *bstNode[K, V] {
	if node.left == nil {
		return node
	}

	return t.min(node.left)
}

func (t *bst[K, _]) Max() *K {
	if t.root == nil {
		return nil
	}

	return &t.max(t.root).key
}

func (t *bst[K, V]) max(node *bstNode[K, V]) *bstNode[K, V] {
	if node.right == nil {
		return node
	}

	return t.max(node.right)
}

func (t *bst[K, V]) Floor(key K) *K {
	largest := t.floor(key, t.root)

	if largest == nil {
		return nil
	}

	return &largest.key
}

func (t *bst[K, V]) floor(key K, node *bstNode[K, V]) *bstNode[K, V] {
	if node == nil {
		return nil
	}

	cmp := t.cmp(&key, &node.key)

	if cmp == 0 {
		return node
	}

	if cmp < 0 {
		return t.floor(key, node.left)
	}

	larger := t.floor(key, node.right)

	if larger != nil {
		return larger
	}

	return node
}

func (t *bst[K, V]) Ceiling(key K) *K {
	smallest := t.ceiling(key, t.root)

	if smallest == nil {
		return nil
	}

	return &smallest.key
}

func (t *bst[K, V]) ceiling(key K, node *bstNode[K, V]) *bstNode[K, V] {
	if node == nil {
		return nil
	}

	cmp := t.cmp(&key, &node.key)

	if cmp == 0 {
		return node
	}

	if cmp > 0 {
		return t.ceiling(key, node.right)
	}

	smaller := t.ceiling(key, node.left)

	if smaller != nil {
		return smaller
	}

	return node
}

func (t *bst[K, V]) Rank(key K) int64 {
	return t.rank(key, t.root)
}

func (t *bst[K, V]) rank(key K, node *bstNode[K, V]) int64 {
	if node == nil {
		return 0
	}

	cmp := t.cmp(&key, &node.key)

	// If we found key in a tree then left subtree
	// will always contain keys less than current node key
	// and right subtree will always ontain greater keys (by BST definition).
	// So we simply return left subtree size
	if cmp == 0 {
		return t.size(node.left)
	}

	// If current node key is bigger than key for which rank is searched
	// we should descend deeper in left subtree
	if cmp < 0 {
		return t.rank(key, node.left)
	}

	// If we found node with key that is less than search key
	// we get the size of the left subtree, add 1 to count current node in
	// rank value and descend deeper in right subtree.
	return 1 + t.size(node.left) + t.rank(key, node.right)
}

func (t *bst[K, V]) KeyByRank(i int64) *K {
	node := t.keyByRank(i, t.root)

	if node == nil {
		return nil
	}

	return &node.key
}

func (t *bst[K, V]) keyByRank(rank int64, node *bstNode[K, V]) *bstNode[K, V] {
	if node == nil {
		return nil
	}

	// We need left subtree size to substract it from our index
	// when we descend deeper in right subtree
	leftSize := t.size(node.left)

	if rank < leftSize {
		return t.keyByRank(rank, node.left)
	}

	if rank > leftSize {
		// We subtract left size subtree
		return t.keyByRank(rank-leftSize-1, node.right)
	}

	return node
}

func (t *bst[K, V]) Contains(key K) bool {
	return t.Get(key) == nil
}

func (t *bst[K, V]) IsEmpty() bool {
	return t.Size() == 0
}

func (t *bst[K, V]) DeleteMin() {

	if t.root == nil {
		return
	}

	t.root = t.deleteMin(t.root)
}

func (t *bst[K, V]) deleteMin(node *bstNode[K, V]) *bstNode[K, V] {
	if node.left == nil {
		return node.right
	}

	node.left = t.deleteMin(node.left)
	node.n = t.size(node.left) + t.size(node.right) + 1

	return node
}

func (t *bst[K, V]) DeleteMax() {

	if t.root == nil {
		return
	}

	t.root = t.deleteMax(t.root)
}

func (t *bst[K, V]) deleteMax(node *bstNode[K, V]) *bstNode[K, V] {
	if node.right == nil {
		return node.left
	}

	node.right = t.deleteMax(node.right)
	node.n = t.size(node.left) + t.size(node.right) + 1

	return node
}

func (t *bst[K, V]) Delete(key K) {
	t.root = t.delete(key, t.root)
}

func (t *bst[K, V]) delete(key K, node *bstNode[K, V]) *bstNode[K, V] {
	if node == nil {
		return nil
	}

	cmp := t.cmp(&key, &node.key)

	if cmp < 0 {
		node.left = t.delete(key, node.left)
	} else if cmp > 0 {
		node.right = t.delete(key, node.right)
	} else {

		// Shortcut: we can return left or right subtree if we have only one of them
		// without size recalculation and pointers juggling
		if node.right == nil {
			return node.left
		}
		if node.left == nil {
			return node.right
		}

		// Needed to delete "min" node in right subtree
		tmp := node
		// We substitute current node with "min" node from right subtree.
		// When "node" variable will be returned to the caller "tmp" node
		// will be erased by "node" value and be marked for garbage collection.
		// At least it should work as described
		node = t.min(tmp.right)
		// We prevent "node" duplication in the tree by deleting it from right subtree
		node.right = t.deleteMin(tmp.right)
		// Left subtree stays unchanged
		node.left = tmp.left
	}
	node.n = t.size(node.left) + t.size(node.right) + 1
	return node
}

func (t *bst[K, V]) KeysBetween(lo, hi K) []K {
	q := queue.NewQueue[K]()
	t.keysBetween(lo, hi, t.root, q)
	keys := make([]K, 0, q.Size())

	for !q.IsEmpty() {
		keys = append(keys, q.Dequeue())
	}

	return keys
}

func (t *bst[K, V]) keysBetween(lo, hi K, node *bstNode[K, V], q queue.Queue[K]) {
	if node == nil {
		return
	}
	cmplo := t.cmp(&lo, &node.key)
	cmphi := t.cmp(&hi, &node.key)

	if cmplo < 0 {
		t.keysBetween(lo, hi, node.left, q)
	}

	if cmplo <= 0 && cmphi >= 0 {
		q.Enqueue(node.key)
	}

	if cmphi > 0 {
		t.keysBetween(lo, hi, node.right, q)
	}
}

func (t *bst[K, V]) Keys() []K {
	if t.IsEmpty() {
		return []K{}
	}

	q := queue.NewQueue[K]()
	t.keysBetween(*t.Min(), *t.Max(), t.root, q)
	keys := make([]K, 0, q.Size())

	for !q.IsEmpty() {
		keys = append(keys, q.Dequeue())
	}

	return keys
}

func (t *bst[K, V]) SizeBetween(lo K, hi K) int64 {
	q := queue.NewQueue[K]()
	t.keysBetween(lo, hi, t.root, q)

	return int64(q.Size())
}
