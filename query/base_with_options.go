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

type baseWithOptions struct {
	*base
}

func newBaseWithOptions() *baseWithOptions {
	return &baseWithOptions{
		newBase(),
	}
}

// Select sets the fields to be returned as query results. Field aliases can be
// set by using "field as alias" syntax instead of just "field". In case of
// multiple Select calls the last one is what will be used.
func (q *baseWithOptions) Select(fields ...string) *baseWithOptions {
	q.values.Set("$select", strings.Join(fields, ","))
	return q
}

// Order sets the sorting order for a given field. In case of multiple Order
// calls for the same field the behavior is undefined.
func (q *baseWithOptions) Order(field string,
	order SortingOrder) *baseWithOptions {
	if order == SortingOrderAscending {
		q.values.Add("$order", field+" asc")
	} else {
		q.values.Add("$order", field+" desc")
	}
	return q
}

// Group groups results based on the given field. In case of multiple Group
// calls the last one is what will be used.
func (q *baseWithOptions) Group(field string) *baseWithOptions {
	q.values.Set("$group", field)
	return q
}

// Having allows filtering conditions on aggregate (Group) results. In case of
// multiple Having calls the last one is what will be used.
func (q *baseWithOptions) Having(having string) *baseWithOptions {
	q.values.Set("$having", having)
	return q
}

// Limit limits the number of results that will be returned by the query. In
// case of multiple Limit calls the last one is what will be used.
func (q *baseWithOptions) Limit(limit int) *baseWithOptions {
	q.values.Set("$limit", fmt.Sprintf("%d", limit))
	return q
}

// Offset skips the first offset results from the query. Used for pagination.
// In case of multiple Offset calls the last one is what will be used.
func (q *baseWithOptions) Offset(offset uint) *baseWithOptions {
	q.values.Set("$offset", fmt.Sprintf("%d", offset))
	return q
}
