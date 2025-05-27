package report

import (
	"sync"
)

// ReportType meanings
const (
	HR ReportType = iota
	IT
	PR
)

// consts
const AmountReportTypes = 3

type (
	ReportType int

	Report struct {
		Id   int        `json:"id"`
		Type ReportType `json:"report_type"`
	}
	ReportQueue[T, V comparable] struct {
		mu    sync.Mutex
		elems []*Report
	}
)

func NewReportQueue() *ReportQueue[int, ReportType] {
	return &ReportQueue[int, ReportType]{
		elems: make([]*Report, 0),
	}
}

func (q *ReportQueue[T, V]) Add(elem *Report) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.elems = append(q.elems, elem)
}

func (q *ReportQueue[T, V]) Peek() *Report {
	if len(q.elems) == 0 {
		return nil
	}
	return q.elems[0]
}

func (q *ReportQueue[T, V]) Remove() *Report {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.elems) == 0 {
		return nil
	}

	removingElem := q.elems[0]
	q.elems = q.elems[1:]
	return removingElem
}

func (q *ReportQueue[T, V]) Size() int {
	return len(q.elems)
}
