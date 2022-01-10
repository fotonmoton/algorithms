package searching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func intCompare(a, b *int) int {
	if *a < *b {
		return -1
	}

	if *a > *b {
		return 1
	}

	return 0
}

func TestPut(t *testing.T) {
	table := NewBST[int, int](intCompare)

	table.Put(1, 10)
	table.Put(2, 20)

	assert.Equal(t, 10, *table.Get(1))
	assert.Equal(t, 20, *table.Get(2))

	// rewrite
	table.Put(1, 11)

	assert.Equal(t, 11, *table.Get(1))
	assert.Equal(t, 20, *table.Get(2))
	assert.Equal(t, int64(2), table.Size())

}

// TODO: test with delete
func TestGet(t *testing.T) {
	table := NewBST[int, int](intCompare)

	assert.Nil(t, table.Get(0))

	table.Put(1, 2)

	assert.Equal(t, 2, *table.Get(1))
}

// TODO: test with delete
func TestSize(t *testing.T) {
	table := NewBST[int, int](intCompare)

	assert.Equal(t, int64(0), table.Size())

	table.Put(1, 1)

	assert.Equal(t, int64(1), table.Size())

	table.Put(2, 2)

	assert.Equal(t, int64(2), table.Size())
}

// TODO: test with delete
func TestMin(t *testing.T) {
	table := NewBST[int, int](intCompare)

	assert.Nil(t, table.Min())

	table.Put(3, 3)

	assert.Equal(t, 3, *table.Min())

	table.Put(2, 2)

	assert.Equal(t, 2, *table.Min())

	table.Put(4, 4)

	assert.Equal(t, 2, *table.Min())

	table.Put(1, 1)

	assert.Equal(t, 1, *table.Min())
}

// TODO: test with delete
func TestMax(t *testing.T) {
	table := NewBST[int, int](intCompare)

	assert.Nil(t, table.Max())

	table.Put(1, 1)
	assert.Equal(t, 1, *table.Max())

	table.Put(2, 2)
	assert.Equal(t, 2, *table.Max())

	table.Put(5, 5)
	assert.Equal(t, 5, *table.Max())

	table.Put(4, 4)
	assert.Equal(t, 5, *table.Max())

	table.Put(3, 3)
	assert.Equal(t, 5, *table.Max())

	table.Put(5, 55)
	assert.Equal(t, 5, *table.Max())

	table.Put(6, 6)
	assert.Equal(t, 6, *table.Max())
}

// TODO: test with delete
func TestFloor(t *testing.T) {
	table := NewBST[int, int](intCompare)

	assert.Nil(t, table.Floor(0))

	table.Put(1, 1)
	assert.Equal(t, 1, *table.Floor(1))

	table.Put(5, 5)
	assert.Equal(t, 5, *table.Floor(5))
	assert.Equal(t, 1, *table.Floor(4))

	table.Put(4, 4)
	assert.Equal(t, 5, *table.Floor(5))
	assert.Equal(t, 4, *table.Floor(4))
	assert.Equal(t, 1, *table.Floor(3))
	assert.Nil(t, table.Floor(0))
}

// TODO: test with delete
func TestCeiling(t *testing.T) {
	table := NewBST[int, int](intCompare)

	assert.Nil(t, table.Ceiling(0))

	table.Put(5, 5)
	assert.Equal(t, 5, *table.Ceiling(5))

	table.Put(4, 4)
	assert.Equal(t, 4, *table.Ceiling(0))
	assert.Equal(t, 5, *table.Ceiling(5))

	table.Put(3, 3)
	assert.Equal(t, 3, *table.Ceiling(0))
	assert.Equal(t, 3, *table.Ceiling(1))
	assert.Equal(t, 3, *table.Ceiling(3))
	assert.Equal(t, 4, *table.Ceiling(4))
	assert.Equal(t, 5, *table.Ceiling(5))
	assert.Nil(t, table.Ceiling(6))
}

// TODO: test with delete
func TestRank(t *testing.T) {
	table := NewBST[int, int](intCompare)

	assert.Equal(t, int64(0), table.Rank(1))

	table.Put(0, 0)
	assert.Equal(t, int64(1), table.Rank(1))

	table.Put(1, 1)
	assert.Equal(t, int64(2), table.Rank(2))

	table.Put(4, 4)
	assert.Equal(t, int64(2), table.Rank(2))
	assert.Equal(t, int64(2), table.Rank(3))
	assert.Equal(t, int64(3), table.Rank(5))

	table.Put(2, 2)
	assert.Equal(t, int64(2), table.Rank(2))
	assert.Equal(t, int64(3), table.Rank(3))
	assert.Equal(t, int64(4), table.Rank(5))

	table.Put(3, 3)
	assert.Equal(t, int64(2), table.Rank(2))
	assert.Equal(t, int64(3), table.Rank(3))
	assert.Equal(t, int64(4), table.Rank(4))
	assert.Equal(t, int64(5), table.Rank(5))
}

// TODO: test with delete
func TestKeyByRank(t *testing.T) {
	table := NewBST[int, int](intCompare)

	assert.Nil(t, table.KeyByRank(1))

	table.Put(0, 0)
	assert.Nil(t, table.KeyByRank(1))
	assert.Equal(t, 0, *table.KeyByRank(table.Rank(0)))

	table.Put(5, 5)
	assert.Equal(t, 5, *table.KeyByRank(table.Rank(5)))
	assert.EqualValues(t, 1, table.Rank(*table.KeyByRank(1)))
}

func TestDeleteMin(t *testing.T) {
	table := NewBST[int, int](intCompare)

	table.DeleteMin()

	table.Put(0, 0)
	assert.EqualValues(t, 1, table.Size())

	table.DeleteMin()
	assert.EqualValues(t, 0, table.Size())

	table.Put(5, 5)
	table.Put(0, 0)
	table.Put(1, 1)
	table.Put(2, 2)

	assert.Equal(t, 0, *table.Get(0))

	table.DeleteMin()

	assert.Nil(t, table.Get(0))
	assert.EqualValues(t, 3, table.Size())
}

func TestDeleteMax(t *testing.T) {
	table := NewBST[int, int](intCompare)

	table.DeleteMin()

	table.Put(0, 0)
	assert.EqualValues(t, 1, table.Size())

	table.DeleteMax()
	assert.EqualValues(t, 0, table.Size())

	table.Put(0, 0)
	table.Put(5, 5)
	table.Put(1, 1)
	table.Put(2, 2)

	assert.Equal(t, 5, *table.Get(5))

	table.DeleteMax()

	assert.Nil(t, table.Get(5))
	assert.EqualValues(t, 3, table.Size())
}

// TODO: add more cases
func TestDelete(t *testing.T) {
	table := NewBST[int, int](intCompare)

	table.Delete(0)

	table.Put(0, 0)

	table.Delete(0)
	assert.EqualValues(t, 0, table.Size())
	assert.Nil(t, table.Get(0))

	table.Put(0, 0)
	table.Put(5, 5)
	table.Put(1, 1)
	table.Put(2, 2)

	assert.Equal(t, 1, *table.Get(1))

	table.Delete(1)
	assert.Nil(t, table.Get(1))
	assert.EqualValues(t, 3, table.Size())

	table.Delete(2)
	table.Delete(5)
	table.Delete(0)
	assert.EqualValues(t, 0, table.Size())
}

func TestKeysBetween(t *testing.T) {
	table := NewBST[int, int](intCompare)

	assert.EqualValues(t, []int{}, table.KeysBetween(0, 10))

	table.Put(1, 1)

	assert.EqualValues(t, []int{}, table.KeysBetween(2, 10))
	assert.EqualValues(t, []int{1}, table.KeysBetween(1, 1))

	table.Put(2, 2)
	table.Put(5, 5)

	assert.EqualValues(t, []int{5}, table.KeysBetween(3, 10))
	assert.EqualValues(t, []int{1, 2, 5}, table.KeysBetween(1, 5))
}

func TestKeys(t *testing.T) {
	table := NewBST[int, int](intCompare)

	assert.EqualValues(t, []int{}, table.Keys())

	table.Put(1, 1)

	assert.EqualValues(t, []int{1}, table.Keys())

	table.Put(2, 2)
	table.Put(5, 5)

	assert.EqualValues(t, []int{1, 2, 5}, table.Keys())

	table.Delete(2)

	assert.EqualValues(t, []int{1, 5}, table.Keys())

}

func TestSizeBetween(t *testing.T) {
	table := NewBST[int, int](intCompare)

	assert.EqualValues(t, 0, table.SizeBetween(0, 10))

	table.Put(1, 1)

	assert.EqualValues(t, 0, table.SizeBetween(2, 10))
	assert.EqualValues(t, 1, table.SizeBetween(1, 1))

	table.Put(2, 2)
	table.Put(5, 5)

	assert.EqualValues(t, 1, table.SizeBetween(3, 10))
	assert.EqualValues(t, 3, table.SizeBetween(1, 5))
}
