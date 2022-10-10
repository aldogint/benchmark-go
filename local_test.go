package benchmark

import (
	"testing"
)

func BenchmarkRangeNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := setupNode(i)
		b.StartTimer()
		RangeNode(data, i)
	}
}

func BenchmarkRangeBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := setupBlock(i)
		b.StartTimer()
		RangeBlock(data, i, 128)
	}
}

func setupBlock(n int) *Block {
	start := 1
	end := 256

	if n < 256 {
		end = 256
	}

	startBlock := &Block{}

	curBlock := startBlock

	for {
		if n < start {
			break
		}

		var val []int
		for i := start; i <= end; i++ {
			val = append(val, i)
		}

		curBlock.Val = val

		if n <= end {
			break
		}

		start = end + 1
		end = start + 255

		if n < end {
			end = n
		}

		nextBlock := &Block{}
		curBlock.Next = nextBlock
		curBlock = nextBlock
	}

	return startBlock
}

func setupNode(n int) *Node {
	start := &Node{
		Val: 0,
	}

	cur := start

	for i := 1; i <= n; i++ {
		next := &Node{
			Val: i,
		}
		cur.Next = next
		cur = next
	}

	return start
}
