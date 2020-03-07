package trees

import (
	"errors"
	"fmt"
)

var (
	errNotExist       = errors.New("index is not existed")
	errTreeNil        = errors.New("tree is null")
	errTreeIndexExist = errors.New("tree index is existed")
)

type Node struct {
	lchild *Node
	rchild *Node
	height int
	index  int
	data   int
}

func max(data1 int, data2 int) int {
	if data1 > data2 {
		return data1
	}
	return data2
}
func GetHeightAVL(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

// 左旋转
//
//    node  BF = 2
//       \
//         prchild     ----->       prchild    BF = 1
//           \                        /   \
//           pprchild               node  pprchild
func llRoation(node *Node) *Node {
	flagTP := node.rchild
	node.rchild = flagTP.lchild
	flagTP.lchild = node
	//todo 设置高度。。。。。。。。。。
	//update node height
	node.height = max(GetHeightAVL(node.lchild), GetHeightAVL(node.rchild)) + 1
	//update father node height
	flagTP.height = max(GetHeightAVL(flagTP.lchild), GetHeightAVL(flagTP.rchild)) + 1
	return flagTP
}

// 右旋转
//             node  BF = -2
//              /
//         plchild     ----->       plchild    BF = 1
//            /                        /   \
//        pplchild                lchild   node
func rrRoation(node *Node) *Node {
	flagTP := node.lchild
	node.lchild = flagTP.rchild
	flagTP.rchild = node
	//todo 设置高度。。。。。。。。。。
	//update node height
	node.height = max(GetHeightAVL(node.lchild), GetHeightAVL(node.rchild)) + 1
	//update father node height
	flagTP.height = max(GetHeightAVL(flagTP.rchild), GetHeightAVL(flagTP.lchild)) + 1
	return flagTP
}

// 先左转再右转
//          node                  node
//         /            左          /     右
//      node1         ---->    node2     --->         node2
//          \                   /                     /   \
//          node2s           node1                 node1  node
func lrRoation(node *Node) *Node {
	roation := llRoation(node.lchild)
	node.lchild = roation
	return rrRoation(node)
}

// 先右转再左转
//       node                  node
//          \          右         \         左
//          node1    ---->       node2     --->      node2
//          /                       \                /   \
//        node2                    node1           node  node1
func rlRoation(node *Node) *Node {
	roation := rrRoation(node.rchild)
	node.rchild = roation
	return llRoation(node)
}

//处理节点高度问题
func handleBF(node *Node) *Node {
	if GetHeightAVL(node.lchild)-GetHeightAVL(node.rchild) == 2 {

		if GetHeightAVL(node.lchild.lchild)-GetHeightAVL(node.lchild.rchild) > 0 { //RR
			node = rrRoation(node)
		} else {
			node = lrRoation(node)
		}

	} else if GetHeightAVL(node.lchild)-GetHeightAVL(node.rchild) == -2 {

		if GetHeightAVL(node.rchild.lchild)-GetHeightAVL(node.rchild.rchild) < 0 { //LL
			node = llRoation(node)
		} else {
			node = rlRoation(node)
		}

	}
	return node
}

func InsertAVL(node *Node, index int, data int) (*Node, error) {
	if node == nil {
		return &Node{lchild: nil, rchild: nil, index: index, data: data, height: 1}, nil
	}

	if node.index > index {
		node.lchild, _ = InsertAVL(node.lchild, index, data)
		node = handleBF(node)
	} else if node.index < index {
		node.rchild, _ = InsertAVL(node.rchild, index, data)
		node = handleBF(node)
	} else {
		return nil, errTreeIndexExist
	}
	node.height = max(GetHeightAVL(node.lchild), GetHeightAVL(node.rchild)) + 1
	return node, nil
}

//删除指定index节点
//查找节点 ---> 删除节点 ----> 调整树结构
//删除节点时既要遵循二叉搜索树的定义又要符合二叉平衡树的要求   ---> 重点处理删除节点的拥有左右子树的情况
func Delete(node *Node, index int) (*Node, error) {
	if node == nil {
		return nil, errNotExist
	}

	if node.index == index {

		//如果没有左子树或者右子树 --->直接返回nil
		if node.lchild == nil && node.rchild == nil {
			return nil, nil
		} else if node.lchild == nil || node.rchild == nil { //若只存在左子树或者右子树

			if node.lchild != nil {
				return node.lchild, nil
			} else {
				return node.rchild, nil
			}
		} else { //左右子树都存在

			//查找前驱，替换当前节点,然后再进行依次删除  ---> 节点删除后，前驱替换当前节点 ---> 需遍历到最后，调整平衡度
			var n *Node
			n = node.lchild
			for {
				if n.rchild == nil {
					break
				}
				n = n.rchild
			}
			n.data, node.data = node.data, n.data
			n.index, node.index = node.index, n.index
			node.lchild, _ = Delete(node.lchild, n.index)
		}

	} else if node.index > index {
		node.lchild, _ = Delete(node.lchild, index)
	} else {
		node.rchild, _ = Delete(node.rchild, index)
	}

	//update node's height
	node.height = max(GetHeightAVL(node.lchild), GetHeightAVL(node.rchild)) + 1
	//update the tree of balance
	return handleBF(node), nil
}

func search(node *Node, index int) (*Node, error) {
	for {
		if node == nil {
			return nil, errNotExist
		}

		if node.index == index {
			return node, nil
		} else if index > node.index {
			node = node.rchild
		} else {
			node = node.lchild
		}
	}

}

//中序遍历树，并根据钩子函数处理数据
func Midtraverse(node *Node, handle func(interface{}) error) error {
	if node == nil {
		return nil
	} else {
		if err := handle(node); err != nil {
			return err
		}
		if err := Midtraverse(node.lchild, handle); err != nil { //处理左子树
			return err
		}
		if err := Midtraverse(node.rchild, handle); err != nil { //处理右子树
			return err
		}
	}
	return nil
}

func PrintAVL() {
	//打印匿名函数
	f := func(node interface{}) error {
		fmt.Println(node)
		//fmt.Printf("this node is tree root node. node.index:%d node.data:%d\n", node.(*Node).index, node.(*Node).data)
		return nil
	}

	tree, _ := InsertAVL(nil, 3, 6)
	tree, _ = InsertAVL(tree, 4, 7)
	tree, _ = InsertAVL(tree, 5, 8)
	tree, _ = InsertAVL(tree, 7, 10)
	tree, _ = InsertAVL(tree, 6, 11)
	tree, _ = InsertAVL(tree, 15, 12)
	fmt.Println("==============")

	if err := Midtraverse(tree, f); err != nil {
		fmt.Printf("Midtraverse failed err:%v\n", err)
	}
}
