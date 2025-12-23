package generate

import "errors"

type errorGroup []error

func (errs *errorGroup) toError() error {
	if len(*errs) == 0 {
		return nil
	}

	errMsg := ""

	for _, err := range *errs {
		if err == nil {
			continue
		}
		errMsg += err.Error() + "\n"

	}

	if errMsg == "" {
		return nil
	}

	return errors.New(errMsg)
}

func (errs *errorGroup) Add(err error) {
	*errs = append(*errs, err)
}
