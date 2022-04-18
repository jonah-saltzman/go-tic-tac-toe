package main

func Eval(b *board) int {
	for i := range Conditions {
		win := true
		for _, pos := range Conditions[i] {
			if b[pos] == 0 || b[pos] != b[Conditions[i][0]] {
				win = false
				break
			}
		}
		if win {
			if b[Conditions[i][0]] == 1 {
				return -Max
			} else {
				return Max
			}
		}
	}
	return 0
}

func Over(b *board) (int, bool) {

	over := true

	for i := range Conditions {

		win := true

		for _, pos := range Conditions[i] {
			if b[pos] == 0 {
				win = false
				over = false
				continue
			}
			if b[pos] != b[Conditions[i][0]] {
				win = false
			}
		}

		if win {
			if b[Conditions[i][0]] == 1 {
				return -Max, true
			} else {
				return Max, true
			}
		}
	}
	return 0, over
}
