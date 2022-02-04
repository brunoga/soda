package query

import (
	"fmt"
	"strings"
)

// SortingOrder represents the sorting order for a query field.
type SortingOrder uint8

const (
	SortingOrderAscending SortingOrder = iota
	SortingOrderDescending
)

type baseWithOptions[T any] struct {
	*base[T]
}

func newBaseWithOptions[T any]() *baseWithOptions[T] {
	return &baseWithOptions[T]{
		newBase[T](),
	}
}

// Select sets the fields to be returned as query results. Field aliases can be
// set by using "field as alias" syntax instead of just "field". In case of
// multiple Select calls the last one is what will be used.
func (q *baseWithOptions[T]) Select(fields ...string) *baseWithOptions[T] {
	q.values.Set("$select", strings.Join(fields, ","))
	return q
}

// Order sets the sorting order for a given field. In case of multiple Order
// calls for the same field the behavior is undefined.
func (q *baseWithOptions[T]) Order(field string,
	order SortingOrder) *baseWithOptions[T] {
	if order == SortingOrderAscending {
		q.values.Add("$order", field+" asc")
	} else {
		q.values.Add("$order", field+" desc")
	}
	return q
}

// Group groups results based on the given field. In case of multiple Group
// calls the last one is what will be used.
func (q *baseWithOptions[T]) Group(field string) *baseWithOptions[T] {
	q.values.Set("$group", field)
	return q
}

// Having allows filtering conditions on aggregate (Group) results. In case of
// multiple Having calls the last one is what will be used.
func (q *baseWithOptions[T]) Having(having string) *baseWithOptions[T] {
	q.values.Set("$having", having)
	return q
}

// Limit limits the number of results that will be returned by the query. In
// case of multiple Limit calls the last one is what will be used.
func (q *baseWithOptions[T]) Limit(limit int) *baseWithOptions[T] {
	q.values.Set("$limit", fmt.Sprintf("%d", limit))
	return q
}

// Offset skips the first offset results from the query. Used for pagination.
// In case of multiple Offset calls the last one is what will be used.
func (q *baseWithOptions[T]) Offset(offset uint) *baseWithOptions[T] {
	q.values.Set("$offset", fmt.Sprintf("%d", offset))
	return q
}

