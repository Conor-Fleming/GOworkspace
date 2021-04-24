package saying

import (
	"fmt"
	"testing"
)

func TestGreetings(t *testing.T) {
	s := Greetings("Ciara")
	if s != "Welcome enjoy this greeting Ciara" {
		t.Error("got", s, "expected", "Welcome, enjoy this greeting Ciara")
	}
}

func ExampleGreetings() {
	fmt.Println(Greetings("Ciara"))
	//Output:
	//Welcome, enjoy this greeting Ciara
}

func BenchmarkGreetings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greetings("Ciara")
	}
}
