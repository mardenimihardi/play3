package delivery

import (
	usecase "play3/internal/usecase"
	"testing"
)

func TestInitDelivery(t *testing.T) {
	type args struct {
		usecase usecase.Usecase
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestInitDelivery: success",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitDelivery(tt.args.usecase)
		})
	}
}
