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

		for index, patient := range patients {
			if index == 0 {
				continue
			}

		}

		fmt.Println(patients)

		datasetsAmount--
	}
}
