package tlog

type Task string

func (t Task) Extend(task string) Task {
	return Task(string(t) + "-" + task)
}
