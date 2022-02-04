package query

// Builder is the standard query builder for the SOSQL API.
//
// See https://dev.socrata.com/docs/queries/ for more details.
type Builder[T any] struct {
	*baseWithOptions[T]
}

// NewBuilder returns a new query builder instance.
func NewBuilder[T any]() *Builder[T] {
	return &Builder[T]{
		newBaseWithOptions[T](),
	}
}

func (q *Builder[T]) Select(fields ...string) *Builder[T] {
	q.baseWithOptions.Select(fields...)
	return q
}

func (q *Builder[T]) From(from string) *Builder[T] {
	q.baseWithOptions.From(from)
	return q
}

// Where sets filtering conditions for the query results. In case of multiple
// Where calls the last one is what will be used.
func (q *Builder[T]) Where(where string) *Builder[T] {
	q.values.Set("$where", where)
	return q
}

func (q *Builder[T]) Order(field string, order SortingOrder) *Builder[T] {
	q.baseWithOptions.Order(field, order)
	return q
}

func (q *Builder[T]) Group(field string) *Builder[T] {
	q.baseWithOptions.Group(field)
	return q
}

func (q *Builder[T]) Having(having string) *Builder[T] {
	q.baseWithOptions.Having(having)
	return q
}

func (q *Builder[T]) Limit(limit int) *Builder[T] {
	q.baseWithOptions.Limit(limit)
	return q
}

func (q *Builder[T]) Offset(offset uint) *Builder[T] {
	q.baseWithOptions.Offset(offset)
	return q
}

