package sum

import "testing"

func TestInts(t *testing.T) {
	t.Errorf("this test failed because I said so")
	t.Fatalf("this test failed and stopped running")
}
