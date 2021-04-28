package erratum

import "fmt"

// Use opens a Resource using ResourceOpener o and performs Frob on input. The
// Frob may panic so we try to recover if so.
func Use(o ResourceOpener, input string) (err error) {

	// Continuously try to open the Resource.
	var res Resource
	res, err = o()
	for err != nil {
		switch err.(type) {
		case TransientError:
			res, err = o()
		default:
			return err
		}
	}
	defer res.Close()

	// Recover from a panic from Frob.
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				res.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			default:
				err = fmt.Errorf("unable to recover from Frob panic: %v", r)
			}
		}
	}()

	res.Frob(input)

	return err
}
