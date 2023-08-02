//go:generate mockgen -source=main.go -destination=./mock_main.go -package=main
package main

import (
	"testing"

	"go.uber.org/mock/gomock"
)

func TestNumber_Plus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	type fields struct {
		num int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "test 1",
			fields: fields{num:10},
			args:   args{num:10},
			want:   20,
		},
		{
			name:   "test 2",
			fields: fields{num:-10},
			args:   args{num:-10},
			want:   -20,
		},
	}
	for _, tt := range tests {
		app := NewMockNumIf(ctrl)
		app.EXPECT().GetNum().Return(tt.fields.num)
		t.Run(tt.name, func(t *testing.T) {
			calc := &ExeCalc{n:app}
			if got := calc.Plus(tt.args.num); got != tt.want {
				t.Errorf("Plus() = %d, want %d", got, tt.want)
			}
		})

	}
}

