package colorhash

import (
	"testing"
)

func TestGetColor(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		alpha     bool
		expectedR uint32
		expectedG uint32
		expectedB uint32
		expectedA uint32
	}{
		{
			name:      "Test with empty byte slice and alpha false",
			input:     []byte{},
			alpha:     false,
			expectedR: 58339,
			expectedG: 45232,
			expectedB: 50372,
			expectedA: 0,
		},
		{
			name:      "Test with empty byte slice and alpha true",
			input:     []byte{},
			alpha:     true,
			expectedR: 58339,
			expectedG: 45232,
			expectedB: 50372,
			expectedA: 16962,
		},
		{
			name:      "Test with 'hello' and alpha false",
			input:     []byte("hello"),
			alpha:     false,
			expectedR: 11308,
			expectedG: 62194,
			expectedB: 19789,
			expectedA: 0,
		},
		{
			name:      "Test with 'hello' and alpha true",
			input:     []byte("hello"),
			alpha:     true,
			expectedR: 11308,
			expectedG: 62194,
			expectedB: 19789,
			expectedA: 47802,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			color := GetColor(tt.input, tt.alpha)
			r, g, b, a := color.RGBA()

			if r != tt.expectedR {
				t.Errorf("Got R: %d, expected: %d", r, tt.expectedR)
			}
			if g != tt.expectedG {
				t.Errorf("Got G: %d, expected: %d", g, tt.expectedG)
			}
			if b != tt.expectedB {
				t.Errorf("Got B: %d, expected: %d", b, tt.expectedB)
			}
			if a != tt.expectedA {
				t.Errorf("Got A: %d, expected: %d", a, tt.expectedA)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		alpha    bool
		expected int64
	}{
		{
			name:     "Test with empty byte slice",
			input:    []byte{},
			expected: 14921924,
		},
		{
			name:     "Test with 'hello'",
			input:    []byte("hello"),
			expected: 2945613,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cint, err := GetInt(tt.input)

			if err != nil {
				t.Error(err)
			}
			if cint != tt.expected {
				t.Errorf("Got %d, expected: %d", cint, tt.expected)
			}
		})
	}
}

func TestGetIntA(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		bitSize  int
		expected int64
	}{
		{
			name:     "Test with empty byte slice",
			input:    []byte{},
			bitSize:  64,
			expected: 3820012610,
		},
		{
			name:     "Test with 'hello'",
			input:    []byte("hello"),
			bitSize:  64,
			expected: 754077114,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cint, err := GetIntA(tt.input)

			if err != nil {
				t.Error(err)
			}
			if cint != tt.expected {
				t.Errorf("Got %d, expected: %d", cint, tt.expected)
			}
		})
	}
}

func TestRGB(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test with empty string",
			input:    "",
			expected: "E3B0C4",
		},
		{
			name:     "Test with 'hello'",
			input:    "hello",
			expected: "2CF24D",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RGB(tt.input)
			if result != tt.expected {
				t.Errorf("Got: %s, expected: %s", result, tt.expected)
			}
		})
	}
}

func TestRGBA(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test with empty string",
			input:    "",
			expected: "E3B0C442",
		},
		{
			name:     "Test with 'hello'",
			input:    "hello",
			expected: "2CF24DBA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RGBA(tt.input)
			if result != tt.expected {
				t.Errorf("Got: %s, expected: %s", result, tt.expected)
			}
		})
	}
}

func TestHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test with empty string",
			input:    "",
			expected: "#E3B0C4",
		},
		{
			name:     "Test with 'hello'",
			input:    "hello",
			expected: "#2CF24D",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HTML(tt.input)
			if result != tt.expected {
				t.Errorf("Got: %s, expected: %s", result, tt.expected)
			}
		})
	}
}

func TestHTMLA(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test with empty string",
			input:    "",
			expected: "#E3B0C442",
		},
		{
			name:     "Test with 'hello'",
			input:    "hello",
			expected: "#2CF24DBA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HTMLA(tt.input)
			if result != tt.expected {
				t.Errorf("Got: %s, expected: %s", result, tt.expected)
			}
		})
	}
}

func TestRGBFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test with empty string",
			input:    "",
			expected: "E3B0C4",
		},
		{
			name:     "Test with 'hello'",
			input:    "hello",
			expected: "2CF24D",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RGBFromString(tt.input)
			if result != tt.expected {
				t.Errorf("Got: %s, expected: %s", result, tt.expected)
			}
		})
	}
}

func TestRGBAFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test with empty string",
			input:    "",
			expected: "E3B0C442",
		},
		{
			name:     "Test with 'hello'",
			input:    "hello",
			expected: "2CF24DBA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RGBAFromString(tt.input)
			if result != tt.expected {
				t.Errorf("Got: %s, expected: %s", result, tt.expected)
			}
		})
	}
}

func TestHTMLFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		a        bool
		expected string
	}{
		{
			name:     "Test with empty string",
			input:    "",
			expected: "#E3B0C4",
		},
		{
			name:     "Test with 'hello'",
			input:    "hello",
			expected: "#2CF24D",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HTMLFromString(tt.input)
			if result != tt.expected {
				t.Errorf("Got: %s, expected: %s", result, tt.expected)
			}
		})
	}
}

func TestHTMLAFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		a        bool
		expected string
	}{
		{
			name:     "Test with empty string",
			input:    "",
			expected: "#E3B0C442",
		},
		{
			name:     "Test with 'hello'",
			input:    "hello",
			expected: "#2CF24DBA",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HTMLAFromString(tt.input)
			if result != tt.expected {
				t.Errorf("Got: %s, expected: %s", result, tt.expected)
			}
		})
	}
}

func TestRGBFromAny(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name:     "Test with empty string",
			input:    "",
			expected: "#E3B0C4",
		},
		{
			name:     "Test with 'hello'",
			input:    "hello",
			expected: "#2CF24D",
		},
		{
			name:     "Test with byte slice of hello",
			input:    []byte("hello"),
			expected: "#0382C6",
		},
		{
			name:     "Test with integer 100",
			input:    100,
			expected: "#AD5736",
		},
		{
			name: "Test with map[string]string",
			input: map[string]string{
				"hello":       "hello there",
				"how are you": "good thanks",
			},
			expected: "#3D3D2A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RGBFromAny(tt.input)
			if result != tt.expected {
				t.Errorf("Got: %s, expected: %s", result, tt.expected)
			}
		})
	}
}

func TestRGBAFromAny(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name:     "Test with empty string",
			input:    "",
			expected: "#E3B0C442",
		},
		{
			name:     "Test with 'hello'",
			input:    "hello",
			expected: "#2CF24DBA",
		},
		{
			name:     "Test with byte slice of hello",
			input:    []byte("hello"),
			expected: "#0382C681",
		},
		{
			name:     "Test with integer 100",
			input:    100,
			expected: "#AD573668",
		},
		{
			name: "Test with map[string]string",
			input: map[string]string{
				"hello":       "hello there",
				"how are you": "good thanks",
			},
			expected: "#3D3D2A98",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RGBAFromAny(tt.input)
			if result != tt.expected {
				t.Errorf("Got: %s, expected: %s", result, tt.expected)
			}
		})
	}
}
