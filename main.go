package main

type IndexStrategy struct {
	Name        string
	Columns     []string
	Type        string // "btree", "hash", etc.
	Cardinality float64
}

type QueryMetrics struct {
	ExecutionTime float64
	RowsScanned   int
	IndexUsed     string
	Cost          float64
}

type SimulationResult struct {
	Strategy IndexStrategy
	Metrics  []QueryMetrics
	Summary  PerformanceSummary
}

type PerformanceSummary struct {
	AverageExecutionTime float64
	TotalRowsScanned     int
	EfficiencyScore      float64
}

