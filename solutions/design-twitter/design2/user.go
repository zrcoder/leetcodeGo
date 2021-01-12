/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package design2

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
	var tweetsList [][]*TweetInfo
	total := 0
	if len(u.tweets) > 0 {
		total += len(u.tweets)
		tweetsList = append(tweetsList, u.tweets)
	}

	for id := range u.followees {
		if users[id] == nil || len(users[id].tweets) == 0 {
			continue
		}
		tweetsList = append(tweetsList, users[id].tweets)
		total += len(users[id].tweets)
	}
	return getLatest(tweetsList, total)
}

func getLatest(tweetsList [][]*TweetInfo, total int) []int {
	const maxSize = 10
	n := maxSize
	if n > total {
		n = total
	}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		markedRow := -1
		maxTime := uint(0)
		for row, tweets := range tweetsList {
			if len(tweets) == 0 {
				continue
			}
			last := tweets[len(tweets)-1]
			if last.Time > maxTime {
				maxTime = last.Time
				markedRow = row
			}
		}
		if markedRow == -1 {
			continue
		}
		result[i] = tweetsList[markedRow][len(tweetsList[markedRow])-1].Id
		tweetsList[markedRow] = tweetsList[markedRow][:len(tweetsList[markedRow])-1]
	}
	return result
}
