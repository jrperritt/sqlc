// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package querytest

import (
	"context"
)

const listAuthors = `-- name: ListAuthors :one
SELECT  id, username, email, name, bio
FROM    authors
WHERE   email = CASE WHEN ?1 = '' then NULL else ?1 END
        OR username = CASE WHEN ?2 = '' then NULL else ?2 END
LIMIT   1
`

type ListAuthorsParams struct {
	Email    interface{}
	Username interface{}
}

func (q *Queries) ListAuthors(ctx context.Context, arg ListAuthorsParams) (Author, error) {
	row := q.db.QueryRowContext(ctx, listAuthors, arg.Email, arg.Username)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Name,
		&i.Bio,
	)
	return i, err
}