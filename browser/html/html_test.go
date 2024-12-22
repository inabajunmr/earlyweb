package html

import (
	"reflect"
	"testing"
)

func TestParseHtml(t *testing.T) {
	tests := []struct {
		input    string
		expected []Tag
	}{
		{
			input: "<DIV><A>Click here</A></DIV>",
			expected: []Tag{
				{Name: "DIV", Children: []Tag{
					{Name: "A", Children: nil, Text: "Click here"},
				}},
			},
		},
		{
			input: "<DIV><A>Click here</A><P>Hello world</P></DIV>",
			expected: []Tag{
				{Name: "DIV", Children: []Tag{
					{Name: "A", Children: nil, Text: "Click here"},
					{Name: "P", Children: nil, Text: "Hello world"},
				}},
			},
		},
		{
			input: "<SPAN>No children</SPAN>",
			expected: []Tag{
				{Name: "SPAN", Children: nil, Text: "No children"},
			},
		},
		{
			input: "<SPAN>hello</SPAN><SPAN>world</SPAN>",
			expected: []Tag{
				{Name: "SPAN", Children: nil, Text: "hello"},
				{Name: "SPAN", Children: nil, Text: "world"},
			},
		},
		{
			input: "<SPAN><ISINDEX>aaa</SPAN>",
			expected: []Tag{
				{Name: "SPAN", Children: []Tag{{Name: "ISINDEX", Children: nil}}, Text: "aaa"},
			},
		},
	}

	for i, test := range tests {
		result := ParseHtml(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Test %d failed: expected %+v, got %+v\n", i+1, test.expected, result)
		}
	}
}
