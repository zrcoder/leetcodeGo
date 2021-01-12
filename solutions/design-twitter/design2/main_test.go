/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package design2

import "fmt"

func Example() {
	tt := Constructor()
	tt.PostTweet(1, 5)
	fmt.Println(tt.GetNewsFeed(1))
	tt.Follow(1, 2)
	tt.PostTweet(2, 6)
	fmt.Println(tt.GetNewsFeed(1))
	tt.Unfollow(1, 2)
	fmt.Println(tt.GetNewsFeed(1))
	// output:
	// [5]
	// [6 5]
	// [5]
}

func Example1() {
	tt := Constructor()
	tt.PostTweet(1, 5)
	tt.Follow(1, 1)
	fmt.Println(tt.GetNewsFeed(1))
	// output:
	// [5]
}

func Example2() {
	tt := Constructor()
	tt.PostTweet(1, 5)
	tt.Follow(1, 2)
	tt.Follow(2, 1)
	fmt.Println(tt.GetNewsFeed(2))
	tt.PostTweet(2, 6)
	fmt.Println(tt.GetNewsFeed(1))
	fmt.Println(tt.GetNewsFeed(2))
	tt.Unfollow(1, 2)
	fmt.Println(tt.GetNewsFeed(1))

	// output:
	// [5]
	// [6 5]
	// [6 5]
	// [5]
}
