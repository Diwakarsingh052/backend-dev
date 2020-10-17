package main
// q3-Default way of passing parameters to a function
type user struct {
	name string
	age  int
}

func defaultParam(user user) user {
	user.name = "Ajay"
	user.age = 29
	return user
}

func main() {
	var u user
	_ = defaultParam(u)
}
