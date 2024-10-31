package linq

import (
	"errors"
)

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
		selector: func(source TSource) TResult {
			return *new(TResult)
		},
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

// toTResultSlice converts a TSource into a TResult
func toTResultSlice[TSource any, TResult any](source []TSource) []TResult {
	result := make([]TResult, len(source))
	for i, item := range source {
		result[i] = any(item).(TResult)
	}
	return result
}

// ToSlice run the query and return all the result in the form of a slice
func (q *Query[TSource, TResult]) ToSlice() []TResult {
	if q.filters != nil {
		q.source = applyFilters(q.source, q.filters)
	}

	// Si selector est nil, on retourne source directement (en supposant TSource == TResult)
	if q.selector == nil {
		return toTResultSlice[TSource, TResult](q.source)
	}

	// Si selector est défini, on l'applique
	return Select(q.source, q.selector)
}

// First run the query and return the first element, or an error if there is nothing
func (q *Query[TSource, TResult]) First() (*TResult, error) {
	if q.filters != nil {
		q.source = applyFilters(q.source, q.filters)
	}
	res := q.ToSlice()
	if len(res) > 0 {
		return &res[0], nil
	}
	return nil, errors.New("no result in the query output")
}

// FirstOrDefault run the query and return def if there is no result
func (q *Query[TSource, TResult]) FirstOrDefault(def *TResult) *TResult {
	if q.filters != nil {
		q.source = applyFilters(q.source, q.filters)
	}
	res := q.ToSlice()
	if len(res) > 0 {
		return &res[0]
	}
	return def
}

// FirstOrNil is like FirstOrDefault, but returns nil if there is no result
func (q *Query[TSource, TResult]) FirstOrNil() *TResult {
	if q.filters != nil {
		q.source = applyFilters(q.source, q.filters)
	}
	res := q.ToSlice()
	if len(res) > 0 {
		return &res[0]
	}
	return nil
}
