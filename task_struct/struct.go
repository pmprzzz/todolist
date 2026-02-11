package task_struct

type Task struct {
	Id          string
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

func (task *Task) GetId() string {
	return task.Id
}

func (task *Task) GetDescription() string {
	return task.Description
}

func (task *Task) GetStatus() string {
	return task.Status
}

func (task *Task) GetCreatedAt() string {
	return task.CreatedAt
}

func (task *Task) GetUpdatedAt() string {
	return task.UpdatedAt
}
