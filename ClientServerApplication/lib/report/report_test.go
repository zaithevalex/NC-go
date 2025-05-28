package report

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name                string
		elems               []*Report
		expectedReportQueue reportQueue[int, ReportType]
	}{
		{
			name:  "Add: 0 reports",
			elems: []*Report{},
			expectedReportQueue: reportQueue[int, ReportType]{
				elems: make([]*Report, 0),
				mutex: sync.Mutex{},
			},
		},
		{
			name: "Add: > 0 reports",
			elems: []*Report{
				{Id: 1, Type: HR},
				{Id: 2, Type: IT},
			},
			expectedReportQueue: reportQueue[int, ReportType]{
				elems: []*Report{
					{Id: 1, Type: HR},
					{Id: 2, Type: IT},
				},
				mutex: sync.Mutex{},
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			rq := NewReportQueue()
			for _, elem := range c.elems {
				rq.Add(elem)
			}
			assert.Equal(t, c.expectedReportQueue, *rq)
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name                string
		originalReportQueue reportQueue[int, ReportType]
		expectedReportQueue reportQueue[int, ReportType]
		expectedReport      *Report
	}{
		{
			name: "Remove: 0 reports",
			originalReportQueue: reportQueue[int, ReportType]{
				mutex: sync.Mutex{},
			},
			expectedReportQueue: reportQueue[int, ReportType]{
				mutex: sync.Mutex{},
			},
			expectedReport: nil,
		},
		{
			name: "Remove: 1 report",
			originalReportQueue: reportQueue[int, ReportType]{
				elems: []*Report{
					{Id: 1, Type: HR},
				},
				mutex: sync.Mutex{},
			},
			expectedReportQueue: reportQueue[int, ReportType]{
				elems: make([]*Report, 0),
				mutex: sync.Mutex{},
			},
			expectedReport: &Report{Id: 1, Type: HR},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			report := c.originalReportQueue.Remove()
			assert.Equal(t, c.expectedReportQueue, c.originalReportQueue)
			assert.Equal(t, c.expectedReport, report)
		})
	}
}

func TestSize(t *testing.T) {
	testCases := []struct {
		name         string
		reportQueue  *reportQueue[int, ReportType]
		expectedSize int
	}{
		{
			name: "Size: 0",
			reportQueue: &reportQueue[int, ReportType]{
				mutex: sync.Mutex{},
				elems: make([]*Report, 0),
			},
			expectedSize: 0,
		},
		{
			name: "Size: > 0",
			reportQueue: &reportQueue[int, ReportType]{
				mutex: sync.Mutex{},
				elems: []*Report{{}, {}, {}},
			},
			expectedSize: 3,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expectedSize, c.reportQueue.Size())
		})
	}
}
