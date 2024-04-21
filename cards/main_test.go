package main

import "testing"

func TestFindFriendWithMaxCardIndex(t *testing.T) {
	type args struct {
		friendsCards map[int]int
	}
	tests := []struct {
		name            string
		args            args
		wantFriendIndex int
	}{
		{
			name: "test",
			args: args{
				friendsCards: map[int]int{0: 3, 1: 3, 2: 2, 3: 6, 4: 5},
			},
			wantFriendIndex: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFriendIndex := FindFriendWithMaxCardIndex(tt.args.friendsCards); gotFriendIndex != tt.wantFriendIndex {
				t.Errorf("FindFriendWithMaxCardIndex() = %v, want %v", gotFriendIndex, tt.wantFriendIndex)
			}
		})
	}
}
