package lib

// ReportType meanings
const (
	HR ReportType = iota
	IT
	PR
)

// consts
const amountReportTypes = 3

type (
	ReportType int

	Report struct {
		Id   int        `json:"id"`
		Type ReportType `json:"report_type"`
	}
	reportQueue[T, V comparable] struct {
		doubleSet *DoubleSet[int, ReportType]
		elems     []*Report
	}
)

func NewReportQueue() *reportQueue[int, ReportType] {
	return &reportQueue[int, ReportType]{
		doubleSet: NewDoubleSet[int, ReportType](),
		elems:     make([]*Report, 0),
	}
}

func (q *reportQueue[T, V]) Add(elem *Report) {
	q.doubleSet.Add(elem.Id, elem.Type)
	q.elems = append(q.elems, elem)
}

func (q *reportQueue[T, V]) Exists(elem Report) bool {
	return q.doubleSet.Exists(elem.Id, elem.Type)
}

func (q *reportQueue[T, V]) Peek() *Report {
	if len(q.elems) == 0 {
		return nil
	}
	return q.elems[0]
}

func (q *reportQueue[T, V]) Remove() *Report {
	if len(q.elems) == 0 {
		return nil
	}

	removingElem := q.elems[0]
	q.elems = q.elems[1:]
	q.doubleSet.Remove(removingElem.Id, removingElem.Type)

	return removingElem
}

func (q *reportQueue[T, V]) Size() int {
	return len(q.elems)
}
