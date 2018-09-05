package generator

import "testing"

func TestGetSTringArrayByText(t *testing.T) {
	var tests = []struct {
		in     string
		first  string
		second string
		third  string
	}{
		{"ほずみん", "ほ", "ず", "み"},
		{"ほずみ", "ほ", "ず", "み"},
		{"ほず", "ほ", "ず", ""},
		{"ほ", "ほ", "", ""},
		{"", "", "", ""},
	}

	for _, tt := range tests {
		arr := getStringArrayByText(tt.in)

		if len(arr) != 3 {
			t.Errorf("len(arr) should be 3, but %v", len(arr))
		}
		if arr[0] != tt.first {
			t.Errorf("arr[0] should be first character, but %v", arr[0])
		}
		if arr[1] != tt.second {
			t.Errorf("arr[1] should be second character, but %v", arr[1])
		}
		if arr[2] != tt.third {
			t.Errorf("arr[2] should be third character, but %v", arr[2])
		}
	}
}

func TestGetFirstCharacterPosition(t *testing.T) {
	x, y := getFirstCharacterPosition(0)

	if x != 0 {
		t.Errorf("x should be 0, but %v", x)
	}
	if y != 0 {
		t.Errorf("y should be 0, but %v", y)
	}
}
