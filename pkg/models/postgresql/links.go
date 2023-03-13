package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/parfenovvs/urlshortener/pkg/models"
)

type LinkModel struct {
	DB *sql.DB
}

func (m *LinkModel) Insert(shortUrl, originalUrl string) error {
	stmt := fmt.Sprintf(`INSERT INTO links
	(short_url, original_url, created)
	values ('%s', '%s', now()::timestamp)`, shortUrl, originalUrl)

	_, err := m.DB.Exec(stmt)
	if err != nil {
		fmt.Println("Cannot insert link")
		return err
	}

	return nil
}

func (m *LinkModel) Get(shortUrl string) (*models.Link, error) {
	stmt := fmt.Sprintf(`SELECT id, short_url, original_url, created FROM links
	WHERE short_url='%s' LIMIT 1`, shortUrl)

	rows, err := m.DB.Query(stmt)
	if err != nil {
		fmt.Println("Cannot get record")
		return nil, err
	}
	defer rows.Close()

	rows.Next()
	link := &models.Link{}
	rows.Scan(&link.ID, &link.ShortUrl, &link.OriginalUrl, &link.Created)

	return link, nil
}
