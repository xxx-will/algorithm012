package main

func lemonadeChange(bills []int) bool {

	m5 := 0
	m10 := 0

	for _, v := range bills {

		switch v {
		case 5:
			m5 = m5 + 1
		case 10:
			m10 = m10 + 1
			m5 = m5 - 1
		case 20:
			//贪心
			if m10 > 0 {
				m10 = m10 - 1
				m5 = m5 - 1
			} else {
				m5 = m5 - 3
			}
		}
		if m5 < 0 {
			return false
		}
	}

	return true
}
