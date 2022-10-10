package benchmark

type Node struct {
	Next *Node
	Val  int
}

type Block struct {
	Next *Block
	Val  []int
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
	var result []int
	result = append(result, b.Val[startIdx:]...)

	var walk func(cur *Block, count int)
	walk = func(cur *Block, count int) {
		end := limit - count + 1

		if end < 1 {
			return
		}

		if end < 257 {
			result = append(result, cur.Val[:end]...)
		} else {
			result = append(result, cur.Val...)
		}

		count += 256

		if cur.Next != nil {
			walk(cur.Next, count)
		}
	}

	walk(b, len(result))

	return result

}

func bs(hay []int, needle int) int {
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
