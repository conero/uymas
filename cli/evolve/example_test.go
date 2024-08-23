package evolve

import (
	"fmt"
	"log"
)

func ExampleNewEvolve() {
	// command struct
	type test struct {
		Command
	}

	evl := NewEvolve()

	// register func
	evl.Command(func() {
		fmt.Println("Evolution For Index.")
	}, "index")

	// register struct
	evl.Command(new(test), "test", "t")
	log.Fatal(evl.Run())
}

//func TestExample(t *testing.T) {
//	ExampleNewEvolve()
//}
