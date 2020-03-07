package trees

import (
	"fmt"
	"testing"
)

func TestRBTree_InsertBalanceFixup(t *testing.T) {
	var rbt RBTree
	rbt.insertNode(9)
	rbt.insertNode(8)
	rbt.insertNode(7)
	rbt.insertNode(6)
	rbt.insertNode(5)
	rbt.insertNode(4)
	rbt.insertNode(3)
	rbt.insertNode(2)
	rbt.insertNode(1)
	deepth := rbt.GetDeepth()
	fmt.Println(deepth)

}

func BenchmarkRBTree_InsertBalanceFixup(b *testing.B) {
	b.Run("1000-9999", func(b *testing.B) {
		var rbt RBTree
		for i := 0; i < b.N; i++ {
			for s := 0; s < 1000; s++ {
				rbt.insertNode(int64(s))
			}
		}
		deepth := rbt.GetDeepth()
		fmt.Println(deepth)
	})

	b.Run("10w-9999", func(b *testing.B) {
		var rbt RBTree
		for i := 0; i < b.N; i++ {
			for s := 0; s < 100000; s++ {
				rbt.insertNode(int64(s))
			}
		}
		deepth := rbt.GetDeepth()
		fmt.Println(deepth)
	})
	b.Run("10000w-9999", func(b *testing.B) {
		var rbt RBTree
		for i := 0; i < b.N; i++ {
			for s := 0; s < 100000000; s++ {
				rbt.insertNode(int64(s))
			}
		}
		deepth := rbt.GetDeepth()
		fmt.Println(deepth)
	})


	b.Run("10w-9999---AVL", func(b *testing.B) {
		var tree *Node
		for i := 0; i < b.N; i++ {
			for s := 1; s < 100000; s++ {
				tree, _ = InsertAVL(nil, s, s)
			}
		}
		avl := GetHeightAVL(tree)
		fmt.Println(avl)

	})
	b.Run("10000w-9999---AVL", func(b *testing.B) {
		var tree *Node
		for i := 0; i < b.N; i++ {
			for s := 1; s < 100000000; s++ {
				tree, _ = InsertAVL(nil, s, s)
			}
		}
		avl := GetHeightAVL(tree)
		fmt.Println(avl)
	})

}
/**

goos: darwin
goarch: amd64
pkg: algorithm/trees
BenchmarkRBTree_InsertBalanceFixup
BenchmarkRBTree_InsertBalanceFixup/1000-9999
17
17
17
17
BenchmarkRBTree_InsertBalanceFixup/1000-9999-8         	   13407	     90763 ns/op
BenchmarkRBTree_InsertBalanceFixup/10w-9999
31
31
31
BenchmarkRBTree_InsertBalanceFixup/10w-9999-8          	      79	  15295231 ns/op
BenchmarkRBTree_InsertBalanceFixup/10000w-9999
50
BenchmarkRBTree_InsertBalanceFixup/10000w-9999-8       	       1	35471252664 ns/op
BenchmarkRBTree_InsertBalanceFixup/10w-9999---AVL
1
1
1
BenchmarkRBTree_InsertBalanceFixup/10w-9999---AVL-8    	     303	   3777580 ns/op
BenchmarkRBTree_InsertBalanceFixup/10000w-9999---AVL
1
BenchmarkRBTree_InsertBalanceFixup/10000w-9999---AVL-8 	       1	3815927046 ns/op
PASS

 */