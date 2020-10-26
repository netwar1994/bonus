package main

func main() {
	purchase := 3333_33
	percent := 1_00
	limit := 100_00

	bonus := (purchase / 100_00) * percent
	if bonus > limit {
		bonus = limit
	}
	println(bonus)
}
