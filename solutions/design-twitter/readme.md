## [设计tweet](https://leetcode-cn.com/problems/design-twitter)

```设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。
你的设计需要支持以下的几个功能：
postTweet(userId, tweetId): 创建一条新的推文
getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
follow(followerId, followeeId): 关注一个用户
unfollow(followerId, followeeId): 取消关注一个用户
```
```
示例:
Twitter twitter = new Twitter();
// 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
twitter.postTweet(1, 5);
// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
twitter.getNewsFeed(1);
// 用户1关注了用户2.
twitter.follow(1, 2);
// 用户2发送了一个新推文 (推文id = 6).
twitter.postTweet(2, 6);
// 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
// 推文id6应当在推文id5之前，因为它是在5之后发送的.
twitter.getNewsFeed(1);
// 用户1取消关注了用户2.
twitter.unfollow(1, 2);
// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
// 因为用户1已经不再关注用户2.
twitter.getNewsFeed(1);
```
面向对象的设计，主要逻辑在User类；

比较难实现的是getNewsFeed， 可用优先队列（小顶堆）来实现[design1](design1)；<br>
注意优化，从后向前遍历用户的tweet列表，如果优先队列里已经有10个元素， 当前tweet的时间不大于堆顶tweet的元素，可以不再遍历该用户的tweet列表<br>

也可以直接对所有tweet数组排序，注意到所有数组本来是已经排序的, 只是遍历下从后向前取前10个[design2](design2)；<br>

getNewsFeed复杂度分析：
假设n为用户及其关注用户所有tweet的个数， m为用户关注的用户加自己总共的用户数
利用小顶堆的实现： 时间复杂度O(nlogn), 最坏情况下所有tweet都要插入堆里一次； 空间复杂度O(1)， 堆里最多有10个元素
直接遍历数组实现： 时间复杂度O(10*m), 空间复杂度O(m)，有一个中间数组

另外，注意一些常见的异常处理，如一个user follow自己