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

	assert.Equal(t, 0, q.Size())
	assert.Equal(t, true, q.IsEmpty())
	assert.Equal(t, -1, q.TopIndex())
}

func TestIPQInsert(t *testing.T) {
	pq := NewIPQ(intDescending, 10)

	pq.Insert(1, 3)
	assert.Equal(t, 3, pq.Top())

	pq.Insert(2, 4)
	assert.Equal(t, 4, pq.Top())

	pq.Insert(3, 1)
	assert.Equal(t, 4, pq.Top())

	pq.Insert(4, 4)
	assert.Equal(t, 4, pq.Top())
}

func TestMoreIPQInsert(t *testing.T) {
	pq := NewIPQ(intDescending, 10)

	pq.Insert(1, 10)

	assert.Equal(t, 1, pq.TopIndex())
	assert.Equal(t, 10, pq.Top())
	assert.Equal(t, 1, pq.Size())
	assert.Equal(t, true, pq.Contains(1))
	assert.Equal(t, false, pq.IsEmpty())

	pq.Insert(2, 20)

	assert.Equal(t, 2, pq.TopIndex())
	assert.Equal(t, 20, pq.Top())
	assert.Equal(t, 2, pq.Size())
	assert.Equal(t, true, pq.Contains(2))
	assert.Equal(t, false, pq.IsEmpty())
}

func TestIPQRemove(t *testing.T) {
	pq := NewIPQ(intDescending, 10)

	pq.Insert(1, 10)

	assert.Equal(t, 1, pq.TopIndex())
	assert.Equal(t, 10, pq.Top())
	assert.Equal(t, 1, pq.Size())
	assert.Equal(t, true, pq.Contains(1))
	assert.Equal(t, false, pq.IsEmpty())

	pq.Insert(2, 20)

	assert.Equal(t, 2, pq.TopIndex())
	assert.Equal(t, 20, pq.Top())
	assert.Equal(t, 2, pq.Size())
	assert.Equal(t, true, pq.Contains(2))
	assert.Equal(t, false, pq.IsEmpty())

	removed := pq.Remove()

	assert.Equal(t, 20, removed)
	assert.Equal(t, 10, pq.Top())
	assert.Equal(t, 1, pq.Size())
	assert.Equal(t, false, pq.Contains(2))
	assert.Equal(t, true, pq.Contains(1))
	assert.Equal(t, false, pq.IsEmpty())

	removed = pq.Remove()

	assert.Equal(t, 10, removed)
	// TODO: should return nil?
	// assert.Equal(t, "", pq.Top())
	assert.Equal(t, 0, pq.Size())
	assert.Equal(t, false, pq.Contains(1))
	assert.Equal(t, true, pq.IsEmpty())
}

func TestIPQRemoveAtIndex(t *testing.T) {
	pq := NewIPQ(intDescending, 10)

	// Top -> 40 - 30 - 20 - 10
	pq.Insert(8, 10)
	pq.Insert(5, 20)
	pq.Insert(3, 30)
	pq.Insert(4, 40)

	assert.Equal(t, 40, pq.Top())
	assert.Equal(t, 4, pq.TopIndex())

	// Top -> 40 - 30 - 10
	removed := pq.RemoveAtIndex(5)

	assert.Equal(t, 20, removed)
	assert.Equal(t, 40, pq.Top())
	assert.Equal(t, 4, pq.TopIndex())

	// Top -> 30 - 10
	removed = pq.RemoveAtIndex(4)

	assert.Equal(t, 40, removed)
	assert.Equal(t, 30, pq.Top())
	assert.Equal(t, 3, pq.TopIndex())

	// Top -> 30 - 20 - 10
	pq.Insert(5, 20)

	assert.Equal(t, 30, pq.Top())
	assert.Equal(t, 3, pq.TopIndex())

	// Top -> 10
	removed = pq.RemoveAtIndex(3)
	assert.Equal(t, 30, removed)
	removed = pq.RemoveAtIndex(5)
	assert.Equal(t, 20, removed)

	assert.Equal(t, 10, pq.Top())
	assert.Equal(t, 8, pq.TopIndex())
	assert.Equal(t, 1, pq.Size())
	assert.Equal(t, false, pq.Contains(5))
	assert.Equal(t, false, pq.Contains(4))
	assert.Equal(t, false, pq.Contains(3))
	assert.Equal(t, true, pq.Contains(8))
	assert.Equal(t, false, pq.IsEmpty())
}

func TestIndexChange(t *testing.T) {
	pq := NewIPQ(intDescending, 10)

	pq.Insert(1, 9)
	pq.Insert(2, 8)
	pq.Insert(3, 12)

	assert.Equal(t, 12, pq.Top())

	pq.Change(3, 7)

	assert.Equal(t, 9, pq.Top())

	pq.Change(2, 10)

	assert.Equal(t, 10, pq.Top())
}

func TestMultiwayMerge(t *testing.T) {

	multiwayMerge := func(streams ...*strings.Reader) string {
		q := NewIPQ(func(t1, t2 rune) bool { return t1 > t2 }, len(streams))
		b := strings.Builder{}

		for i, stream := range streams {
			rune, _, _ := stream.ReadRune()
			q.Insert(i, rune)
		}

		for !q.IsEmpty() {
			b.WriteRune(q.Top())
			streamIndex := q.TopIndex()
			q.Remove()
			rune, _, err := streams[streamIndex].ReadRune()
			if err != io.EOF {
				q.Insert(streamIndex, rune)
			}
		}

		return b.String()
	}

	// ordered "streams"
	firstStream := strings.NewReader("ABCFGIIZ")
	secondStream := strings.NewReader("BDHPQQ")
	thirdStream := strings.NewReader("ABEFJN")

	expected := "AABBBCDEFFGHIIJNPQQZ"
	actual := multiwayMerge(firstStream, secondStream, thirdStream)

	assert.Equal(t, expected, actual)
}
