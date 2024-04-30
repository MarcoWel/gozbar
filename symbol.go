// Package gozbar symbol bindings.
// Read the ZBar documents for details
package gozbar

// #cgo LDFLAGS: -lzbar
// #include <zbar.h>
import "C"

// Symbol is a wrapper around a zbar symbol.
type Symbol struct {
	symbol *C.zbar_symbol_t
}

// Next returns the next symbol or nil if there is none.
func (s *Symbol) Next() *Symbol {
	n := C.zbar_symbol_next(s.symbol)

	if n == nil {
		return nil
	}

	return &Symbol{
		symbol: n,
	}
}

// Data returns the scanned data for this symbol.
func (s *Symbol) Data() string {
	sym := C.zbar_symbol_get_data(s.symbol)

	if sym == nil {
		return ""
	}

	return C.GoString(sym)
}

// Type returns the symbol type.
// Compare it with types in constants to get the accurate symbol type.
func (s *Symbol) Type() C.zbar_symbol_type_t {
	return C.zbar_symbol_get_type(s.symbol)
}

// Each will iterate over all symbols after this symbol.
// passing them into the provided callback
func (s *Symbol) Each(f func(string)) {
	t := s

	for {
		f(t.Data())

		t = t.Next()

		if t == nil {
			break
		}
	}
}

// Type retrieves type of decoded symbol.
func (s *Symbol) Type() int {
	return int(C.zbar_symbol_get_type(s.symbol))
}

// Name retrieves string name for symbol encoding.
func (s *Symbol) Name() string {
	cstr := C.zbar_get_symbol_name(C.zbar_symbol_get_type(s.symbol))
	return C.GoString(cstr)
}

// AddonName retrieves string name for addon encoding
func (s *Symbol) AddonName() string {
	return C.GoString(C.zbar_get_addon_name(C.zbar_symbol_get_type(s.symbol)))
}

// Quality retrieves a symbol confidence metric.
// Returns an unscaled, relative quantity: larger values are better than smaller values,
// where "large" and "small" are application dependent.
// Note:
//
//	expect the exact definition of this quantity to change as the metric is refined.
//	Currently, only the ordered relationship between two values is defined and will remain stable in the future
func (s *Symbol) Quality() int {
	return int(C.zbar_symbol_get_quality(s.symbol))
}

// Retrieve the number of points in the location polygon.
// The location polygon defines the image area that the symbol was extracted from.
// Returns the number of points in the location polygon
// Note:
//
//	this is currently not a polygon, but the scan locations where the symbol was decoded
func (s *Symbol) LocSize() uint {
	return uint(C.zbar_symbol_get_loc_size(s.symbol))
}

// LocX retrieves location polygon x-coordinates.
// Points are specified by 0-based index.
// Returns:
//
//	the x-coordinate for a point in the location polygon.
//	-1 if index is out of range
func (s *Symbol) LocX(index uint) int {
	return int(C.zbar_symbol_get_loc_x(s.symbol, C.unsigned(index)))
}

// LocY retrieves location polygon y-coordinates.
// Points are specified by 0-based index.
// Returns:
//
//	the y-coordinate for a point in the location polygon.
//	-1 if index is out of range
func (s *Symbol) LocY(index uint) int {
	return int(C.zbar_symbol_get_loc_y(s.symbol, C.unsigned(index)))
}
