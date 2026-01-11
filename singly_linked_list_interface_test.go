package main

import (
	"reflect"
	"testing"
)

// Helper function to convert list to slice for easy comparison
func (l *List) toSlice() []interface{} {
	result := []interface{}{}
	cur := l
	for cur != nil && cur.val != nil {
		result = append(result, cur.val)
		cur = cur.next
	}
	return result
}

// Helper function to create a list from values
func createList(values []interface{}) *List {
	if len(values) == 0 {
		return &List{}
	}
	list := &List{val: values[0]}
	cur := list
	for i := 1; i < len(values); i++ {
		cur.next = &List{val: values[i]}
		cur = cur.next
	}
	return list
}

func TestAppend(t *testing.T) {
	tests := []struct {
		name     string
		values   []interface{}
		expected []interface{}
	}{
		{
			name:     "append to empty list",
			values:   []interface{}{"first"},
			expected: []interface{}{"first"},
		},
		{
			name:     "append multiple values",
			values:   []interface{}{"a", "b", "c"},
			expected: []interface{}{"a", "b", "c"},
		},
		{
			name:     "append integers",
			values:   []interface{}{1, 2, 3},
			expected: []interface{}{1, 2, 3},
		},
		{
			name:     "append mixed types",
			values:   []interface{}{"string", 42, 3.14},
			expected: []interface{}{"string", 42, 3.14},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &List{}
			for _, val := range tt.values {
				list.append(val)
			}
			got := list.toSlice()
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("append() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPrepend(t *testing.T) {
	tests := []struct {
		name     string
		initial  []interface{}
		prepend  interface{}
		expected []interface{}
	}{
		{
			name:     "prepend to empty list",
			initial:  []interface{}{},
			prepend:  "first",
			expected: []interface{}{"first"},
		},
		{
			name:     "prepend to list with one element",
			initial:  []interface{}{"second"},
			prepend:  "first",
			expected: []interface{}{"first", "second"},
		},
		{
			name:     "prepend to list with multiple elements",
			initial:  []interface{}{"b", "c"},
			prepend:  "a",
			expected: []interface{}{"a", "b", "c"},
		},
		{
			name:     "prepend multiple times",
			initial:  []interface{}{},
			prepend:  "c",
			expected: []interface{}{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := createList(tt.initial)
			if tt.name == "prepend multiple times" {
				// Special case: prepend multiple items
				list = &List{}
				list.prepend("c")
				list.prepend("b")
				list.prepend("a")
			} else {
				list.prepend(tt.prepend)
			}
			got := list.toSlice()
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("prepend() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		name     string
		list     *List
		expected int
	}{
		{
			name:     "empty list (nil val)",
			list:     &List{},
			expected: 1, // A List{} node exists, even with nil val
		},
		{
			name:     "single element",
			list:     createList([]interface{}{"one"}),
			expected: 1,
		},
		{
			name:     "multiple elements",
			list:     createList([]interface{}{1, 2, 3, 4, 5}),
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.length()
			if got != tt.expected {
				t.Errorf("length() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestContains(t *testing.T) {
	list := createList([]interface{}{"apple", "banana", "cherry"})

	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{
			name:     "contains first element",
			value:    "apple",
			expected: true,
		},
		{
			name:     "contains middle element",
			value:    "banana",
			expected: true,
		},
		{
			name:     "contains last element",
			value:    "cherry",
			expected: true,
		},
		{
			name:     "does not contain value",
			value:    "orange",
			expected: false,
		},
		{
			name:     "empty list",
			value:    "anything",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var testList *List
			if tt.name == "empty list" {
				testList = &List{}
			} else {
				testList = list
			}
			got := testList.contains(tt.value)
			if got != tt.expected {
				t.Errorf("contains(%v) = %v, want %v", tt.value, got, tt.expected)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name           string
		initial        []interface{}
		deleteValue    interface{}
		expected       []interface{}
		shouldDelete   bool
	}{
		{
			name:         "delete first element",
			initial:      []interface{}{"a", "b", "c"},
			deleteValue:  "a",
			expected:     []interface{}{"b", "c"},
			shouldDelete: true,
		},
		{
			name:         "delete middle element",
			initial:      []interface{}{"a", "b", "c"},
			deleteValue:  "b",
			expected:     []interface{}{"a", "c"},
			shouldDelete: true,
		},
		{
			name:         "delete last element",
			initial:      []interface{}{"a", "b", "c"},
			deleteValue:  "c",
			expected:     []interface{}{"a", "b"},
			shouldDelete: true,
		},
		{
			name:         "delete only element",
			initial:      []interface{}{"only"},
			deleteValue:  "only",
			expected:     []interface{}{},
			shouldDelete: true,
		},
		{
			name:         "delete non-existent element",
			initial:      []interface{}{"a", "b", "c"},
			deleteValue:  "z",
			expected:     []interface{}{"a", "b", "c"},
			shouldDelete: false,
		},
		{
			name:         "delete from empty list",
			initial:      []interface{}{},
			deleteValue:  "anything",
			expected:     []interface{}{},
			shouldDelete: false,
		},
		{
			name:         "delete duplicate (first occurrence)",
			initial:      []interface{}{"a", "a", "b"},
			deleteValue:  "a",
			expected:     []interface{}{"a", "b"},
			shouldDelete: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := createList(tt.initial)
			list.delete(tt.deleteValue)
			got := list.toSlice()
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("delete(%v) = %v, want %v", tt.deleteValue, got, tt.expected)
			}
		})
	}
}

func TestCombinedOperations(t *testing.T) {
	t.Run("append then delete then prepend", func(t *testing.T) {
		list := &List{}
		list.append("b")
		list.append("c")
		list.delete("b")
		list.prepend("a")
		expected := []interface{}{"a", "c"}
		got := list.toSlice()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("combined operations = %v, want %v", got, expected)
		}
	})

	t.Run("multiple prepends and appends", func(t *testing.T) {
		list := &List{}
		list.append("middle")
		list.prepend("first")
		list.append("last")
		expected := []interface{}{"first", "middle", "last"}
		got := list.toSlice()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("multiple operations = %v, want %v", got, expected)
		}
	})

	t.Run("delete all elements one by one", func(t *testing.T) {
		list := createList([]interface{}{"a", "b", "c"})
		list.delete("b")
		list.delete("a")
		list.delete("c")
		expected := []interface{}{}
		got := list.toSlice()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("delete all = %v, want %v", got, expected)
		}
	})
}

func TestIntegerValues(t *testing.T) {
	t.Run("append integers", func(t *testing.T) {
		list := &List{}
		list.append(1)
		list.append(2)
		list.append(3)
		expected := []interface{}{1, 2, 3}
		got := list.toSlice()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("integers = %v, want %v", got, expected)
		}
	})

	t.Run("delete integer", func(t *testing.T) {
		list := createList([]interface{}{1, 2, 3})
		list.delete(2)
		expected := []interface{}{1, 3}
		got := list.toSlice()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("delete integer = %v, want %v", got, expected)
		}
	})
}
