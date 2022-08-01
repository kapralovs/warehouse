package tasks

type TaskUseCase interface{
	Start()
	Complete()
}