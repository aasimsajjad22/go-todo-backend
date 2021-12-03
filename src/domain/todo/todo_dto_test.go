package todo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTodoValidateError(t *testing.T) {
	todo := Todo{Description: ""}
	err := todo.Validate()

	assert.NotNil(t, err)
	assert.EqualValues(t, "please add description", err.Message)
	assert.EqualValues(t, 400, err.Status)
}

func TestTodoValidateSuccess(t *testing.T) {
	todo := Todo{Description: "  Test Description  "}
	err := todo.Validate()

	assert.Nil(t, err)
	assert.EqualValues(t, "Test Description", todo.Description)
}
