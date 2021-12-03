package todo

import (
	"github.com/aasimsajjad22/go-todo-backend/utils/errors"
	"strings"
)

type Todo struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
}

func (t *Todo) Validate() *errors.RestErr {
	t.Description = strings.TrimSpace(t.Description)
	if t.Description == "" {
		return errors.NewBadRequestError("please add description")
	}
	return nil
}
