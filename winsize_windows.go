package winsize

// Get is not implemented
func Get() (s *Size, err error) {
	return nil, ErrNotImplemented
}

// Set is not implemented
func Set(*Size) (err error) {
	return ErrNotImplemented
}
