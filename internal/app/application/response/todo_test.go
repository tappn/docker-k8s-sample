package response

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/tappn/docker-k8s-sample/internal/app/domain/model"
)

func Test_NewTodo(t *testing.T) {
	t.Parallel()
	now := time.Now()
	todo1 := &model.Todo{
		ID:          1,
		Title:       "title",
		Description: "description",
		Deadline:    now,
	}

	expected1 := &Todo{
		ID:          1,
		Title:       "title",
		Description: "description",
		Deadline:    now,
	}

	tests := map[string]struct {
		todo     *model.Todo
		expected *Todo
	}{
		"new": {
			todo:     todo1,
			expected: expected1,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			resp := NewTodo(test.todo)
			if diff := cmp.Diff(resp, test.expected); diff != "" {
				t.Errorf("(-expected +actual)\n%s", diff)
			}
		})
	}
}
