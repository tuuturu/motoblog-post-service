package entrypoints

const maxTruncatedLength = 160

func truncateContent(original string) string {
	if len(original) > maxTruncatedLength {
		return original[0:maxTruncatedLength]
	}

	return original
}
