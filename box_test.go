package fang

import "testing"

func TestBoxMode(t *testing.T) {
	tests := []struct {
		name string
		opts helpOptions
		want boxMode
	}{
		{"default", helpOptions{}, boxFilled},
		{"without fill", helpOptions{noFill: true}, boxOutlined},
		{"without box", helpOptions{noBox: true}, boxNone},
		{"both, box wins", helpOptions{noFill: true, noBox: true}, boxNone},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.opts.boxMode(); got != tt.want {
				t.Errorf("boxMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
