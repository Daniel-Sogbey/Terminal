package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
)

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Published int       `json:"published"`
	Pages     int       `json:"pages"`
	Genres    []string  `json:"genres"`
	Rating    float32   `json:"rating"`
	Version   int32     `json:"version"`
}

type BookModel struct {
	DB *sql.DB
}

// insert a book into the db
func (b BookModel) Insert(book *Book) error {
	stmt := `insert into books (title,published, pages, genres, rating) values ($1,$2,$3,$4,$5)
	returning id, created_at, version`

	args := []interface{}{book.Title, book.Published, book.Pages, pq.Array(book.Genres), book.Rating}

	return b.DB.QueryRow(stmt, args...).Scan(&book.ID, &book.CreatedAt, &book.Version)
}

// get a single book by id
func (b BookModel) Get(id int64) (*Book, error) {
	if id < 1 {
		return nil, errors.New("record not found")
	}

	query := `select (id,title, published, pages,genres, rating, version,created_at) from books where id = $1`

	var book Book
	err := b.DB.QueryRow(query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Published,
		&book.Pages,
		pq.Array(&book.Genres),
		&book.Rating,
		&book.Version,
		&book.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("recod not found")
		default:
			return nil, err
		}

	}

	return &book, nil
}

func (b *BookModel) Update(book *Book) error {
	stmt := `update books set title = $1, published = $2, pages = $3, genres = $4, rating =$5, version = version+1
 where id = $6 and version = $7 returning version`

	args := []interface{}{book.Title, book.Published, book.Pages, pq.Array(book.Genres), book.Rating, book.ID, book.Version}

	return b.DB.QueryRow(stmt, args...).Scan(&book.Version)
}

func (b *BookModel) Delete(id int) error {
	stmt := `delete from books where id = $1`

	results, err := b.DB.Exec(stmt, id)

	if err != nil {
		return err
	}

	rowAffected, err := results.RowsAffected()

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errors.New("record not found")
	}

	return nil

}

func (b *BookModel) GetAll() ([]*Book, error) {
	query := `select * from books order by id`

	rows, err := b.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*Book

	for rows.Next() {
		var book Book

		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Published,
			&book.Pages,
			pq.Array(&book.Genres),
			&book.Rating,
			&book.Version,
			&book.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil

}
