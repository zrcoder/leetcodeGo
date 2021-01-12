/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package design2

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
