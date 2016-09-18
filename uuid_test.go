package uuid

import "testing"

func TestNew(t *testing.T) {
	a := New()
	for i := 0; i < 10; i++ {
		b := New()
		if b == a {
			t.Fatal("should not generate duplicate uuid")
		}
		if len(b) != ByteSize*2+4 {
			t.Fatal("should be a standardized length")
		}
	}
}

func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = New()
	}
}
