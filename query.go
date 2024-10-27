package linq

import "errors"

type IQueryable[TSource any, TResult any] interface {
	ToSlice() TSource[]
	First() (*TSource, error)
	FirstOrNil() *TSource
	FirstOrDefault(d *TSource) *TSource
	Select(func(source TSource) TResult) []TResult
	Where(func(source TSource) bool) []TSource
	Skip(x uint) []TSource
	Limit(x uint) []TSource
}

type Query[T any] struct {
	content []T
}

func (q *Query[T]) ToSlice() []T {
	return q.content
}

func (q *Query[T]) First() (*T, error) {
	if len(q.content) != 0 {
		return &q.content[0], nil
	}
	return nil, errors.New("no matching elements")
}

func (q *Query[T]) FirstOrDefault(d *T) *T {
	if len(q.content) != 0 {
		return &q.content[0]
	}
	return d
}

// FirstOrNil is like FirstOrDefault, but shorter if you want to return nil
func (q *Query[T]) FirstOrNil() *T {
	if len(q.content) != 0 {
		return &q.content[0]
	}
	return nil
}
