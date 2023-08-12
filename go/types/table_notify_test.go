package types

import (
	"fmt"
	"testing"
	"time"

	"github.com/candiddev/shared/go/assert"
	"github.com/google/uuid"
)

func TestTableNotifyFromString(t *testing.T) {
	aa := uuid.New()
	ah := uuid.New()
	id := uuid.New()

	tests := map[string]struct {
		input string
		want  *TableNotify
	}{
		"bad": {
			input: "I'm not valid JSON",
			want:  &TableNotify{},
		},
		"good": {
			input: fmt.Sprintf(`{"authAccountID":"%s","authHouseholdID":"%s","id":"%s","table":"test"}`, aa, ah, id),
			want: &TableNotify{
				AuthAccountID:   &aa,
				AuthHouseholdID: &ah,
				ID:              id,
				Table:           "test",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := TableNotifyFromString(tc.input)
			assert.Equal(t, got, tc.want)
		})
	}
}

func TestTableNotifyString(t *testing.T) {
	now := time.Now()
	id := uuid.New()

	n := TableNotify{
		AuthAccountID:   &id,
		AuthHouseholdID: &id,
		ID:              uuid.New(),
		Operation:       TableNotifyOperationUpdate,
		Table:           "auth_account",
		Updated:         &now,
	}

	tests := map[string]struct {
		input TableNotify
		want  string
	}{
		"no updated": {
			input: TableNotify{
				ID:        n.ID,
				Operation: n.Operation,
				Table:     n.Table,
			},
			want: fmt.Sprintf(`{"id":%q,"operation":%d,"table":%q,"updated":null}`, n.ID, n.Operation, n.Table),
		},
		"string": {
			input: n,
			want:  fmt.Sprintf(`{"id":%q,"operation":%d,"table":%q,"updated":%q}`, n.ID, n.Operation, n.Table, n.Updated.Format(time.RFC3339Nano)),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.input.String(), tc.want)
		})
	}
}
