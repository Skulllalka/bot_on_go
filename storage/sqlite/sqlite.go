package sqlite

import (
	"context"
	"database/sql"

	e "github.com/Skulllalka/bot_on_go/lib"
	"github.com/Skulllalka/bot_on_go/storage"
	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, e.Wrap("can't connect to db", err)
	}

	if err := db.Ping(); err != nil {
		return nil, e.Wrap("can't ping db", err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) Save(ctx context.Context, p *storage.Page) error {
	query := `INSERT INTO pages (url, user_name) VALUES (?, ?)`
	if _, err := s.db.ExecContext(ctx, query, p.URL, p.UserName); err != nil {
		return e.Wrap("can't write a query", err)
	}
	return nil
}

func (s *Storage) PickRandom(ctx context.Context, username string) (*storage.Page, error) {
	query := `SELECT url FROM pages WHERE user_name=? ORDER BY RANDOM() LIMIT 1`
	var url string
	err := s.db.QueryRowContext(ctx, query, username).Scan(&url)
	if err == sql.ErrNoRows {
		return nil, storage.ErrNoSavedPages
	}
	if err != nil {
		return nil, e.Wrap("can't pick url", err)
	}
	return &storage.Page{
		URL:      url,
		UserName: username,
	}, nil
}

func (s *Storage) Remove(ctx context.Context, p *storage.Page) error {
	query := `DELETE FROM pages WHERE url = ? AND user_name = ?`

	if _, err := s.db.ExecContext(ctx, query, p.URL, p.UserName); err != nil {
		return e.Wrap("can't remove page", err)
	}
	return nil
}

func (s *Storage) IsExists(ctx context.Context, p *storage.Page) (bool, error) {
	query := `SELECT COUNT(*) FROM pages WHERE url = ? AND user_name = ?`

	var count int
	err := s.db.QueryRowContext(ctx, query, p.URL, p.UserName).Scan(&count)
	if err != nil {
		return false, e.Wrap("can't get query", err)
	}
	return count > 0, nil
}

func (s *Storage) Init(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS pages (url TEXT, user_name TEXT)`
	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return e.Wrap("can't crate main table", err)
	}
	return nil
}
