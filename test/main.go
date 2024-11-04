// Package testing demonstrates comprehensive testing concepts in Go
package testing

import (
	"errors"
	"math"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// Calculator represents a simple calculator with basic operations
type Calculator struct {
	memory float64
	mu     sync.Mutex
}

// Add performs addition and stores result in memory
func (c *Calculator) Add(a, b float64) float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.memory = a + b
	return c.memory
}

// Multiply performs multiplication and stores result in memory
func (c *Calculator) Multiply(a, b float64) float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.memory = a * b
	return c.memory
}

// SquareRoot calculates the square root of a number
func (c *Calculator) SquareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("cannot calculate square root of negative number")
	}
	return math.Sqrt(n), nil
}

// DataStore represents an interface for data storage
type DataStore interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

// MockDataStore implements DataStore interface for testing
type MockDataStore struct {
	mock.Mock
}

func (m *MockDataStore) Get(key string) (string, error) {
	args := m.Called(key)
	return args.String(0), args.Error(1)
}

func (m *MockDataStore) Set(key, value string) error {
	args := m.Called(key, value)
	return args.Error(0)
}

// Basic unit test example
func TestAdd(t *testing.T) {
	calc := &Calculator{}

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"zero case", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%v, %v) = %v; want %v",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Testify assert example
func TestMultiplyWithAssert(t *testing.T) {
	calc := &Calculator{}
	result := calc.Multiply(4, 5)
	assert.Equal(t, 20.0, result, "multiplication should work correctly")
	assert.NotEqual(t, 0.0, result, "result should not be zero")
}

// Testify require example
func TestSquareRootWithRequire(t *testing.T) {
	calc := &Calculator{}
	result, err := calc.SquareRoot(16)
	require.NoError(t, err, "should not return error for positive number")
	require.Equal(t, 4.0, result, "square root of 16 should be 4")

	// This will stop the test immediately if there's an error
	_, err = calc.SquareRoot(-1)
	require.Error(t, err, "should return error for negative number")
}

// Mock example
func TestWithMockDataStore(t *testing.T) {
	mockStore := new(MockDataStore)

	// Set up expectations
	mockStore.On("Get", "test_key").Return("test_value", nil)
	mockStore.On("Set", "test_key", "new_value").Return(nil)

	// Test Get
	value, err := mockStore.Get("test_key")
	assert.NoError(t, err)
	assert.Equal(t, "test_value", value)

	// Test Set
	err = mockStore.Set("test_key", "new_value")
	assert.NoError(t, err)

	// Verify that all expectations were met
	mockStore.AssertExpectations(t)
}

// Test suite example
type CalculatorTestSuite struct {
	suite.Suite
	calc *Calculator
}

func (suite *CalculatorTestSuite) SetupTest() {
	suite.calc = &Calculator{}
}

func (suite *CalculatorTestSuite) TestAddInSuite() {
	result := suite.calc.Add(2, 3)
	suite.Equal(5.0, result)
}

func (suite *CalculatorTestSuite) TestMultiplyInSuite() {
	result := suite.calc.Multiply(2, 3)
	suite.Equal(6.0, result)
}

func TestCalculatorSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}

// Benchmark example
func BenchmarkAdd(b *testing.B) {
	calc := &Calculator{}
	for i := 0; i < b.N; i++ {
		calc.Add(2, 3)
	}
}

// Parallel benchmark example
func BenchmarkAddParallel(b *testing.B) {
	calc := &Calculator{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			calc.Add(2, 3)
		}
	})
}

// Example of sub-benchmarks
func BenchmarkCalculator(b *testing.B) {
	calc := &Calculator{}

	b.Run("Add", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			calc.Add(2, 3)
		}
	})

	b.Run("Multiply", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			calc.Multiply(2, 3)
		}
	})
}
