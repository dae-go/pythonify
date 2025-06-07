package py

import (
	"maps"
	"slices"
)

type Dict map[any]any

func (d *Dict) Get(key any) any { return (*d)[key] }

func (d *Dict) Keys() []any { return slices.Collect(maps.Keys(*d)) }

func (d *Dict) Values() []any { return slices.Collect(maps.Values(*d)) }

func (d *Dict) Items() [][2]any {
	result := make([][2]any, 0, len(*d))
	for k, v := range *d {
		result = append(result, [2]any{k, v})
	}
	return result
}

func (d *Dict) Copy() Dict { return maps.Clone(*d) }

func (d *Dict) Pop(key any, defaultValue ...any) any {
	var dv any

	if len(defaultValue) > 0 {
		dv = defaultValue[0]
	}
	v, ok := (*d)[key]

	if ok {
		return v
	}
	return dv
}

func (d *Dict) PopItem() ([2]any, bool) {
	for k, v := range *d {
		delete(*d, k)
		return [2]any{k, v}, true
	}
	return [2]any{}, false
}

func (d *Dict) SetDefault(key any, defaultValue any) any {
	if value, exists := (*d)[key]; exists {
		return value
	}
	(*d)[key] = defaultValue
	return defaultValue
}

func (d *Dict) Update(other Dict) { maps.Copy(*d, other) }

func (d *Dict) Clear() { clear(*d) }

func (d Dict) ToSlice() [][2]any { return d.Items() }
