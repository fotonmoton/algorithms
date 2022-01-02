package priority_queue

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIPQ(t *testing.T) {
	q := NewIPQ(intDescending, 1)
	assert.NotNil(t, q)

	assert.Equal(t, 0, q.size())
	assert.Equal(t, true, q.isEmpty())
	assert.Equal(t, -1, q.topIndex())
	// TODO: maybe should return nil?
	// assert.Equal(t, 0, q.top())
}

func TestIPQInsert(t *testing.T) {
	pq := NewIPQ(intDescending, 10)

	pq.insert(1, 3)
	assert.Equal(t, 3, pq.top())

	pq.insert(2, 4)
	assert.Equal(t, 4, pq.top())

	pq.insert(3, 1)
	assert.Equal(t, 4, pq.top())

	pq.insert(4, 4)
	assert.Equal(t, 4, pq.top())
}

func TestMoreIPQInsert(t *testing.T) {
	pq := NewIPQ(intDescending, 10)

	pq.insert(1, 10)

	assert.Equal(t, 1, pq.topIndex())
	assert.Equal(t, 10, pq.top())
	assert.Equal(t, 1, pq.size())
	assert.Equal(t, true, pq.contains(1))
	assert.Equal(t, false, pq.isEmpty())

	pq.insert(2, 20)

	assert.Equal(t, 2, pq.topIndex())
	assert.Equal(t, 20, pq.top())
	assert.Equal(t, 2, pq.size())
	assert.Equal(t, true, pq.contains(2))
	assert.Equal(t, false, pq.isEmpty())
}

func TestIPQRemove(t *testing.T) {
	pq := NewIPQ(intDescending, 10)

	pq.insert(1, 10)

	assert.Equal(t, 1, pq.topIndex())
	assert.Equal(t, 10, pq.top())
	assert.Equal(t, 1, pq.size())
	assert.Equal(t, true, pq.contains(1))
	assert.Equal(t, false, pq.isEmpty())

	pq.insert(2, 20)

	assert.Equal(t, 2, pq.topIndex())
	assert.Equal(t, 20, pq.top())
	assert.Equal(t, 2, pq.size())
	assert.Equal(t, true, pq.contains(2))
	assert.Equal(t, false, pq.isEmpty())

	removed := pq.remove()

	assert.Equal(t, 20, removed)
	assert.Equal(t, 10, pq.top())
	assert.Equal(t, 1, pq.size())
	assert.Equal(t, false, pq.contains(2))
	assert.Equal(t, true, pq.contains(1))
	assert.Equal(t, false, pq.isEmpty())

	removed = pq.remove()

	assert.Equal(t, 10, removed)
	// TODO: should return nil?
	// assert.Equal(t, "", pq.top())
	assert.Equal(t, 0, pq.size())
	assert.Equal(t, false, pq.contains(1))
	assert.Equal(t, true, pq.isEmpty())
}

func TestIPQRemoveAtIndex(t *testing.T) {
	pq := NewIPQ(intDescending, 10)

	// top -> 40 - 30 - 20 - 10
	pq.insert(8, 10)
	pq.insert(5, 20)
	pq.insert(3, 30)
	pq.insert(4, 40)

	assert.Equal(t, 40, pq.top())
	assert.Equal(t, 4, pq.topIndex())

	// top -> 40 - 30 - 10
	removed := pq.removeAtIndex(5)

	assert.Equal(t, 20, removed)
	assert.Equal(t, 40, pq.top())
	assert.Equal(t, 4, pq.topIndex())

	// top -> 30 - 10
	removed = pq.removeAtIndex(4)

	assert.Equal(t, 40, removed)
	assert.Equal(t, 30, pq.top())
	assert.Equal(t, 3, pq.topIndex())

	// top -> 30 - 20 - 10
	pq.insert(5, 20)

	assert.Equal(t, 30, pq.top())
	assert.Equal(t, 3, pq.topIndex())

	// top -> 10
	removed = pq.removeAtIndex(3)
	assert.Equal(t, 30, removed)
	removed = pq.removeAtIndex(5)
	assert.Equal(t, 20, removed)

	assert.Equal(t, 10, pq.top())
	assert.Equal(t, 8, pq.topIndex())
	assert.Equal(t, 1, pq.size())
	assert.Equal(t, false, pq.contains(5))
	assert.Equal(t, false, pq.contains(4))
	assert.Equal(t, false, pq.contains(3))
	assert.Equal(t, true, pq.contains(8))
	assert.Equal(t, false, pq.isEmpty())
}

func TestIndexChange(t *testing.T) {
	pq := NewIPQ(intDescending, 10)

	pq.insert(1, 9)
	pq.insert(2, 8)
	pq.insert(3, 12)

	assert.Equal(t, 12, pq.top())

	pq.change(3, 7)

	assert.Equal(t, 9, pq.top())

	pq.change(2, 10)

	assert.Equal(t, 10, pq.top())
}

func TestMultiwayMerge(t *testing.T) {
	// ordered "streams"
	firstStream := strings.NewReader("ABCFGIIZ")
	secondStream := strings.NewReader("BDHPQQ")
	thirdStream := strings.NewReader("ABEFJN")

	allStreams := []*strings.Reader{firstStream, secondStream, thirdStream}

	expected := "AABBBCDEFFGHIIJNPQQZ"
	actual := ""

	pq := NewIPQ(func(t1, t2 string) bool { return t1 > t2 }, len(allStreams))

	for i, stream := range allStreams {
		rune, _, _ := stream.ReadRune()
		pq.insert(i, string(rune))
	}

	for !pq.isEmpty() {
		actual += string(pq.top())
		streamIndex := pq.topIndex()
		pq.remove()
		rune, _, err := allStreams[streamIndex].ReadRune()
		if err != io.EOF {
			pq.insert(streamIndex, string(rune))
		}
	}

	assert.Equal(t, expected, actual)
}
