// Code generated by "stringer -type Format"; DO NOT EDIT.

package edgedb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Binary-98]
	_ = x[JSON-106]
	_ = x[JSONElements-74]
	_ = x[Null-110]
}

const (
	_Format_name_0 = "JSONElements"
	_Format_name_1 = "Binary"
	_Format_name_2 = "JSON"
	_Format_name_3 = "Null"
)

func (i Format) String() string {
	switch {
	case i == 74:
		return _Format_name_0
	case i == 98:
		return _Format_name_1
	case i == 106:
		return _Format_name_2
	case i == 110:
		return _Format_name_3
	default:
		return "Format(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}