package report

import "runtime"

type PlanningReport struct {
	TotalAllocated  float64
	TotalHeapInUse  float64
	TotalStackInUse float64
	Elapsed        float64
}

func NewPlanningReport(before runtime.MemStats, after runtime.MemStats) PlanningReport {
	return PlanningReport{
		TotalAllocated:  float64(after.TotalAlloc-before.TotalAlloc) / 1024,
		TotalHeapInUse:  float64(after.HeapAlloc-before.HeapAlloc) / 1024,
		TotalStackInUse: float64(after.StackInuse-before.StackInuse) / 1024,
	}
}
