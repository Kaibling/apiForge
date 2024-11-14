package apierror

type MultiError struct {
	errors []error
}

func NewMultiError() *MultiError {
	return &MultiError{}
}

func (m *MultiError) Add(err error) {
	m.errors = append(m.errors, err)
}

func (m *MultiError) HasError() bool {
	for _, e := range m.errors {
		if e != nil {
			return true
		}
	}

	return false
}

func (m *MultiError) GetErrors() []string {
	errString := []string{}

	for _, e := range m.errors {
		if e != nil {
			errString = append(errString, e.Error())
		}
	}

	return errString
}
