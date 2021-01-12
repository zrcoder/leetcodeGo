package pointer

/*
138. 复制带随机指针的链表

给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。

要求返回这个链表的 深拷贝。

我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

val：一个表示 Node.val 的整数。
random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。


示例 1：



输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
示例 2：



输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]
示例 3：



输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]
示例 4：

输入：head = []
输出：[]
解释：给定的链表为空（空指针），因此返回 null。


提示：

-10000 <= Node.val <= 10000
Node.random 为空（null）或指向链表中的节点。
节点数目不超过 1000 。
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 借助哈希表、切片的实现，哈希表记录节点的“索引”
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	s := make([]*Node, 0)
	m := make(map[*Node]int, 0)
	i := 0
	for p := head; p != nil; p = p.Next {
		n := &Node{Val: p.Val}
		s = append(s, n)
		m[p] = i
		i++
	}
	i = 0
	for p := head; p != nil; p = p.Next {
		if i < len(s)-1 {
			s[i].Next = s[i+1]
		}
		if p.Random != nil {
			s[i].Random = s[m[p.Random]]
		}
		i++
	}
	return s[0]
}

// 只用一个哈希表，哈希表存储原节点-克隆节点键值对
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	visited := map[*Node]*Node{}
	dummy := new(Node)
	q := dummy
	for p := head; p != nil; p = p.Next {
		n := getClone(p, visited)
		n.Next = getClone(p.Next, visited)
		n.Random = getClone(p.Random, visited)
		q.Next = n
		q = n
	}
	dummy, q = nil, dummy.Next
	return q
}

func getClone(n *Node, visited map[*Node]*Node) *Node {
	if n == nil {
		return nil
	}
	if visited[n] != nil {
		return visited[n]
	}
	r := &Node{Val: n.Val}
	visited[n] = r
	return r
}

// 省去哈希表的做法：先在原链表每个节点后一位增加一个克隆副本，再根据原链表的random指针确定克隆副本节点的random指针，最后分离原链表和克隆副本链表
func copyRandomList3(head *Node) *Node {
	if head == nil {
		return nil
	}
	for p := head; p != nil; {
		n := &Node{Val: p.Val}
		p.Next, n.Next = n, p.Next
		p = n.Next
	}
	for p := head; p != nil; p = p.Next.Next {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
	}
	result := head.Next
	for p, q := head, head.Next; q != nil; {
		p.Next = q.Next
		p = p.Next
		if p == nil {
			return result
		}
		q.Next = p.Next
		q = q.Next
	}
	return result
}
