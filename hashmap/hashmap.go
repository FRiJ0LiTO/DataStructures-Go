package hasmap

import "fmt"

type HashMap[K comparable, V any] struct {
	Put      func(K, V)
	Get      func(K) (V, bool)
	Remove   func(K)
	Keys     func() []K
	PrintMap func()
}

// Map creates and returns a new HashMap with generic key and value types.
// The returned HashMap provides methods to put, get, remove keys, retrieve all keys, and print the map.
//
// Type Parameters:
//
//	K (comparable): The type of keys in the HashMap. Must be comparable.
//	V (any): The type of values in the HashMap.
//
// Returns:
//
//	HashMap[K, V]: A new HashMap with the specified key and value types.
func Map[K comparable, V any]() HashMap[K, V] {
	data := make(map[K]V)

	return HashMap[K, V]{
		// Put inserts or updates the value associated with the given key.
		//
		// Parameters:
		//   key (K): The key to insert or update.
		//   value (V): The value to associate with the key.
		Put: func(key K, value V) {
			data[key] = value
		},
		// Get retrieves the value associated with the given key.
		//
		// Parameters:
		//   key (K): The key to look up.
		//
		// Returns:
		//   (V, bool): The value associated with the key and a boolean indicating if the key was found.
		Get: func(key K) (V, bool) {
			v, found := data[key]
			return v, found
		},
		// Remove deletes the value associated with the given key.
		//
		// Parameters:
		//   key (K): The key to delete.
		Remove: func(key K) {
			delete(data, key)
		},
		// Keys returns a slice of all keys in the HashMap.
		//
		// Returns:
		//   []K: A slice of all keys in the HashMap.
		Keys: func() []K {
			keys := make([]K, 0, len(data))
			for k := range data {
				keys = append(keys, k)
			}
			return keys
		},
		// PrintMap prints all key-value pairs in the HashMap.
		PrintMap: func() {
			for k, v := range data {
				fmt.Println(k, v)
			}
		},
	}
}
