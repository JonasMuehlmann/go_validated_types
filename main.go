package main

const runs = 10000000

var results = make([]int, runs)

func main() {
	//benchRegular()
	benchValidated()
}

func benchValidated() {
	for i := 0; i < runs; i++ {
		foo := NewSetOnceIntBare()
		foo.Set(i)
		results[i] = foo.Get()
	}
}
func benchRegular() {
	for i := 0; i < runs; i++ {
		foo := 0
		foo = i
		results[i] = foo
	}
}
