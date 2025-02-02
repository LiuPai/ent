// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeRelatives holds the string denoting the relatives edge name in mutations.
	EdgeRelatives = "relatives"
	// EdgeLikedTweets holds the string denoting the liked_tweets edge name in mutations.
	EdgeLikedTweets = "liked_tweets"
	// EdgeTweets holds the string denoting the tweets edge name in mutations.
	EdgeTweets = "tweets"
	// EdgeJoinedGroups holds the string denoting the joined_groups edge name in mutations.
	EdgeJoinedGroups = "joined_groups"
	// EdgeFriendships holds the string denoting the friendships edge name in mutations.
	EdgeFriendships = "friendships"
	// EdgeRelationship holds the string denoting the relationship edge name in mutations.
	EdgeRelationship = "relationship"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "likes"
	// EdgeUserTweets holds the string denoting the user_tweets edge name in mutations.
	EdgeUserTweets = "user_tweets"
	// Table holds the table name of the user in the database.
	Table = "users"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "user_groups"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "friendships"
	// RelativesTable is the table that holds the relatives relation/edge. The primary key declared below.
	RelativesTable = "relationships"
	// LikedTweetsTable is the table that holds the liked_tweets relation/edge. The primary key declared below.
	LikedTweetsTable = "tweet_likes"
	// LikedTweetsInverseTable is the table name for the Tweet entity.
	// It exists in this package in order to avoid circular dependency with the "tweet" package.
	LikedTweetsInverseTable = "tweets"
	// TweetsTable is the table that holds the tweets relation/edge. The primary key declared below.
	TweetsTable = "user_tweets"
	// TweetsInverseTable is the table name for the Tweet entity.
	// It exists in this package in order to avoid circular dependency with the "tweet" package.
	TweetsInverseTable = "tweets"
	// JoinedGroupsTable is the table that holds the joined_groups relation/edge.
	JoinedGroupsTable = "user_groups"
	// JoinedGroupsInverseTable is the table name for the UserGroup entity.
	// It exists in this package in order to avoid circular dependency with the "usergroup" package.
	JoinedGroupsInverseTable = "user_groups"
	// JoinedGroupsColumn is the table column denoting the joined_groups relation/edge.
	JoinedGroupsColumn = "user_id"
	// FriendshipsTable is the table that holds the friendships relation/edge.
	FriendshipsTable = "friendships"
	// FriendshipsInverseTable is the table name for the Friendship entity.
	// It exists in this package in order to avoid circular dependency with the "friendship" package.
	FriendshipsInverseTable = "friendships"
	// FriendshipsColumn is the table column denoting the friendships relation/edge.
	FriendshipsColumn = "user_id"
	// RelationshipTable is the table that holds the relationship relation/edge.
	RelationshipTable = "relationships"
	// RelationshipInverseTable is the table name for the Relationship entity.
	// It exists in this package in order to avoid circular dependency with the "relationship" package.
	RelationshipInverseTable = "relationships"
	// RelationshipColumn is the table column denoting the relationship relation/edge.
	RelationshipColumn = "user_id"
	// LikesTable is the table that holds the likes relation/edge.
	LikesTable = "tweet_likes"
	// LikesInverseTable is the table name for the TweetLike entity.
	// It exists in this package in order to avoid circular dependency with the "tweetlike" package.
	LikesInverseTable = "tweet_likes"
	// LikesColumn is the table column denoting the likes relation/edge.
	LikesColumn = "user_id"
	// UserTweetsTable is the table that holds the user_tweets relation/edge.
	UserTweetsTable = "user_tweets"
	// UserTweetsInverseTable is the table name for the UserTweet entity.
	// It exists in this package in order to avoid circular dependency with the "usertweet" package.
	UserTweetsInverseTable = "user_tweets"
	// UserTweetsColumn is the table column denoting the user_tweets relation/edge.
	UserTweetsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"user_id", "group_id"}
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"user_id", "friend_id"}
	// RelativesPrimaryKey and RelativesColumn2 are the table columns denoting the
	// primary key for the relatives relation (M2M).
	RelativesPrimaryKey = []string{"user_id", "relative_id"}
	// LikedTweetsPrimaryKey and LikedTweetsColumn2 are the table columns denoting the
	// primary key for the liked_tweets relation (M2M).
	LikedTweetsPrimaryKey = []string{"user_id", "tweet_id"}
	// TweetsPrimaryKey and TweetsColumn2 are the table columns denoting the
	// primary key for the tweets relation (M2M).
	TweetsPrimaryKey = []string{"user_id", "tweet_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
)
