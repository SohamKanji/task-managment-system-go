package db

import (
	"context"
	"math/rand"
	"testing"
	"time"

	db "github.com/SohamKanji/task-management-system-go/db/sqlc"
	"github.com/SohamKanji/task-management-system-go/utils"
	"github.com/stretchr/testify/require"
)

func verifyAccount(t *testing.T, due time.Time, status string, priority int64, title string, description string, task db.Task) {
	require.NotEmpty(t, task)
	require.Equal(t, due.Year(), task.Due.Year())
	require.Equal(t, due.Month(), task.Due.Month())
	require.Equal(t, due.Day(), task.Due.Day())
	require.Equal(t, status, task.Status)
	require.Equal(t, priority, task.Priority)
	require.Equal(t, title, task.Title)
	require.Equal(t, description, task.Description)
}

func TestCreateTask(t *testing.T) {
	arg := db.CreateTaskParams{
		Due:         time.Now(),
		Status:      utils.GetRandomStatus(),
		Priority:    rand.Int63n(3),
		Title:       utils.GetRandomString(6),
		Description: utils.GetRandomString(10),
	}
	task, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	verifyAccount(t, arg.Due, arg.Status, arg.Priority, arg.Title, arg.Description, task)
}

func TestGetTask(t *testing.T) {
	arg := db.CreateTaskParams{
		Due:         time.Now(),
		Status:      utils.GetRandomStatus(),
		Priority:    rand.Int63n(3),
		Title:       utils.GetRandomString(6),
		Description: utils.GetRandomString(10),
	}
	task1, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)

	task2, err := testQueries.GetTask(context.Background(), task1.ID)
	require.NoError(t, err)
	verifyAccount(t, task1.Due, task1.Status, task1.Priority, task1.Title, task1.Description, task2)
}

func TestListTasks(t *testing.T) {
	ids := make([]int64, 0)
	for i := 0; i < 10; i++ {
		arg := db.CreateTaskParams{
			Due:         time.Now(),
			Status:      utils.GetRandomStatus(),
			Priority:    rand.Int63n(3),
			Title:       utils.GetRandomString(6),
			Description: utils.GetRandomString(10),
		}
		task, err := testQueries.CreateTask(context.Background(), arg)
		require.NoError(t, err)
		ids = append(ids, task.ID)
	}

	tasks, err := testQueries.ListTasks(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(tasks), 10)
	taskMap := make(map[int64]struct{})
	for _, task := range tasks {
		require.NotEmpty(t, task)
		taskMap[task.ID] = struct{}{}
	}
	for _, id := range ids {
		_, ok := taskMap[id]
		require.True(t, ok)
	}
}

func TestUpdateTask(t *testing.T) {
	arg := db.CreateTaskParams{
		Due:         time.Now(),
		Status:      utils.GetRandomStatus(),
		Priority:    rand.Int63n(3),
		Title:       utils.GetRandomString(6),
		Description: utils.GetRandomString(10),
	}

	task1, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)

	update_arg := db.UpdateTaskParams{
		ID:          task1.ID,
		Due:         time.Now(),
		Status:      utils.GetRandomStatus(),
		Priority:    rand.Int63n(3),
		Title:       utils.GetRandomString(6),
		Description: utils.GetRandomString(10),
	}

	task2, err := testQueries.UpdateTask(context.Background(), update_arg)
	require.NoError(t, err)
	verifyAccount(t, update_arg.Due, update_arg.Status, update_arg.Priority, update_arg.Title, update_arg.Description, task2)
}

func TestDeleteTask(t *testing.T) {
	arg := db.CreateTaskParams{
		Due:         time.Now(),
		Status:      utils.GetRandomStatus(),
		Priority:    rand.Int63n(3),
		Title:       utils.GetRandomString(6),
		Description: utils.GetRandomString(10),
	}
	task1, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)

	err = testQueries.DeleteTask(context.Background(), task1.ID)
	require.NoError(t, err)

	task2, err := testQueries.GetTask(context.Background(), task1.ID)
	require.Error(t, err)
	require.Empty(t, task2)
}
