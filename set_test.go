package gods_test

import (
	"testing"

	gods "github.com/gweithio/go-ds"
)

func TestGoSet(t *testing.T) {
	stringSet := gods.MakeSet[string]()
	stringSet.Add("Hello")
	stringSet.Add("World")
	stringSet.Add("Again")

	if !stringSet.Contains("Hello") {
		t.Error("Set does not contain Hello")
	}

	stringSet.Remove("World")

	if stringSet.Contains("World") {
		t.Error("World still exists when it should be removed")
	}
}

func TestQueue(t *testing.T) {
	floatQueue := gods.MakeQueue[float32](5)

	floatQueue.Enqueue(1.1)
	floatQueue.Enqueue(2.2)
	floatQueue.Enqueue(3.4)

	if floatQueue.Size != 3 {
		t.Error("Queue must be greater than 3")
	}

	value, ok := floatQueue.Dequeue()

	if value != 3.4 {
		t.Error("Value that was Dequeued should be 3")
	}

	if !ok {
		t.Error("The queue is already empty")
	}

}
