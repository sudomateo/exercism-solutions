package erratum

// Use opens a resource and performs work on it.
func Use(o ResourceOpener, input string) (ret error) {

	// Continuously try to open the resource.
	var resource Resource
	for {
		r, err := o()
		if err != nil {
			if _, ok := err.(TransientError); ok {
				continue
			}
			return err
		}
		resource = r
		break
	}
	defer resource.Close()

	// Recover from a panic.
	defer func() {
		if r := recover(); r != nil {
			if ferr, ok := r.(FrobError); ok {
				resource.Defrob(ferr.defrobTag)
				ret = ferr
				return
			}
			ret = r.(error)
			return
		}
	}()

	resource.Frob(input)

	return ret
}
