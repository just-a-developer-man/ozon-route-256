package main

import (
	"testing"
)

func TestMoveRight(t *testing.T) {
	type args struct {
		patients *[]Patient
		index    int
		freesUp  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				patients: &[]Patient{
					Patient{
						id:          0,
						move:        '0',
						appointment: 100000,
					}, Patient{
						id:          3,
						move:        '0',
						appointment: 100000,
					},
					Patient{
						id:          1,
						move:        '0',
						appointment: 100001,
					},
					Patient{
						id:          2,
						move:        '0',
						appointment: 100001,
					}},
				index:   1,
				freesUp: 200000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MoveRight(tt.args.patients, tt.args.index, tt.args.freesUp); (err != nil) != tt.wantErr {
				t.Errorf("MoveRight() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMoveLeft(t *testing.T) {
	var err error
	freesUp := 200000
	patients := []Patient{
		Patient{
			id:          0,
			move:        '0',
			appointment: 1,
		}, Patient{
			id:          3,
			move:        '0',
			appointment: 1,
		},
		Patient{
			id:          1,
			move:        '0',
			appointment: 2,
		},
		Patient{
			id:          2,
			move:        '0',
			appointment: 2,
		},
		Patient{
			id:          2,
			move:        '0',
			appointment: 4,
		},
		Patient{
			id:          2,
			move:        '0',
			appointment: 4,
		}}

	for index := range patients {
		if index == 0 {
			continue
		}
		if !NeedToMove(patients, index) {
			continue
		}
		if MoveLeft(&patients, index-1) != nil {
			err = MoveRight(&patients, index, freesUp)
		}
		if err != nil {
			break
		}
	}
}
