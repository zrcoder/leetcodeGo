/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package all_oone_data_structure

/*
实现一个数据结构支持以下操作：

Inc(key) - 插入一个新的值为 1 的 key。或者使一个存在的 key 增加一，保证 key 不为空字符串。
Dec(key) - 如果这个 key 的值是 1，那么把他从数据结构中移除掉。否者使一个存在的 key 值减一。如果这个 key 不存在，这个函数不做任何事情。key 保证不为空字符串。
GetMaxKey() - 返回 key 中值最大的任意一个。如果没有元素存在，返回一个空字符串""。
GetMinKey() - 返回 key 中值最小的任意一个。如果没有元素存在，返回一个空字符串""。
挑战：以 O(1) 的时间复杂度实现所有操作。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/all-oone-data-structure
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
能够在常数级时间添加或删除元素，哈希表、链表可当此任
要能在常数级时间获取到最大值或最小值，维护一个排序的数据结构即可，考虑到增加和删除需要常数级，这里选链表
注意到一个key只能从0连续递增上来，或从某个数连续递减下去，可以在常数时间找到列表中需要插入或删除的节点
*/
type ListNode struct {
	Val  int
	prev *ListNode
	next *ListNode
}

type List struct {
	head *ListNode
	tail *ListNode
}

func (l *List) PushFront(n *ListNode) {
	n.prev = nil
	if l.head != nil {
		l.head.prev = n
	} else {
		l.tail = n
	}
	n.next = l.head
	l.head = n
}

func (l *List) InsertBefore(n, before *ListNode) {
	if before != l.head {
		n.prev, n.next = before.prev, before
		n.prev.next, before.prev = n, n
	} else {
		n.prev, n.next = nil, before
		before.prev, l.head = n, n
	}
}

func (l *List) InsertAfter(n, after *ListNode) {
	if after != l.tail {
		n.prev, n.next = after, after.next
		n.next.prev, after.next = n, n
	} else {
		n.prev, n.next = after, nil
		after.next, l.tail = n, n
	}
}

func (l *List) PopFront() *ListNode {
	if l.head == nil {
		return nil
	}
	h := l.head
	if h.next != nil {
		l.head = h.next
	} else {
		l.head, l.tail = nil, nil
	}
	return h
}

func (l *List) PopBack() *ListNode {
	if l.tail == nil {
		return nil
	}
	r := l.tail
	if r.prev != nil {
		l.tail = r.prev
	} else {
		l.head, l.tail = nil, nil
	}
	return r
}

func (l *List) Remove(n *ListNode) {
	if l.head == n {
		l.PopFront()
		return
	} else if l.tail == n {
		l.PopBack()
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
}

//
//type AllOne struct {
//	keyCount  map[string]int
//	countKey  map[int]map[string]bool
//	countNode map[int]*ListNode
//	list      *List
//}
//
//func Constructor() AllOne {
//	return AllOne{
//		keyCount:  map[string]int{},
//		countKey:  map[int]map[string]bool{},
//		countNode: map[int]*ListNode{},
//		list:      &List{},
//	}
//}
//
//func (a *AllOne) Inc(key string) {
//	a.keyCount[key]++
//	v := a.keyCount[key]
//	a.addKV(key, v, true)
//	if v != 1 {
//		a.removeKV(key, v-1)
//	}
//}
//
//func (a *AllOne) Dec(key string) {
//	if _, ok := a.keyCount[key]; !ok {
//		return
//	}
//	a.keyCount[key]--
//	v := a.keyCount[key]
//	if v != 0 {
//		a.addKV(key, v, false)
//	}
//	a.removeKV(key, v+1)
//}
//
//func (a *AllOne) addKV(key string, val int, onInc bool) {
//	if m, ok := a.countKey[val]; ok {
//		m[key] = true
//	} else {
//		a.countKey[val] = map[string]bool{key: true}
//	}
//
//	if len(a.countKey[val]) == 1 {
//		n := &ListNode{Val: val}
//		a.countNode[val] = n
//		if onInc {
//			if val != 1 {
//				a.list.InsertAfter(n, a.countNode[val-1])
//			} else {
//				a.list.PushFront(n)
//			}
//		} else {
//			a.list.InsertBefore(n, a.countNode[val+1])
//		}
//	}
//}
//
//func (a *AllOne) removeKV(key string, val int) {
//	m := a.countKey[val]
//	if len(m) != 1 {
//		delete(m, key)
//	} else {
//		delete(a.countKey, val)
//		a.list.Remove(a.countNode[val])
//		delete(a.countNode, val)
//	}
//}
//
//func (a *AllOne) GetMaxKey() string {
//	return a.getOneKey(a.list.tail)
//}
//
//func (a *AllOne) GetMinKey() string {
//	return a.getOneKey(a.list.head)
//}
//
//func (a *AllOne) getOneKey(n *ListNode) string {
//	if nil == n {
//		return ""
//	}
//	for k, _ := range a.countKey[n.Val] {
//		return k
//	}
//	return ""
//}
//
//
//
//type AllOne struct {
//	count     map[string]int                   // 记录key的个数
//	countNode map[int]map[string]*list.Element // 记录个数为定值的链表节点，节点值为key
//	list      *list.List                       // 按照key的个数排序
//}
//
///** Initialize your data structure here. */
//func Constructor() AllOne {
//	return AllOne{
//		count:     make(map[string]int, 0),
//		countNode: make(map[int]map[string]*list.Element, 0),
//		list:      list.New(),
//	}
//}
//
///** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
//func (a *AllOne) Inc(key string) {
//	c, ok := a.count[key]
//	a.count[key]++
//	if ok {
//		e := a.countNode[c][key]
//		if len(a.countNode[c+1]) == 0 {
//			// a.list.MoveToBack(e)
//			a.countNode[c+1] = make(map[string]*list.Element, 1)
//		} else {
//			// choose any node in a.countNode[c+1] and move e after it
//			var node *list.Element
//			for _, n := range a.countNode[c+1] {
//				node = n
//				break
//			}
//			a.list.MoveAfter(e, node)
//		}
//		delete(a.countNode[c], key)
//		a.countNode[c+1][key] = e
//	} else {
//		e := a.list.PushFront(key)
//		if a.countNode[1] == nil {
//			a.countNode[1] = make(map[string]*list.Element, 1)
//		}
//		a.countNode[1][key] = e
//	}
//}
//
///** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
//func (a *AllOne) Dec(key string) {
//	c, ok := a.count[key]
//	if !ok {
//		return
//	}
//	if c == 1 {
//		delete(a.count, key)
//		_ = a.list.Remove(a.countNode[1][key])
//		delete(a.countNode[1], key)
//	} else {
//		a.count[key] = c - 1
//		e := a.countNode[c][key]
//		if len(a.countNode[c-1]) > 0 {
//			// choose any node in a.countNode[c-1] and insert e after it
//			var node *list.Element
//			for _, n := range a.countNode[c-1] {
//				node = n
//				break
//			}
//			a.list.MoveAfter(e, node)
//		}
//		delete(a.countNode[c], key)
//		a.countNode[c-1][key] = e
//	}
//}
//
///** Returns one of the keys with maximal value. */
//func (a *AllOne) GetMaxKey() string {
//	if a.list.Len() == 0 {
//		return ""
//	}
//	return a.list.Back().Value.(string)
//}
//
///** Returns one of the keys with Minimal value. */
//func (a *AllOne) GetMinKey() string {
//	if a.list.Len() == 0 {
//		return ""
//	}
//	fmt.Println(a.count)
//	fmt.Println(a.countNode)
//	return a.list.Front().Value.(string)
//}
