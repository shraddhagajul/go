package twofer

//ShareWith : returns One for X, one for me.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
