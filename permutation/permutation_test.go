package util

import (
    "testing"
)

func TestNewAlphaNode(t *testing.T) {
    set1 := []any{1, 2, 3}
    set2 := []any{"x", "y", "z"}
    set3 := []any{"A", "B"}
    sets := [][]any{set1, set2, set3}
    Product(sets)
}
