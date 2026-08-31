package repository

import "time"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	Createdat time.Time `json:"created_at"`
}

type InMemoryRepo struct {
	tasks  map[int]*Task
	nextID int
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}

func (r *InMemoryRepo) Create(t *Task) *Task {
	t.ID = r.nextID
	r.nextID++
	t.Createdat = time.Now()
	if t.Status == "" {
		t.Status = "pending"
	}
	r.tasks[t.ID] = t
	return t
}

func (r *InMemoryRepo) Get(id int) (*Task, bool) {
	task, ok := r.tasks[id]
	return task, ok
}

func (r *InMemoryRepo) List() []*Task {
	result := make([]*Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		result = append(result, t)
	}
	return result
}

func (r *InMemoryRepo) Update(id int, title *string, status *string) (*Task, bool) {
	task, ok := r.Get(id)
	if !ok {
		return nil, false
	}
	if title != nil {
		task.Title = *title
	}
	if status != nil {
		task.Status = *status
	}
	return task, true
}

func (r *InMemoryRepo) Delete(id int) bool {
	_, ok := r.tasks[id]
	if ok {
		delete(r.tasks, id)
	}
	return ok
}
