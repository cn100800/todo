package db

import "testing"

func TestInsert(t *testing.T) {
	type args struct {
		table string
		m     map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Insert",
			args: args{
				table: "todo",
				m: map[string]string{
					"hello_key": "hello",
					"world_key": "world",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Insert(tt.args.table, tt.args.m)
		})
	}
}
