package base

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var shortStr = "my_string" + string(make([]byte, 16))

func BenchmarkConcatShort10(b *testing.B) {
	concaIter(b, shortStr, 10)
	fmt.Println()
}

func BenchmarkConcatShort100(b *testing.B) {
	concaIter(b, shortStr, 100)
	fmt.Println()
}

func BenchmarkConcatShort1000(b *testing.B) {
	concaIter(b, shortStr, 1000)
	fmt.Println()
}

func BenchmarkConcatShort10000(b *testing.B) {
	concaIter(b, shortStr, 10000)
	fmt.Println()

}

func BenchmarkConcatShort100000(b *testing.B) {
	concaIter(b, shortStr, 100000)
	fmt.Println()
}


var longStr = "my_string" + string(make([]byte, 256))

func BenchmarkConcatShort10Long(b *testing.B) {
	concaIter(b, longStr, 10)
	fmt.Println()
}

func BenchmarkConcatShort100Long(b *testing.B) {
	concaIter(b, longStr, 100)
	fmt.Println()
}

func BenchmarkConcatShort1000Long(b *testing.B) {
	concaIter(b, longStr, 1000)
	fmt.Println()
}

func BenchmarkConcatShort10000Long(b *testing.B) {
	concaIter(b, longStr, 10000)
	fmt.Println()

}

func BenchmarkConcatShort100000Long(b *testing.B) {
	concaIter(b, longStr, 100000)
	fmt.Println()
}

func concaIter(b *testing.B, str string, i int) {
	b.Run("Plus", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			var s string
			for is := 0; is < i; is++ {
				s += str
			}
			_ = s

		}
	})

	b.Run("Buffer", func(b *testing.B) {
		var buf bytes.Buffer
		for n := 0; n < b.N; n++ {
			buf.Reset()
			for s := 0; s < i; s++ {
				buf.WriteString(str)
			}
			_ = buf.String()

		}
	})

	b.Run("Join", func(b *testing.B) {
		stringStr := make([]string, i)
		for sp := 0; sp < i; sp++ {
			stringStr[sp] = str
		}
		b.ResetTimer()

		for np := 0; np < b.N; np++ {
			_ = strings.Join(stringStr, "")

		}
	})

	b.Run("Builder", func(b *testing.B) {
		var buf strings.Builder
		for n := 0; n < b.N; n++ {
			buf.Reset()
			for s := 0; s < i; s++ {
				buf.WriteString(str)
			}
			_ = buf.String()

		}
	})

	b.Run("Sprint", func(b *testing.B) {
		sprintArgs := make([]interface{}, i)

		for i := 0; i < i; i++ {
			sprintArgs[i] = str
		}

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			_ = fmt.Sprint(sprintArgs...)
		}
	})

	b.Run("Sprintf", func(b *testing.B) {
		sprintfArgs := make([]interface{}, i)

		for i := 0; i < i; i++ {
			sprintfArgs[i] = str
		}
		repeat := strings.Repeat("%s", i)

		b.ResetTimer()

		for n := 0; n < b.N; n++ {
			_ = fmt.Sprintf(repeat,sprintfArgs...)
		}
	})
}
