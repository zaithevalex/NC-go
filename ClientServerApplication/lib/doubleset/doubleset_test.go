package doubleset

import (
	"ClientServerApplication/lib/report"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name              string
		keys              [][2]any
		expectedDoubleSet *doubleSet[int, report.ReportType]
	}{
		{
			name: "Add: not repetitive, one report type",
			keys: [][2]any{
				{1, report.HR},
				{2, report.IT},
				{3, report.PR},
			},
			expectedDoubleSet: &doubleSet[int, report.ReportType]{
				mutex: sync.Mutex{},
				set: map[int]map[report.ReportType]struct{}{
					1: {report.HR: struct{}{}},
					2: {report.IT: struct{}{}},
					3: {report.PR: struct{}{}},
				},
			},
		},
		{
			name: "Add: not repetitive, all report types",
			keys: [][2]any{
				{1, report.HR},
				{1, report.IT},
				{1, report.PR},
				{2, report.IT},
				{3, report.PR},
			},
			expectedDoubleSet: &doubleSet[int, report.ReportType]{
				mutex: sync.Mutex{},
				set: map[int]map[report.ReportType]struct{}{
					1: {report.HR: struct{}{},
						report.IT: struct{}{},
						report.PR: struct{}{},
					},
					2: {report.IT: struct{}{}},
					3: {report.PR: struct{}{}},
				},
			},
		},
		{
			name: "Add: repetitive",
			keys: [][2]any{
				{1, report.HR},
				{1, report.HR},
				{2, report.IT},
				{3, report.PR},
			},
			expectedDoubleSet: &doubleSet[int, report.ReportType]{
				mutex: sync.Mutex{},
				set: map[int]map[report.ReportType]struct{}{
					1: {report.HR: struct{}{}},
					2: {report.IT: struct{}{}},
					3: {report.PR: struct{}{}},
				},
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			ds := NewDoubleSet[int, report.ReportType]()
			for _, set := range c.keys {
				ds.Add(set[0].(int), set[1].(report.ReportType))
			}

			assert.Equal(t, c.expectedDoubleSet, ds)
		})
	}
}

func TestExists(t *testing.T) {
	testCases := []struct {
		name           string
		keys           [2]any
		doubleSet      *doubleSet[int, report.ReportType]
		expectedResult bool
	}{
		{
			name: "Exists",
			keys: [2]any{1, report.HR},
			doubleSet: &doubleSet[int, report.ReportType]{
				mutex: sync.Mutex{},
				set: map[int]map[report.ReportType]struct{}{
					1: {report.HR: struct{}{}},
					2: {report.IT: struct{}{}},
					3: {report.PR: struct{}{}},
				},
			},
			expectedResult: true,
		},
		{
			name: "Not exists",
			keys: [2]any{1, report.IT},
			doubleSet: &doubleSet[int, report.ReportType]{
				mutex: sync.Mutex{},
				set: map[int]map[report.ReportType]struct{}{
					1: {report.HR: struct{}{}},
					2: {report.IT: struct{}{}},
					3: {report.PR: struct{}{}},
				},
			},
			expectedResult: false,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.doubleSet.Exists(c.keys[0].(int), c.keys[1].(report.ReportType)), c.expectedResult)
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name              string
		keys              [][2]any
		originalDoubleSet *doubleSet[int, report.ReportType]
		expectedDoubleSet *doubleSet[int, report.ReportType]
	}{
		{
			name:              "Remove: empty original dataset",
			keys:              [][2]any{{1, report.HR}},
			originalDoubleSet: &doubleSet[int, report.ReportType]{},
			expectedDoubleSet: &doubleSet[int, report.ReportType]{},
		},
		{
			name: "Remove: original dataset contains removing element",
			keys: [][2]any{{1, report.HR}},
			originalDoubleSet: &doubleSet[int, report.ReportType]{
				mutex: sync.Mutex{},
				set: map[int]map[report.ReportType]struct{}{
					1: {report.HR: struct{}{}},
					2: {report.IT: struct{}{}},
				},
			},
			expectedDoubleSet: &doubleSet[int, report.ReportType]{
				mutex: sync.Mutex{},
				set: map[int]map[report.ReportType]struct{}{
					2: {report.IT: struct{}{}},
				},
			},
		},
		{
			name: "Remove: original dataset contains simple element",
			keys: [][2]any{{1, report.HR}},
			originalDoubleSet: &doubleSet[int, report.ReportType]{
				mutex: sync.Mutex{},
				set: map[int]map[report.ReportType]struct{}{
					1: {report.HR: struct{}{}},
				},
			},
			expectedDoubleSet: &doubleSet[int, report.ReportType]{
				mutex: sync.Mutex{},
				set:   nil,
			},
		},
		{
			name: "Remove: original dataset does not contain removing element",
			keys: [][2]any{{1, report.PR}},
			originalDoubleSet: &doubleSet[int, report.ReportType]{
				mutex: sync.Mutex{},
				set: map[int]map[report.ReportType]struct{}{
					1: {report.HR: struct{}{}},
					2: {report.IT: struct{}{}},
				},
			},
			expectedDoubleSet: &doubleSet[int, report.ReportType]{
				mutex: sync.Mutex{},
				set: map[int]map[report.ReportType]struct{}{
					1: {report.HR: struct{}{}},
					2: {report.IT: struct{}{}},
				},
			},
		},
		{
			name: "Remove: original dataset does not contain removing element",
			keys: [][2]any{{3, report.HR}},
			originalDoubleSet: &doubleSet[int, report.ReportType]{
				mutex: sync.Mutex{},
				set: map[int]map[report.ReportType]struct{}{
					1: {report.HR: struct{}{}},
					2: {report.IT: struct{}{}},
				},
			},
			expectedDoubleSet: &doubleSet[int, report.ReportType]{
				mutex: sync.Mutex{},
				set: map[int]map[report.ReportType]struct{}{
					1: {report.HR: struct{}{}},
					2: {report.IT: struct{}{}},
				},
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			for _, k := range c.keys {
				c.originalDoubleSet.Remove(k[0].(int), k[1].(report.ReportType))
			}
			assert.Equal(t, c.expectedDoubleSet, c.originalDoubleSet)
		})
	}
}
