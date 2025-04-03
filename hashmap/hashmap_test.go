package hasmap

import (
	"sort"
	"testing"
)

// TestPut tests the Put method of the HashMap.
// It verifies that putting elements into the map updates the map correctly
// and that the number of keys in the map is as expected.
//
// Parameters:
//
//	t (*testing.T): The testing context.
func TestPut(t *testing.T) {
	m := Map[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)

	if len(m.Keys()) != 3 {
		t.Errorf("m.Keys() = %v; want 3", len(m.Keys()))
	}
}

// TestGet tests the Get method of the HashMap.
// It verifies that getting elements from the map returns the correct values
// and that the presence of keys is correctly identified.
//
// Parameters:
//
//	t (*testing.T): The testing context.
func TestGet(t *testing.T) {
	m := Map[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	val, ok := m.Get("a")
	if !ok {
		t.Errorf("Get(\"a\") = %v; want true", ok)
	}
	if val != 1 {
		t.Errorf("Get(\"a\") = %v; want 1", val)
	}
	_, ok = m.Get("c")
	if ok {
		t.Errorf("Get(\"c\") = %v; want false", ok)
	}
}

// TestRemove tests the Remove method of the HashMap.
// It verifies that removing elements from the map updates the map correctly
// and that the removed key is no longer present in the map.
//
// Parameters:
//
//	t (*testing.T): The testing context.
func TestRemove(t *testing.T) {
	m := Map[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)
	m.Remove("b")
	if len(m.Keys()) != 1 {
		t.Errorf("m.Keys() = %v; want 1", len(m.Keys()))
	}
	_, ok := m.Get("b")
	if ok {
		t.Errorf("Get(\"b\") = %v; want false", ok)
	}
}

// TestKeys tests the Keys method of the HashMap.
// It verifies that the keys method returns all keys in the map,
// and that the keys are correctly sorted and retrievable.
//
// Parameters:
//
//	t (*testing.T): The testing context.
func TestKeys(t *testing.T) {
	m := Map[string, int]()
	m.Put("a", 1)
	m.Put("b", 2)

	keys := m.Keys()
	if len(keys) != 2 {
		t.Errorf("m.Keys() = %v; want 2", len(keys))
	}

	sort.Strings(keys)
	expectedKeys := []string{"a", "b"}
	for i, key := range keys {
		if key != expectedKeys[i] {
			t.Errorf("keys[%d] = %v; want %v", i, key, expectedKeys[i])
		}

		if _, ok := m.Get(key); !ok {
			t.Errorf("m.Get(%v) = _, %v; want true", key, ok)
		}
	}
}

func TestGetEmpty(t *testing.T) {
	m := Map[string, int]()
	_, ok := m.Get("nonexistent")
	if ok {
		t.Errorf("Get on empty map returned ok=true; want false")
	}
}

func TestPutOverwrite(t *testing.T) {
	m := Map[string, int]()
	m.Put("a", 1)
	m.Put("a", 100)

	val, ok := m.Get("a")
	if !ok {
		t.Errorf("Get(\"a\") = _, %v; want true", ok)
	}
	if val != 100 {
		t.Errorf("Get(\"a\") = %v, _; want 100", val)
	}
	if len(m.Keys()) != 1 {
		t.Errorf("len(m.Keys()) = %v; want 1", len(m.Keys()))
	}
}
