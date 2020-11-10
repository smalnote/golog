package closure

// Singleton return the singleton value
var Singleton func() *int

func initSingleton() func() *int {
	v := 0
	return func() *int {
		return &v
	}
}

func init() {
	Singleton = initSingleton()
}
