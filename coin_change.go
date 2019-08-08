package coding_tasks

func ChangeMoney(total int) (quarters, dimes, nickels, pennies int) {
	for total > 0 {
		if total > 25 {
			quarters = total / 25
			total = total % 25
			continue
		} else if total > 10 {
			dimes = total / 10
			total = total % 10
		} else if total > 5 {
			nickels = total / 5
			total = total % 5
		} else {
			pennies = total
			total = 0
		}
	}
	return
}