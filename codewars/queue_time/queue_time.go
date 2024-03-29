package queue_time

/* There is a queue for the self-checkout tills at the supermarket. Your task is write a function to calculate the total time required for all the customers to check out!

input ->
	customers: an array of positive integers representing the queue. Each integer represents a customer, and its value is the amount of time they require to check out.
	n: a positive integer, the number of checkout tills.

output ->
	The function should return an integer, the total time required. */

func hasCustomers(customers []int) bool {
	return len(customers) > 0
}

func mapIsEmpty(m map[int]int) bool {
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func MySolution(customers []int, n int) int {
	t := 0
	checkTills := make(map[int]int, n)

	for i := 0; i < n; i++ {
		if hasCustomers(customers) {
			checkTills[i] = customers[0]
			customers = append(customers[1:])
		}
	}

	for !mapIsEmpty(checkTills) || hasCustomers(customers) {
		for i := range checkTills {
			if checkTills[i] == 0 {
				if hasCustomers(customers) {
					checkTills[i] = customers[0]
					checkTills[i] -= 1
					customers = append(customers[1:])
				}
			} else {
				checkTills[i] -= 1
			}
		}
		t++
	}

	return t
}
