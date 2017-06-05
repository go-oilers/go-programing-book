package concat_test

// Exercise 1.3: Experiment to measure the difference in running time bet ween our potentially
// inefficient versions and the one that uses strings.Join. Section 1.6 illustrates part of the
// time package, and Section 11.4 shows how to write benchmark tests for systematic performance
// evaluation.Exercise 1.3: Experiment to measure the dif ference in running time bet ween our potentially
// inefficient versions and the one that uses strings.Join. Section 1.6 illustrates part of the
// time package, and Section 11.4 shows how to write benchmark tests for systematic performance
// /evaluation.

import (
	"strings"
	"testing"
)

var args = []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog", "10", "100", "1000", "10000"}

func concat(args []string) {
	var final, temp = "", ""
	for _, str := range args {
		final += temp + str
		temp = ""
	}
}

func Testconcat(t *testing.B) {

	for i := 0; i < t.N; i++ {
		concat(args)
	}
}

func TestJoin(t *testing.B) {

	for i := 0; i < t.N; i++ {
		strings.Join(args, " ")
	}

}
