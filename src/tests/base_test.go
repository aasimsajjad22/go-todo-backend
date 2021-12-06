package tests

import (
	"fmt"
	"github.com/aasimsajjad22/go-todo-backend/app"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	go app.StartApplication()
	fmt.Println("about to start functional test cases .....")
	os.Exit(m.Run())
}
