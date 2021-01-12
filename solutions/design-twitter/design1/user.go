/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package design1

import "container/heap"

type User struct {
	tweets    []*TweetInfo
	followees map[int]struct{}
	Id        int
}

func NewUser(id int) *User {
	return &User{Id: id, followees: map[int]struct{}{}}
}

func (u *User) PostTweet(id int, time uint) {
	u.tweets = append(u.tweets, NewTweetInfo(id, time))
}

func (u *User) Follow(id int) {
	if u.Id == id {
		return
	}
	u.followees[id] = struct{}{}
}

func (u *User) Unfollow(id int) {
	delete(u.followees, id)
}

func (u *User) GetNewsFeed(users map[int]*User) []int {
	pq := NewPriorityQueue()
	pushQueue(pq, u.tweets)
	for id := range u.followees {
		if users[id] == nil || len(users[id].tweets) == 0 {
			continue
		}
		pushQueue(pq, users[id].tweets)
	}
	result := make([]int, pq.Len())
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = heap.Pop(pq).(*TweetInfo).Id
	}
	return result
}

func pushQueue(pq *Heap, infos []*TweetInfo) {
	const maxSize = 10
	for i := len(infos) - 1; i >= 0; i-- {
		info := infos[i]
		if pq.Len() < maxSize {
			heap.Push(pq, info)
		} else if top := heap.Pop(pq).(*TweetInfo); top.Time < info.Time {
			heap.Push(pq, info)
		} else {
			heap.Push(pq, top)
			break
		}
	}
}
