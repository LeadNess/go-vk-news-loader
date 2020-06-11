package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type GroupsStorage interface {
	InsertGroup(group Group) error
	InsertGroups(groups []Group) error
}

type PostsStorage interface {
	InsertPost(post Post) error
	InsertPosts(post []Post) error
}

type NewsStorage interface {
	GroupsStorage
	PostsStorage
	CreateSchema() error
}

type Storage struct {
	 db *sqlx.DB
}

func OpenConnection(user, password, host, port, dbName string) (*Storage, error) {
	conStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName)
	db, err := sqlx.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return &Storage{db: db}, err
}

func (s *Storage) CreateSchema() error {
	res := s.db.MustExec(schema)
	_, err := res.RowsAffected()
	return err
}

func (s *Storage) InsertGroup(group Group) error {
	sql := `
		INSERT INTO 
			groups (group_id, domain, name, followers_count) 
		VALUES 
			(:group_id, :domain, :name, :followers_count)`
	res := s.db.MustExec(sql, group)
	_, err := res.RowsAffected()
	return err
}

func (s *Storage) InsertGroups(groups []Group) error {
	for _, group := range groups {
		if err := s.InsertGroup(group); err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) InsertPost(post Post) error {
	sql := `
		INSERT INTO 
			posts (post_id, group_id, date, title, text, likes_count, views_count, comments_count, reposts_count) 
		VALUES 
			(:post_id, :group_id, :date, :title, :text, :likes_count, :views_count, :comments_count, :reposts_count)`
	res := s.db.MustExec(sql, post)
	_, err := res.RowsAffected()
	return err
}

func (s *Storage) InsertPosts(posts []Post) error {
	for _, post := range posts {
		if err := s.InsertPost(post); err != nil {
			return err
		}
	}
	return nil
}
