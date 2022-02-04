package query

// FullTextBuilder is the full text query builder for the SOSQL API.
//
// See https://dev.socrata.com/docs/queries/q.html for more details.
type FullTextBuilder[T any] struct {
	*baseWithOptions[T]
}

// NewFullTextBuilder returns a new full text query builder instance.
func NewFullTextBuilder[T any]() *FullTextBuilder[T] {
	return &FullTextBuilder[T]{
		newBaseWithOptions[T](),
	}
}

func (q *FullTextBuilder[T]) Select(fields ...string) *FullTextBuilder[T] {
	q.baseWithOptions.Select(fields...)
	return q
}

func (q *FullTextBuilder[T]) From(from string) *FullTextBuilder[T] {
	q.base.From(from)
	return q
}

// Q sets the full text query to be made. In case of multiple Q calls the last
// one is what will be used.
func (q *FullTextBuilder[T]) Q(text string) *FullTextBuilder[T] {
	q.values.Add("$q", text)
	return q
}

func (q *FullTextBuilder[T]) Order(field string, order SortingOrder) *FullTextBuilder[T] {
	q.baseWithOptions.Order(field, order)
	return q
}

func (q *FullTextBuilder[T]) Group(field string) *FullTextBuilder[T] {
	q.baseWithOptions.Group(field)
	return q
}

func (q *FullTextBuilder[T]) Having(having string) *FullTextBuilder[T] {
	q.baseWithOptions.Having(having)
	return q
}

func (q *FullTextBuilder[T]) Limit(limit int) *FullTextBuilder[T] {
	q.baseWithOptions.Limit(limit)
	return q
}

func (q *FullTextBuilder[T]) Offset(offset uint) *FullTextBuilder[T] {
	q.baseWithOptions.Offset(offset)
	return q
}

