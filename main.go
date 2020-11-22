package main

func main() {
	purchase := 3333_33
	percent := 1
	limit := 100

	bonus := (purchase / 100_00) * percent
	if bonus > limit {
		bonus = limit
	}
	println(bonus)
}
