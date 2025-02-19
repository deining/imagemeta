// Code generated by "stringer -type=TagSource"; DO NOT EDIT.

package imagemeta

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EXIF-1]
	_ = x[IPTC-2]
	_ = x[XMP-4]
}

const (
	_TagSource_name_0 = "EXIFIPTC"
	_TagSource_name_1 = "XMP"
)

var (
	_TagSource_index_0 = [...]uint8{0, 4, 8}
)

func (i TagSource) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _TagSource_name_0[_TagSource_index_0[i]:_TagSource_index_0[i+1]]
	case i == 4:
		return _TagSource_name_1
	default:
		return "TagSource(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
