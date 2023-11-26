package battle

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	{
		ints := newQueue[int](3)
		if ints.Len() != 0 {
			t.Fatalf("empty queue size is not 0")
		}

		ints.Push(10)
		ints.Push(20)
		if ints.Len() != 2 {
			t.Fatalf("expected the len to be 2")
		}
		if ints.Pop() != 10 {
			t.Fatalf("popped invalid value")
		}
		if ints.Pop() != 20 {
			t.Fatalf("popped invalid value")
		}

		for i := 0; i < 3; i++ {
			ints.Push(10 * i)
		}
		for i := 0; i < 3; i++ {
			if ints.Pop() != 10*i {
				t.Fatalf("popped invalid value")
			}
		}

		if ints.Len() != 0 {
			t.Fatalf("empty queue size is not 0")
		}

		rng := rand.New(rand.NewSource(time.Now().UnixNano()))

		for n := 0; n < 100; n++ {
			x := rng.Int()
			ints.Push(x)
			if ints.Len() != 1 {
				t.Fatalf("expected the queue to have length of 1 (have %d)", ints.Len())
			}
			if ints.Pop() != x {
				t.Fatalf("popped invalid value")
			}
			if ints.Len() != 0 {
				t.Fatalf("expected the queue to be empty")
			}
		}

		if ints.Len() != 0 {
			t.Fatalf("empty queue size is not 0")
		}

		for n := 0; n < 100; n++ {
			x := rng.Int()
			y := rng.Int()
			ints.Push(y)
			ints.Push(x)
			if ints.Len() != 2 {
				fmt.Printf("%v %v\n", ints.head, ints.tail)
				t.Fatalf("expected the queue to have length of 2 (have %d)", ints.Len())
			}
			if ints.Pop() != y {
				t.Fatalf("popped invalid value")
			}
			if ints.Pop() != x {
				t.Fatalf("popped invalid value")
			}
			if ints.Len() != 0 {
				t.Fatalf("expected the queue to be empty")
			}
		}

		if ints.Len() != 0 {
			t.Fatalf("empty queue size is not 0")
		}

		for n := 0; n < 100; n++ {
			x := rng.Int()
			y := rng.Int()
			z := rng.Int()
			ints.Push(x)
			ints.Push(y)
			ints.Push(z)
			if ints.Len() != 3 {
				fmt.Printf("%v %v\n", ints.head, ints.tail)
				t.Fatalf("expected the queue to have length of 3 (have %d)", ints.Len())
			}
			if ints.Pop() != x {
				t.Fatalf("popped invalid value")
			}
			if ints.Len() != 2 {
				t.Fatalf("expected the queue to have length of 2")
			}
			if ints.Pop() != y {
				t.Fatalf("popped invalid value")
			}
			if ints.Pop() != z {
				t.Fatalf("popped invalid value")
			}
			if ints.Len() != 0 {
				t.Fatalf("expected the queue to be empty")
			}
		}

		if ints.Len() != 0 {
			t.Fatalf("empty queue size is not 0")
		}
	}

	for sizeTest := 2; sizeTest < 20; sizeTest++ {
		ints := newQueue[int](sizeTest)
		if ints.Len() != 0 {
			t.Fatalf("empty queue size is not 0")
		}

		rng := rand.New(rand.NewSource(time.Now().UnixNano()))

		for n := 0; n < 120; n++ {
			num := rng.Intn(ints.Cap())
			var values []int
			for i := 0; i < num; i++ {
				values = append(values, rng.Int())
			}
			for _, v := range values {
				ints.Push(v)
			}
			for i, v := range values {
				if ints.Len() != len(values)-i {
					t.Fatalf("values[%d] cap=%d: invalid length: have %d, want %d", i, cap(values), ints.Len(), len(values)-i)
				}
				have := ints.Pop()
				if v != have {
					t.Fatalf("values[%d] mismatch: have %d, want %d", i, v, have)
				}
				if ints.Len()+1 != len(values)-i {
					t.Fatalf("values[%d] cap=%d: invalid length: have %d, want %d", i, cap(values), ints.Len(), len(values)-i)
				}
			}
		}

		if ints.Len() != 0 {
			t.Fatalf("empty queue size is not 0")
		}
	}
}
