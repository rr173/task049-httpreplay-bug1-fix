package replay

import "testing"

func TestProbe_EmptyKeysReturnsEmptySlice(t *testing.T) {
	r := New()
	keys := r.Keys()
	if keys == nil {
		t.Fatal("Keys returned nil; want a non-nil empty slice")
	}
	if len(keys) != 0 {
		t.Fatalf("len(Keys()) = %d, want 0", len(keys))
	}
}
