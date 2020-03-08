package trees

/**
红黑树的实现
*/

//用来标记是红节点、黑色节点
const (
	RED   int = 0
	BLACK int = 1
)

type RBNode struct {
	value               int64
	color               int
	left, right, parent *RBNode
}

//树的节点  只包含一个红黑节点
type RBTree struct {
	root   *RBNode
	cur    *RBNode
	create *RBNode
}

const (
	//右旋
	RIGHTROTATE int = 10
	//左旋
	LEFTROTATE int = 11
)

/**
    插入情况 大体四种
1、红黑树为空
2、插入节点的key已经存在了
3、插入节点的父节点为黑色节点
4、插入节点的父节点为红色节点
	4-1、叔叔节点存在且为红色
	4-2、叔叔节点不存在或为黑色，且插入节点的父节点为祖父节点的左子节点。
		4-2-1 插入节点为父节点的左子节点
		4-2-2 插入节点为父节点的右子节点
	4-3、叔叔节点不存在或为黑色，且插入节点的父节点为祖父节点的右子节点。
		4-3-1 插入节点为父节点的左子节点
		4-3-2 插入节点为父节点的右子节点
*/

func (rbtree *RBTree) insertNode(value int64) {
	rbtree.create = new(RBNode)
	rbtree.create.value = value
	rbtree.create.color = RED
	if rbtree.root == nil {
		rbtree.root = rbtree.create
		rbtree.root.color = BLACK
		rbtree.root.parent = nil
		return
	}

	rbtree.cur = rbtree.root
	for {
		if value < rbtree.cur.value {
			if rbtree.cur.left == nil {
				rbtree.cur.left = rbtree.create
				rbtree.create.parent = rbtree.cur
				break
			} else {
				rbtree.cur = rbtree.cur.left
			}
		} else if value > rbtree.cur.value {
			if rbtree.cur.right == nil {
				rbtree.cur.right = rbtree.create
				rbtree.create.parent = rbtree.cur
				break
			} else {
				rbtree.cur = rbtree.cur.right
			}
		} else {
			//value is exits
			return
		}

	}
	rbtree.InsertBalanceFixup(rbtree.create)

}

/**
修复红黑二叉树

修复次数不超过三次  ----有待验证
*/
func (rbtree *RBTree) InsertBalanceFixup(insertNode *RBNode) {
	var uncle *RBNode

	for insertNode.color == RED && insertNode.parent.color == RED {
		if insertNode.parent == insertNode.parent.parent.left {
			uncle = insertNode.parent.parent.right
		} else {
			uncle = insertNode.parent.parent.left
		}

		if uncle != nil && uncle.color == RED {
			uncle.color, insertNode.parent.color = BLACK, BLACK
			insertNode = insertNode.parent.parent
			if insertNode == rbtree.root || insertNode == nil {
				return
			}
			insertNode.color = RED
		} else {
			if insertNode.parent == insertNode.parent.parent.left {
				if insertNode == insertNode.parent.right {
					insertNode = insertNode.parent
					rbtree.LeftRotate(insertNode)
				}
				insertNode = insertNode.parent
				insertNode.color = BLACK
				insertNode = insertNode.parent
				insertNode.color = RED
				rbtree.RightRotate(insertNode)
			} else {
				if insertNode == insertNode.parent.left {
					insertNode = insertNode.parent
					rbtree.RightRotate(insertNode)
				}
				insertNode = insertNode.parent
				insertNode.color = BLACK
				insertNode = insertNode.parent
				insertNode.color = RED
				rbtree.LeftRotate(insertNode)
			}
			return
		}
	}
}

/**
 * 旋转图解(以向左旋转为例)：
 *     |                               |
 *     ○ <-left rotate                 ●
 *      \              ----------\    / \
 *       ●             ----------/   ○   ●r
 *      / \                           \
 *    l●   ●r                         l●
 *
 *
 *
 *     |                               |
 *     ○ <-left rotate                 ●
 *      \              ----------\    / \
 *       ●             ----------/   ○   ●
 *        \                           \
 *         ●                          nil <-don't forget it should be nil
 */
func (rbt *RBTree) LeftRotate(node *RBNode) {
	if node.right == nil {
		return
	}

	rightChild := node.right
	rbt.replaceNode(node, rightChild)
	node.right = rightChild.left
	if node.right != nil {
		node.right.parent = node
	}

	rightChild.left = node
	node.parent = rightChild
}

/**
---->主要处理父节点
替换新、旧节点的父节点
 */
func (rbt *RBTree) replaceNode(old *RBNode, new *RBNode) {
	if old.parent == nil {
		rbt.root = new
	} else {
		if old.parent.left == old {
			old.parent.left = new
		} else {
			old.parent.right = new
		}
	}

	if new != nil {
		new.parent = old.parent
	}
}
/**
围绕node节点右旋  和左旋相反
1、先子节点
2、再子节点父节点

3、旋转节点
4、旋转节点父节点
 */
func (rbt *RBTree) RightRotate(node *RBNode) {
	if node.left == nil {
		return
	}
	leftChild := node.left
	//父节点的替换
	rbt.replaceNode(node, leftChild)

	//子节点本身该放哪里？
	node.left = leftChild.right
	//子节点的父节点是谁？
	if node.left != nil {
		node.left.parent = node
	}
	//旋转接点该放哪里？
	leftChild.right = node
	//旋转接点的父节点是谁？
	node.parent = leftChild
}

func (rbt *RBTree) GetDeepth() int {
	var getDeepth func(node *RBNode) int

	getDeepth = func(node *RBNode) int {
		if node == nil {
			return 0
		}
		if node.left == nil && node.right == nil {
			return 1
		}
		var (
			ldeepth int = getDeepth(node.left)
			rdeepth int = getDeepth(node.right)
		)
		if ldeepth > rdeepth {
			return ldeepth + 1
		} else {
			return rdeepth + 1
		}
	}

	return getDeepth(rbt.root)
}
