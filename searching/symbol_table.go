package searching

// TODO: think about pointer semantics: where pointers should be used?
// Does go compiler silently convert values to pointers when they are leave table?
type SymbolTable[K any, V any] interface {
	Put(K, V)               // add value V with associated key K to symbol table
	Get(K) *V               // get value V with associated key K to symbol table, nil if value is absent
	Size() int64            // number of key-value pairs
	Min() *K                // smallest key
	Max() *K                // largest key
	Floor(K) *K             // largest key less than or equal to K
	Ceiling(K) *K           // smallest key greater or equal to K
	Rank(K) int64           // number of keys less than K. Rank(*Index(in)) = in
	KeyByRank(int64) *K     // key of specified rank. *Index(Rank(K)) = K
	Contains(K) bool        // check if key K exists in symbol table
	IsEmpty() bool          // check if symbol table is empty
	DeleteMin()             // delete value with smallest key
	DeleteMax()             // delete value with largest key
	Delete(K)               // delete value associated with key K.
	KeysBetween(K, K) []K   // keys between two other keys in sorted order
	Keys() []K              // all existing keys in sorted order
	SizeBetween(K, K) int64 // number of keys between two keys
}
