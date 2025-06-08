package py

import (
	"maps"
	"slices"
)

type Dict[K comparable, V any] map[K]V

func NewDict[K comparable, V any](items ...any) Dict[K, V] {
	d := make(Dict[K, V])

	if len(items) == 0 {
		return d
	}

	// Handle single map argument
	if len(items) == 1 {
		switch v := items[0].(type) {
		case map[K]V:
			maps.Copy(d, v)
			return d
		case [][2]any:
			for _, pair := range v {
				if key, ok := pair[0].(K); ok {
					if val, ok := pair[1].(V); ok {
						d[key] = val
					}
				}
			}
			return d
		}
	}

	for i := 0; i < len(items)-1; i += 2 {
		if key, ok := items[i].(K); ok {
			if val, ok := items[i+1].(V); ok {
				d[key] = val
			}
		}
	}

	return d
}

func (d *Dict[K, V]) Get(key K) V { return (*d)[key] }

func (d *Dict[K, V]) Keys() []K { return slices.Collect(maps.Keys(*d)) }

func (d *Dict[K, V]) Values() []V { return slices.Collect(maps.Values(*d)) }

func (d *Dict[K, V]) Items() [][2]any {
	result := make([][2]any, 0, len(*d))
	for k, v := range *d {
		result = append(result, [2]any{k, v})
	}
	return result
}

func (d *Dict[K, V]) Copy() Dict[K, V] { return maps.Clone(*d) }

func (d *Dict[K, V]) Pop(key K, defaultValue ...V) V {
	if val, exists := (*d)[key]; exists {
		delete(*d, key)
		return val
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	var zero V
	return zero
}

func (d *Dict[K, V]) PopItem() ([2]any, bool) {
	for k, v := range *d {
		delete(*d, k)
		return [2]any{k, v}, true
	}
	return [2]any{}, false
}

func (d *Dict[K, V]) SetDefault(key K, defaultValue V) V {
	if value, exists := (*d)[key]; exists {
		return value
	}
	(*d)[key] = defaultValue
	return defaultValue
}

func (d *Dict[K, V]) Update(other Dict[K, V]) { maps.Copy(*d, other) }

func (d *Dict[K, V]) Clear() { clear(*d) }

func (d Dict[K, V]) ToSlice() [][2]any { return d.Items() }
