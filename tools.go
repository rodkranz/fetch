package fetch

// MustString ignore any error and return a string
func MustString(s string, _ error) string {
	return s
}

// MustBytes ignore any error and return a slice of bytes
func MustBytes(b []byte, _ error) []byte {
	if b == nil {
		b = []byte{}
	}

	return b
}
