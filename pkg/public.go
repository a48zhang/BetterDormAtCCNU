package pkg

import "strconv"

func Get5thBit(uid string) (string, int) {
	role := "???"
	// 区分都是学号第五位，本科是2，硕士是1，博士是0，工号是6或9
	id, _ := strconv.Atoi(uid)
	if id <= 1000000000 || id >= 9999999999 {
		role = "Are you a test user?"
	} else {
		b := (id / 1000000) % 10
		if b == 1 {
			role = "Master"
		} else if b == 2 {
			role = "Bachelor"
		} else if b == 0 {
			role = "Doctor"
		} else if b == 6 || b == 9 {
			role = "Teacher"
		} else {
			role = "Are you a test user?"
		}
	}
	return role, id
}
