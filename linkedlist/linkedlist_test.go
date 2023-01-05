package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	{
		ll := Constructor()

		ll.AddAtHead(1)
		ll.AddAtTail(3)
		ll.AddAtIndex(1, 2)
		assert.Equal(t, ll.Get(1), 2)
		ll.DeleteAtIndex(1)
		assert.Equal(t, ll.Get(1), 3)
	}

	{
		ll := Constructor()

		ll.AddAtHead(2)
		ll.DeleteAtIndex(1)
		ll.AddAtHead(2)
		ll.AddAtHead(7)
		ll.AddAtHead(3)
		ll.AddAtHead(2)
		ll.AddAtHead(5)
		ll.AddAtTail(5)
		assert.Equal(t, ll.Get(5), 2)
		ll.DeleteAtIndex(6)
		ll.DeleteAtIndex(4)
	}

	{
		ll := Constructor()

		ll.AddAtHead(5)
		ll.AddAtIndex(1, 2)
		assert.Equal(t, ll.Get(1), 2)
		ll.AddAtHead(6)
		ll.AddAtTail(2)
		assert.Equal(t, ll.Get(3), 2)
		ll.AddAtTail(1)
		assert.Equal(t, ll.Get(5), -1)
		ll.AddAtHead(2)
		assert.Equal(t, ll.Get(2), 5)
		ll.AddAtHead(6)
	}
}
