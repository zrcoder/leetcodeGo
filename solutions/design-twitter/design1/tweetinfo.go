/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package design1

type TweetInfo struct {
	Id   int
	Time uint
}

func NewTweetInfo(id int, time uint) *TweetInfo {
	return &TweetInfo{Id: id, Time: time}
}
