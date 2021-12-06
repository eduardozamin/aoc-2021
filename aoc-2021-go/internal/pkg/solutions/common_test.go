package solutions

import (
	"reflect"
	"testing"
)

func Test_csvRow2intArray(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"should convert multi-number csv into int array",
			args{"1,2,3"},
			[]int{1, 2, 3},
			false},
		{"should convert single-number csv into int array",
			args{"99"},
			[]int{99},
			false},
		{"should thrown an error for an empty string",
			args{""},
			nil,
			true},
		{"should thrown an error if non-numbers are provided",
			args{"1,2,3,4,a,6"},
			nil,
			true},
		{"should thrown an error if empty entries are provided",
			args{"1,2,3,4,,,6"},
			nil,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := csvRow2intArray(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("csvRow2intArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("csvRow2intArray() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeEmpty(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"should remove empty strings from a slice",
			args{[]string{"2", "3", "", "  ", "5"}},
			[]string{"2", "3", "5"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeEmpty(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
