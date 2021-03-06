package escapeanalysis

// run command below at directory this file in
// go build -gcflags " -m -m "
func escapeAnalysis() {
	createUserV1()
	var v *eaUser
	createUserV2(&v)
	println(v.name)
}

type eaUser struct {
	name string
}

//go:noinline
func createUserV1() eaUser {
	u := eaUser{
		name: "Alpha",
	}

	println("&u in createUserV1=", &u)

	return u
}

// **eaUser 指向指针的指针, 可以改变指针的指向
// *eaUser 指向eaUser的指针, 可以改变指向eaUser的内容
//go:noinline
func createUserV2(v **eaUser) eaUser {
	u := eaUser{
		name: "Bravo",
	}

	eafoo(&u) // does not cause u escape to heap

	*v = &u // u escape to heap

	return u
}

//go:noinline
func eafoo(u *eaUser) {
	u.name = "Ceasar"
}
