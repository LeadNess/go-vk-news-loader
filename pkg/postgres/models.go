package postgres

import "time"

var schema = `
CREATE TABLE groups (
	group_id INTEGER 
             PRIMARY KEY,

	domain   TEXT 
             UNIQUE 
             NOT NULL,

    name     TEXT
             UNIQUE
             NOT NULL,

	followers_count INTEGER
             NOT_NULL
             CHECK (followers_count >= 0)
);

CREATE TABLE posts (
    post_id  INTEGER
             PRIMARY KEY,

	group_id INTEGER
             REFERENCES groups (group_id)
             ON DELETE SET NULL,

    date     TIMESTAMP
             NOT NULL,

    title    TEXT
             NOT NULL,

    text     TEXT,

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
             CHECK (reposts >= 0)
);`

type Group struct {
	ID             int    `db:"group_id"`
	Domain         string `db:"domain"`
	Name           string `db:"name"`
	FollowersCount int    `db:"followers_count"`
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
