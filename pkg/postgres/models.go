package postgres

import (
	"time"
)

var schema = `
CREATE TABLE groups (
	group_id INTEGER 
             PRIMARY KEY,

	screen_name TEXT 
             UNIQUE 
             NOT NULL,

    name     TEXT
             UNIQUE
             NOT NULL,

	members_count INTEGER
             NOT_NULL
             CHECK (followers_count >= 0)
);

CREATE TABLE posts (
    post_id  INTEGER
             NOT NULL,

	group_id INTEGER
             REFERENCES groups (group_id)
             ON DELETE SET NULL,

    date     TIMESTAMP
             NOT NULL,

    title    TEXT
             NOT NULL,

    text     TEXT
			 NOT NULL,

    likes_count INTEGER
             NOT NULL
             CHECK (likes >= 0),

    views_count INTEGER
             NOT NULL
             CHECK (views >= 0),

    comments_count INTEGER
             NOT NULL
             CHECK (comments >= 0),

    reposts_count INTEGER
             NOT NULL
             CHECK (reposts >= 0),

	PRIMARY KEY (post_id, date)
);`

type Group struct {
	ID           int    `db:"group_id"`
	ScreenName   string `db:"screen_name"`
	Name         string `db:"name"`
	MembersCount int    `db:"members_count"`
}

type Post struct {
	ID            int       `db:"post_id"`
	GroupID       int       `db:"group_id"`
	Date          time.Time `db:"date"`
	Title         string    `db:"title"`
	Text          string    `db:"text"`
	LikesCount    int       `db:"likes_count"`
	ViewsCount    int       `db:"views_count"`
	CommentsCount int       `db:"comments_count"`
	RepostsCount  int       `db:"reposts_count"`
}
