package vm

import "testing"

func TestNewMemory(t *testing.T) {

	tests := []struct {
		input int
		want  int
	}{
		{input: 3, want: 4},
	}

	var m *Memory
	var size int

	for _, test := range tests {
		m = NewMemory(test.input)
		size = len(m.mem)
		if size != test.want {
			t.Fatalf("expected %d got %d", test.want, size)
		}
	}

}
