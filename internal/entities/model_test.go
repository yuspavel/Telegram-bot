package entities

import (
	"testing"
)

func TestEntity(t *testing.T) {
	tests := []struct {
		name     string
		quote    string
		author   string
		expected string
	}{
		{
			name:     "Valide quote",
			quote:    "To be or not to be",
			author:   "Shakespeare",
			expected: "To be or not to be",
		},
		{
			name:     "Empty quote",
			quote:    "",
			author:   "Unknown",
			expected: "",
		},
		{
			name:     "Qoute with special character",
			quote:    "To be' or not to be",
			author:   "Shakespeare",
			expected: "To be' or not to be",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quote := &Quote{
				Text:   tt.quote,
				Author: tt.author}
			if quote.Text != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, quote.Text)
			}
			if quote.Author != tt.author {
				t.Errorf("Expected author %s, got %s", tt.author, quote.Author)
			}
			/*if !strings.Contains(quote.Text, "'") {
				t.Errorf("Expected %s, got %s", tt.expected, quote.Text)
			}*/

		})
	}
}
