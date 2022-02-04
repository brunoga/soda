package query

// FullTextBuilder is the full text query builder for the SOSQL API.
//
// See https://dev.socrata.com/docs/queries/q.html for more details.
type FullTextBuilder struct {
	*baseWithOptions
}

// NewFullTextBuilder returns a new full text query builder instance.
func NewFullTextBuilder() *FullTextBuilder {
	return &FullTextBuilder{
		newBaseWithOptions(),
	}
}

func (q *FullTextBuilder) Select(fields ...string) *FullTextBuilder {
	q.baseWithOptions.Select(fields...)
	return q
}

func (q *FullTextBuilder) From(from string) *FullTextBuilder {
	q.base.From(from)
	return q
}

// Q sets the full text query to be made. In case of multiple Q calls the last
// one is what will be used.
func (q *FullTextBuilder) Q(text string) *FullTextBuilder {
	q.values.Add("$q", text)
	return q
}

func (q *FullTextBuilder) Order(field string, order SortingOrder) *FullTextBuilder {
	q.baseWithOptions.Order(field, order)
	return q
}

func (q *FullTextBuilder) Group(field string) *FullTextBuilder {
	q.baseWithOptions.Group(field)
	return q
}

func (q *FullTextBuilder) Having(having string) *FullTextBuilder {
	q.baseWithOptions.Having(having)
	return q
}

func (q *FullTextBuilder) Limit(limit int) *FullTextBuilder {
	q.baseWithOptions.Limit(limit)
	return q
}

func (q *FullTextBuilder) Offset(offset uint) *FullTextBuilder {
	q.baseWithOptions.Offset(offset)
	return q
}
