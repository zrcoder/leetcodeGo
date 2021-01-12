/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package design_twitter

// tweetinfo.go
type TweetInfo struct {
	Id   int
	Time uint
}

func NewTweetInfo(id int, time uint) *TweetInfo {
	return &TweetInfo{Id: id, Time: time}
}

// user.go
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
		markedRow := 0
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
		result[i] = tweetsList[markedRow][len(tweetsList[markedRow])-1].Id
		tweetsList[markedRow] = tweetsList[markedRow][:len(tweetsList[markedRow])-1]
	}
	return result
}

// main.go
type Twitter struct {
	Users map[int]*User
	time  uint
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{Users: make(map[int]*User, 0)}
}

/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.time++
	t.getOrAddUser(userId).PostTweet(tweetId, t.time)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed.
Each item in the news feed must be posted by users who the user followed or by the user herself.
Tweets must be ordered from most recent to least recent.
*/
func (t *Twitter) GetNewsFeed(userId int) []int {
	return t.getOrAddUser(userId).GetNewsFeed(t.Users)
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	t.getOrAddUser(followerId).Follow(followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	t.getOrAddUser(followerId).Unfollow(followeeId)
}

/** search and return a user, if not present, generate and add one **/
func (t *Twitter) getOrAddUser(userId int) *User {
	user, ok := t.Users[userId]
	if !ok {
		user = NewUser(userId)
		t.Users[userId] = user
	}
	return user
}
