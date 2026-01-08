package main

import (
	"fmt"
	"math"
)

// Vector2D represents a 2D vector
type Vector2D struct {
	X, Y float64
}

// Add adds two vectors
func (v Vector2D) Add(other Vector2D) Vector2D {
	return Vector2D{v.X + other.X, v.Y + other.Y}
}

// Magnitude returns the length of the vector
func (v Vector2D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// DotProduct calculates dot product
func (v Vector2D) DotProduct(other Vector2D) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Matrix2x2 represents a 2x2 matrix
type Matrix2x2 struct {
	A, B, C, D float64
}

// Multiply multiplies two matrices
func (m Matrix2x2) Multiply(other Matrix2x2) Matrix2x2 {
	return Matrix2x2{
		A: m.A*other.A + m.B*other.C,
		B: m.A*other.B + m.B*other.D,
		C: m.C*other.A + m.D*other.C,
		D: m.C*other.B + m.D*other.D,
	}
}

// Determinant calculates the determinant
func (m Matrix2x2) Determinant() float64 {
	return m.A*m.D - m.B*m.C
}

// MathUtils contains utility functions
type MathUtils struct{}

// Factorial calculates n!
func (MathUtils) Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * MathUtils{}.Factorial(n-1)
}

// IsPrime checks if number is prime
func (MathUtils) IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Fibonacci generates fibonacci sequence up to n terms
func (MathUtils) Fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	
	fib := make([]int, n)
	fib[0], fib[1] = 0, 1
	
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

// GCD calculates greatest common divisor
func (MathUtils) GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println("=== Go Math Project Demo ===\n")


	// Vector operations
	fmt.Println("--- Vector Operations ---")
	v1 := Vector2D{3, 4}
	v2 := Vector2D{1, 2}
	fmt.Printf("v1: (%.1f, %.1f)\n", v1.X, v1.Y)
	fmt.Printf("v2: (%.1f, %.1f)\n", v2.X, v2.Y)
	fmt.Printf("v1 + v2: (%.1f, %.1f)\n", v1.Add(v2).X, v1.Add(v2).Y)
	fmt.Printf("Magnitude of v1: %.2f\n", v1.Magnitude())
	fmt.Printf("Dot product v1·v2: %.1f\n\n", v1.DotProduct(v2))

	// Matrix operations
	fmt.Println("--- Matrix Operations ---")
	m1 := Matrix2x2{1, 2, 3, 4}
	m2 := Matrix2x2{2, 0, 1, 2}
	fmt.Printf("Matrix 1: [%.0f %.0f; %.0f %.0f]\n", m1.A, m1.B, m1.C, m1.D)
	fmt.Printf("Matrix 2: [%.0f %.0f; %.0f %.0f]\n", m2.A, m2.B, m2.C, m2.D)
	result := m1.Multiply(m2)
	fmt.Printf("M1 × M2: [%.0f %.0f; %.0f %.0f]\n", result.A, result.B, result.C, result.D)
	fmt.Printf("Determinant of M1: %.1f\n\n", m1.Determinant())

	// Math utilities
	fmt.Println("--- Math Utilities ---")
	utils := MathUtils{}
	
	fmt.Printf("Factorial of 5: %d\n", utils.Factorial(5))
	fmt.Printf("Is 17 prime? %t\n", utils.IsPrime(17))
	fmt.Printf("Is 20 prime? %t\n", utils.IsPrime(20))
	fmt.Printf("Fibonacci(10): %v\n", utils.Fibonacci(10))
	fmt.Printf("GCD(48, 18): %d\n", utils.GCD(48, 18))
}