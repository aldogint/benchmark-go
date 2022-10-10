package benchmark

type Node struct {
	Next *Node
	Val  int
}

type Block struct {
	Next *Block
	Val  [256]int
}

// func RangeNode(n *Node, limit int) []int {
// 	var result []int

// 	result = append(result, n.Val)
// 	count := 1

// 	for {
// 		if count >= limit || n.Next == nil {
// 			break
// 		}

// 		result = append(result, n.Next.Val)
// 	}

// 	return result
// }

func RangeNode(n *Node, limit int) (res []int) {
	var result []int
	var walk func(cur *Node, count int)

	walk = func(cur *Node, count int) {
		result = append(result, cur.Val)

		if cur.Next != nil && count < limit {
			walk(cur.Next, count+1)
		}
	}

	walk(n, 0)

	return result
}

func RangeBlock(b *Block, limit int, start int) []int {
	startIdx := bs(b.Val, start)

	if startIdx == -1 {
		return nil
	}

	result := make([]int, len(b.Val)-startIdx)
	copy(result, b.Val[startIdx:])

	var walk func(cur *Block, count int)
	walk = func(cur *Block, count int) {
		end := limit - count + 1

		if end < 257 {
			copy(result, b.Val[:end])
		}

		count += 256

		if cur.Next != nil {
			walk(cur.Next, count)
		}
	}

	return result

}

func bs(hay [256]int, needle int) int {
	low := 0
	high := len(hay) - 1

	for low <= high {
		med := (low + high) / 2

		if hay[med] == needle {
			return med
		}

		if hay[med] < needle {
			low = med + 1
		} else {
			high = med - 1
		}
	}

	return -1
}
