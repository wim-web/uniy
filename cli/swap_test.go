package main

import (
	"testing"
)

// TestCalculateAverageCost は CalculateAverageCost 関数のテストを行います
func TestCalculateAverageCost(t *testing.T) {
	tests := []struct {
		name           string
		orders         []OrderData
		expectedResult float64
	}{
		{
			name: "正常なケース1",
			orders: []OrderData{
				{"WDOGE(WRAPPED-DOGE)", "UNIX", 10, 100, 0},
			},
			expectedResult: 0.1,
		},
		{
			name: "正常なケース2",
			orders: []OrderData{
				{"WDOGE(WRAPPED-DOGE)", "UNIX", 10, 100, 0},
				{"WDOGE(WRAPPED-DOGE)", "UNIX", 10, 200, 0},
				{"WDOGE(WRAPPED-DOGE)", "UNIX", 10, 300, 0},
			},
			expectedResult: 0.05,
		},
		{
			name: "正常なケース3",
			orders: []OrderData{
				{"WDOGE(WRAPPED-DOGE)", "UNIX", 10, 100, 0},
				{"UNIX", "WDOGE(WRAPPED-DOGE)", 50, 5, 0},
				{"WDOGE(WRAPPED-DOGE)", "UNIX", 10, 100, 0},
			},
			expectedResult: 0.1,
		},
		{
			name:           "空のケース",
			orders:         []OrderData{},
			expectedResult: 0,
		},
		{
			name: "Amt1が0のケース",
			orders: []OrderData{
				{"WDOGE(WRAPPED-DOGE)", "UNIX", 10000000000, 0, 0},
			},
			expectedResult: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CalculateAverageCost(test.orders, "UNIX")
			if result != test.expectedResult {
				t.Errorf("Test %s failed: expected %f, got %f", test.name, test.expectedResult, result)
			}
		})
	}
}
