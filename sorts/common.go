package sorts

// StepCallback вызывается при каждом заметном изменении массива во время сортировки.
type StepCallback func([]int)

// emitStep копирует текущее состояние и передает его колбэку.
func emitStep(cb StepCallback, arr []int) {
	if cb == nil {
		return
	}
	snapshot := append([]int(nil), arr...)
	cb(snapshot)
}
