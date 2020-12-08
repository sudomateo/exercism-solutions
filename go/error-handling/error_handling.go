package erratum

import "errors"

// Use opens a Resource using ResourceOpener o and performs Frob on input. The
// Frob may panic so we try to recover if so.
func Use(o ResourceOpener, input string) (err error) {

	// Continuously try to open the Resource.
	var res Resource
	res, err = o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		res, err = o()
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
				err = errors.New("unable to recover from Frob panic")
			}
		}
	}()

	res.Frob(input)

	return err
}
