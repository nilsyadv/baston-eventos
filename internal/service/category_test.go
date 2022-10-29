package service

import (
	"reflect"
	"testing"

	uuid "github.com/satori/go.uuid"

	custerror "github.com/nilsyadv/baston-eventos/error"
	"github.com/nilsyadv/baston-eventos/internal/model"
)

func TestGetCategory(t *testing.T) {
	type args struct {
		category *model.Category
		id       uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want *custerror.CustomeError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCategory(tt.args.category, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}
