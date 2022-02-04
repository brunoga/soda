package query

// Builder is the standard query builder for the SOSQL API.
//
// See https://dev.socrata.com/docs/queries/ for more details.
type Builder struct {
	*baseWithOptions
}

// NewBuilder returns a new query builder instance.
func NewBuilder() *Builder {
	return &Builder{
		newBaseWithOptions(),
	}
}

func (q *Builder) Select(fields ...string) *Builder {
	q.baseWithOptions.Select(fields...)
	return q
}

func (q *Builder) From(from string) *Builder {
	q.baseWithOptions.From(from)
	return q
}

// Where sets filtering conditions for the query results. In case of multiple
// Where calls the last one is what will be used.
func (q *Builder) Where(where string) *Builder {
	q.values.Set("$where", where)
	return q
}

func (q *Builder) Order(field string, order SortingOrder) *Builder {
	q.baseWithOptions.Order(field, order)
	return q
}

func (q *Builder) Group(field string) *Builder {
	q.baseWithOptions.Group(field)
	return q
}

func (q *Builder) Having(having string) *Builder {
	q.baseWithOptions.Having(having)
	return q
}

func (q *Builder) Limit(limit int) *Builder {
	q.baseWithOptions.Limit(limit)
	return q
}

func (q *Builder) Offset(offset uint) *Builder {
	q.baseWithOptions.Offset(offset)
	return q
}
