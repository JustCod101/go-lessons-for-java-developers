package examples

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// 基础测试
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}
}

func TestAddTableDriven(t *testing.T) {
	// 表驱动测试 (Table-Driven Tests)
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 1, 2, 3},
		{"negative", -1, -2, -3},
		{"zero", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
