package main

import (
	"reflect"
	"testing"
)

func TestHandleWithoutZ(t *testing.T) {
	t1, _ := GenerateCounters("XXYYXY")
	t2, _ := GenerateCounters("XYYX")
	t3, _ := GenerateCounters("XYYY")
	t4, _ := GenerateCounters("XYYXY")
	type args struct {
		queue    string
		counters Counters
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "XXYYXY - CORRECT",
			args: args{
				queue:    "XXYYXY",
				counters: t1,
			},
			want: "Yes",
		},
		{
			name: "XYYX - INCORRECT",
			args: args{
				queue:    "XYYX",
				counters: t2,
			},
			want: "No",
		},
		{
			name: "XYYY - INCORRECT",
			args: args{
				queue:    "XYYY",
				counters: t3,
			},
			want: "No",
		},
		{
			name: "XYYXY - INCORRECT",
			args: args{
				queue:    "XYYXY",
				counters: t4,
			},
			want: "No",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleWithoutZ(tt.args.queue, tt.args.counters); !reflect.DeepEqual(got.isCorrect(), tt.want) {
				t.Errorf("HandleWithoutZ() = %v, want %v", got.isCorrect(), tt.want)
			}
		})
	}
}

func TestHandleWithoutY(t *testing.T) {
	t1, _ := GenerateCounters("XXZZXZ")
	t2, _ := GenerateCounters("XZZX")
	t3, _ := GenerateCounters("XZZZ")
	t4, _ := GenerateCounters("XZZXZ")
	type args struct {
		queue    string
		counters Counters
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "XXZZXZ - CORRECT",
			args: args{
				queue:    "XXZZXZ",
				counters: t1,
			},
			want: "Yes",
		},
		{
			name: "XZZX - INCORRECT",
			args: args{
				queue:    "XZZX",
				counters: t2,
			},
			want: "No",
		},
		{
			name: "XZZZ - INCORRECT",
			args: args{
				queue:    "XZZZ",
				counters: t3,
			},
			want: "No",
		},
		{
			name: "XZZXZ - INCORRECT",
			args: args{
				queue:    "XZZXZ",
				counters: t4,
			},
			want: "No",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleWithoutY(tt.args.queue, tt.args.counters); !reflect.DeepEqual(got.isCorrect(), tt.want) {
				t.Errorf("HandleWithoutY() = %v, want %v", got.isCorrect(), tt.want)
			}
		})
	}
}

func TestHandleWithoutX(t *testing.T) {
	t1, _ := GenerateCounters("YYZZYZ")
	t2, _ := GenerateCounters("YZZY")
	t3, _ := GenerateCounters("YZZZ")
	t4, _ := GenerateCounters("YZZYZ")
	type args struct {
		queue    string
		counters Counters
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "YYZZYZ - CORRECT",
			args: args{
				queue:    "YYZZYZ",
				counters: t1,
			},
			want: "Yes",
		},
		{
			name: "YZZY - INCORRECT",
			args: args{
				queue:    "YZZY",
				counters: t2,
			},
			want: "No",
		},
		{
			name: "YZZZ - INCORRECT",
			args: args{
				queue:    "YZZZ",
				counters: t3,
			},
			want: "No",
		},
		{
			name: "YZZYZ - INCORRECT",
			args: args{
				queue:    "YZZYZ",
				counters: t4,
			},
			want: "No",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleWithoutX(tt.args.queue, tt.args.counters); !reflect.DeepEqual(got.isCorrect(), tt.want) {
				t.Errorf("HandleWithoutY() = %v, want %v", got.isCorrect(), tt.want)
			}
		})
	}
}

func TestHandleGeneral(t *testing.T) {
	// t1, _ := GenerateCounters("XXYYXZ")
	// t2, _ := GenerateCounters("XXZZXY")
	t3, _ := GenerateCounters("YYXXYXYZZXYXZZZYZZYY")
	type args struct {
		queue    string
		counters Counters
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	name: "XXYYXZ",
		// 	args: args{
		// 		queue:    "XXYYXZ",
		// 		counters: t1,
		// 	},
		// 	want: "Yes",
		// },
		// {
		// 	name: "XXZZXY",
		// 	args: args{
		// 		queue:    "XXZZXY",
		// 		counters: t2,
		// 	},
		// 	want: "Yes",
		// },
		{
			name: "YYXXYXYZZXYXZZZYZZYY",
			args: args{
				queue:    "YYXXYXYZZXYXZZZYZZYY",
				counters: t3,
			},
			want: "Yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleGeneral(tt.args.queue, tt.args.counters); !reflect.DeepEqual(got.isCorrect(), tt.want) {
				t.Errorf("HandleGeneral() = %v, want %v", got.isCorrect(), tt.want)
			}
		})
	}
}
