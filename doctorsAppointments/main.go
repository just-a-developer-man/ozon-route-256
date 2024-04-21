package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type Patient struct {
	id          int
	move        rune
	appointment int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Solution(in, out)
}

func Solution(in io.Reader, out io.Writer) {
	var datasetsAmount int
	fmt.Fscan(in, &datasetsAmount)

	for datasetsAmount != 0 {
		var freesUp, patientsAmount int
		var err error
		fmt.Fscan(in, &freesUp, &patientsAmount)

		patients := make([]Patient, 0)

		for id := 0; id < patientsAmount; id++ {
			appointment := 0
			fmt.Fscan(in, &appointment)
			patients = append(patients, Patient{
				id:          id,
				move:        '0',
				appointment: appointment,
			})
		}

		sort.Slice(patients, func(i, j int) bool {
			return patients[i].appointment < patients[j].appointment
		})

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
				fmt.Fprintln(out, "x")
				break
			}
		}

		if err != nil {
			continue
		}

		sort.Slice(patients, func(i, j int) bool {
			return patients[i].id < patients[j].id
		})

		for id := 0; id < patientsAmount; id++ {
			fmt.Fprint(out, string(patients[id].move))
		}

		fmt.Fprintln(out, "")

		datasetsAmount--
	}
}

func NeedToMove(patients []Patient, index int) bool {
	if patients[index].appointment == patients[index-1].appointment {
		return true
	}
	return false
}

func MoveLeft(patients *[]Patient, index int) error {
	if index < 0 {
		return fmt.Errorf("")
	}

	if (*patients)[index].move != '0' {
		return fmt.Errorf("")
	}

	if index == 0 {
		if (*patients)[index].appointment <= 1 {
			return fmt.Errorf("")
		}
		(*patients)[index].appointment--
		(*patients)[index].move = '-'
		return nil
	}

	if (*patients)[index].appointment <= 1 {
		return fmt.Errorf("")
	}

	if (*patients)[index].appointment-(*patients)[index-1].appointment <= 1 {
		err := MoveLeft(patients, index-1)
		if err != nil {
			return fmt.Errorf("")
		}
	}

	(*patients)[index].appointment--
	(*patients)[index].move = '-'
	return nil

}

func MoveRight(patients *[]Patient, index int, freesUp int) error {
	if index >= len(*patients) {
		return fmt.Errorf("")
	}

	if (*patients)[index].move != '0' {
		return fmt.Errorf("")
	}

	if index == len(*patients)-1 {
		if (*patients)[index].appointment >= freesUp {
			return fmt.Errorf("")
		}
		(*patients)[index].appointment++
		(*patients)[index].move = '+'
		return nil
	}

	if (*patients)[index+1].appointment-(*patients)[index].appointment <= 1 {
		err := MoveRight(patients, index+1, freesUp)
		if err != nil {
			return fmt.Errorf("")
		}
	}

	(*patients)[index].appointment++
	(*patients)[index].move = '+'
	return nil
}
