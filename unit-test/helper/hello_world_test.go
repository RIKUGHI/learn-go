package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Bambang",
			request: "Bambang",
		},
		{
			name:    "Cahyo",
			request: "Cahyo",
		},
		{
			name:    "Budi",
			request: "Budi Nugraha",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Bambang", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bambang")
		}
	})
	b.Run("Cahyo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Cahyo")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bambang")
	}
}

func BenchmarkHelloWorldBambang(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bambang")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Bambang",
			request:  "Bambang",
			expected: "Hello Bambang",
		},
		{
			name:     "Cahyo",
			request:  "Cahyo",
			expected: "Hello Cahyo",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Bambang", func(t *testing.T) {
		result := HelloWorld("Bambang")
		require.Equal(t, "Hello Bambang", result, "Result must be 'Hello Bambang'")
	})
	t.Run("Cahyo", func(t *testing.T) {
		result := HelloWorld("Cahyo")
		require.Equal(t, "Hello Cahyo", result, "Result must be 'Hello Cahyo'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
	// after
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bambang")

	require.Equal(t, "Hello Bambang", result, "Result must be 'Hello Bambang'")
	fmt.Println("TestHelloWorldAssert with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Bambang")

	assert.Equal(t, "Hello Bambang", result, "Result must be 'Hello Bambang'")
	fmt.Println("TestHelloWorldAssert with Assert Done")
}

func TestHelloWorldBambang(t *testing.T) {
	result := HelloWorld("Bambang")

	if result != "Hello Bambang" {
		t.Error("Result must be Hello Bambang")
	}

	fmt.Println("TestHelloWorldBambang Done")
}

func TestHelloWordCahyo(t *testing.T) {
	result := HelloWorld("Cahyo")

	if result != "Hello Cahyo" {
		t.Fatal("Result must be Hello Cahyo")
	}

	fmt.Println("TestHelloWorldCahyo Done")
}
