package cli

import "fmt"

type ExitError struct {
	Code int
	Err  error
}

func (e ExitError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("exit %d", e.Code)
	}
	return e.Err.Error()
}

func usageError(err error) ExitError {
	return ExitError{Code: 1, Err: err}
}

func notFoundError(err error) ExitError {
	return ExitError{Code: 2, Err: err}
}

func internalError(err error) ExitError {
	return ExitError{Code: 3, Err: err}
}
