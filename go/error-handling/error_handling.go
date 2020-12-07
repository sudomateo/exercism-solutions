package erratum

// Use opens a resource and performs work on it.
func Use(o ResourceOpener, input string) (err error) {

	// Continuously try to open the res.
	var res Resource
	res, err = o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		res, err = o()
	}
	defer res.Close()

	// Recover from a panic.
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				res.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			default:
				err = nil
			}
		}
	}()

	res.Frob(input)

	return err
}
