/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package cache

import (
	"container/list"
)

/*
146. LRU缓存机制 https://leetcode-cn.com/problems/lru-cache

运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

进阶:

你是否可以在 O(1) 时间复杂度内完成这两种操作？

示例:

LRUCache lis = new LRUCache( 2 );

lis.put(1, 1);
lis.put(2, 2);
lis.get(1);       // 返回  1
lis.put(3, 3);    // 该操作会使得密钥 2 作废
lis.get(2);       // 返回 -1 (未找到)
lis.put(4, 4);    // 该操作会使得密钥 1 作废
lis.get(1);       // 返回 -1 (未找到)
lis.get(3);       // 返回  3
lis.get(4);       // 返回  4
*/
type Pair struct {
	Key, Val int
}

type LRUCache struct {
	lis      *list.List
	m        map[int]*list.Element
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		lis:      list.New(),
		m:        make(map[int]*list.Element, 0),
		Capacity: capacity,
	}
}

func (s *LRUCache) Get(key int) int {
	e, ok := s.m[key]
	if !ok {
		return -1
	}
	s.lis.MoveToFront(e)
	return e.Value.(*Pair).Val
}

func (s *LRUCache) Put(key int, value int) {
	pair := &Pair{Key: key, Val: value}
	if e, ok := s.m[key]; ok {
		_ = s.lis.Remove(e)
	} else if s.lis.Len() == s.Capacity {
		back := s.lis.Remove(s.lis.Back()).(*Pair)
		delete(s.m, back.Key)
	}
	s.m[key] = s.lis.PushFront(pair)
}
