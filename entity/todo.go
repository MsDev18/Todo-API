package entity

type Todo struct {
	ID     uint
	Task   string
	// we have 3 status start / done / progress
	// create new type alies for status 
	Status Status
	// Refrence to cateogry entity
	CategoryID uint
}

type Status string

const (
	TodoStart    Status = "start"
	TodoProgress Status = "progress"
	TodoDone     Status = "done"
)
