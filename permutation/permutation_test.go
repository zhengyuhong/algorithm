package permutation

import (
    "testing"
)

func TestNewAlphaNode(t *testing.T) {
    set1 := []any{1, 2, 3}
    set2 := []any{"x", "y", "z"}
    set3 := []any{"A", "B"}
    sets := [][]any{set1, set2, set3}
    ret := Product(sets)
    if len(ret) != 3 * 3 * 2 {
        t.Fail()
        t.Log()
    }
}
