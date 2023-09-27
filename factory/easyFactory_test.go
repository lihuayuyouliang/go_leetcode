package factory

import (
	"reflect"
	"testing"
)

// 测试
func TestNewIParser(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IParser
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: jsonParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
