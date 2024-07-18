package models

import "time"

var (
	comment1 = Comment{
		CommentID: 1,
		AritcleID: 1,
		Message:   "Hello",
		CreatedAt: time.Now(),
	}

	comment2 = Comment{
		CommentID: 2,
		AritcleID: 1,
		Message:   "world",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID:          1,
		Title:       "first article",
		Contents:    "This is first article",
		UserName:    "user1",
		NiceNum:     10,
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}

	Article2 = Article{
		ID:        2,
		Title:     "second article",
		Contents:  "This is second article",
		UserName:  "user2",
		NiceNum:   20,
		CreatedAt: time.Now(),
	}
)
