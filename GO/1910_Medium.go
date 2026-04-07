func removeOccurrences(s string, part string) string {
    for {
			before, after, found := strings.Cut(s, part)
			s = before + after
			if !found {
				return s
			}

		}
}