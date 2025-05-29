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
	reportQueue[T, V comparable] struct {
		elems []*Report
		mutex sync.Mutex
	}
)

func NewReportQueue() *reportQueue[int, ReportType] {
	return &reportQueue[int, ReportType]{
		elems: make([]*Report, 0),
		mutex: sync.Mutex{},
	}
}

func (q *reportQueue[T, V]) Add(elem *Report) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.elems = append(q.elems, elem)
}

func (q *reportQueue[T, V]) Remove() *Report {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.elems) == 0 {
		return nil
	}

	removingElem := q.elems[0]
	q.elems = q.elems[1:]
	return removingElem
}

func (q *reportQueue[T, V]) Size() int {
	return len(q.elems)
}
