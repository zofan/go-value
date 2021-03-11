package value

func StringChunks(s string, size int) (r []string) {
	for i := 0; i < len(s); i += size {
		r = append(r, s[i:i+size])
	}

	return r
}
