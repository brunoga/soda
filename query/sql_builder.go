package query

// SQLBuilder is the SQL query builder for the SOSQL API.
//
// See https://dev.socrata.com/docs/queries/query.html for more details.
type SQLBuilder[T any] struct {
	*base[T]
}

// NewSQLBuilder returns a new SQL query builder instance.
func NewSQLBuilder[T any]() *SQLBuilder[T] {
	return &SQLBuilder[T]{
		newBase[T](),
	}
}

func (q *SQLBuilder[T]) From(from string) *SQLBuilder[T] {
	q.base.From(from)
	return q
}

// Query allows setting the SQL query to be made. In case of multiple Q calls
// the last one is what will be used.
func (q *SQLBuilder[T]) Query(query string) *SQLBuilder[T] {
	q.values.Set("$query", query)
	return q
}

