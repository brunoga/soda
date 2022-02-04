package query

// SQLBuilder is the SQL query builder for the SOSQL API.
//
// See https://dev.socrata.com/docs/queries/query.html for more details.
type SQLBuilder struct {
	*base
}

// NewSQLBuilder returns a new SQL query builder instance.
func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{
		newBase(),
	}
}

func (q *SQLBuilder) From(from string) *SQLBuilder {
	q.base.From(from)
	return q
}

// Query allows setting the SQL query to be made. In case of multiple Q calls
// the last one is what will be used.
func (q *SQLBuilder) Query(query string) *SQLBuilder {
	q.values.Set("$query", query)
	return q
}
