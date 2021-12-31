package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var intCompare = func(t1, t2 int) bool { return t1 < t2 }

func TestNew(t *testing.T) {
	pq := NewPQ(intCompare)
	assert.NotNil(t, pq)
}

func TestSize(t *testing.T) {
	pq := NewPQ(intCompare)
	assert.Equal(t, 0, pq.size())
	pq.insert(1)
	assert.Equal(t, 1, pq.size())
	pq.delete()
	assert.Equal(t, 0, pq.size())

}

func TestIsEmpty(t *testing.T) {
	pq := NewPQ(intCompare)
	assert.Equal(t, true, pq.isEmpty())
	pq.insert(1)
	assert.Equal(t, false, pq.isEmpty())
}

func TestInsert(t *testing.T) {
	pq := NewPQ(intCompare)
	pq.insert(1)
	assert.Equal(t, 1, pq.top())
	pq.insert(4)
	assert.Equal(t, 4, pq.top())
	pq.insert(3)
	assert.Equal(t, 4, pq.top())
	pq.insert(5)
	assert.Equal(t, 5, pq.top())
}

func TestDelete(t *testing.T) {
	pq := NewPQ(intCompare)
	pq.insert(1)
	pq.insert(2)
	pq.insert(6)
	pq.insert(5)
	pq.insert(4)
	pq.insert(3)
	assert.Equal(t, 6, pq.delete())
	assert.Equal(t, 5, pq.delete())
	assert.Equal(t, 4, pq.delete())
	assert.Equal(t, 3, pq.delete())
	assert.Equal(t, 2, pq.delete())
	assert.Equal(t, 1, pq.delete())
}
