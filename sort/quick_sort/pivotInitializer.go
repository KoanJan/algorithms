package quick_sort

// PivotInitializer returns the pivot index and save the pivot value on the leftest size
type PivotInitializer func([]int, int, int) int
