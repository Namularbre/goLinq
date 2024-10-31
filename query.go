package linq

import "errors"

// Query is a struct the represent a query
type Query[TSource any, TResult any] struct {
	source   []TSource
	filters  []func(TSource) bool
	selector func(TSource) TResult
}

// NewQuery creates a query from the source slice
func NewQuery[TSource any, TResult any](source []TSource) *Query[TSource, TResult] {
	return &Query[TSource, TResult]{
		source: source,
	}
}

// Where adds a predicate in the query
func (q *Query[TSource, TResult]) Where(predicate func(TSource) bool) *Query[TSource, TResult] {
	q.filters = append(q.filters, predicate)
	return q
}

// Select set the selector of the query
func (q *Query[TSource, TResult]) Select(selector func(TSource) TResult) *Query[TSource, TResult] {
	q.selector = selector
	return q
}

// Skip skips n elements in the source
func (q *Query[TSource, TResult]) Skip(skip uint) *Query[TSource, TResult] {
	q.source = q.source[skip:]
	return q
}

// Take keep only n element from the source
func (q *Query[TSource, TResult]) Take(take uint) *Query[TSource, TResult] {
	q.source = q.source[0:take]
	return q
}

// applyFilters is a tool function the applies the filters of a query
func applyFilters[T any](source []T, filters []func(T) bool) []T {
	var filtered []T
	for _, item := range source {
		include := true
		for _, filter := range filters {
			if !filter(item) {
				include = false
				break
			}
		}
		if include {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// ToSlice run the query and return all the result in the form of a slice
func (q *Query[TSource, TResult]) ToSlice() []TResult {
	q.source = applyFilters(q.source, q.filters)
	return Select(q.source, q.selector)
}

// First run the query and return the first element, or an error if there is nothing
func (q *Query[TSource, TResult]) First() (*TResult, error) {
	q.source = applyFilters(q.source, q.filters)
	res := Select(q.source, q.selector)
	if len(res) > 0 {
		return &res[0], nil
	}
	return nil, errors.New("no result in the query output")
}

// FirstOrDefault run the query and return def if there is no result
func (q *Query[TSource, TResult]) FirstOrDefault(def *TResult) *TResult {
	q.source = applyFilters(q.source, q.filters)
	res := Select(q.source, q.selector)
	if len(res) > 0 {
		return &res[0]
	}
	return def
}

// FirstOrNil is like FirstOrDefault, but return nil if there is no result
func (q *Query[TSource, TResult]) FirstOrNil() *TResult {
	q.source = applyFilters(q.source, q.filters)
	res := Select(q.source, q.selector)
	if len(res) > 0 {
		return &res[0]
	}
	return nil
}
