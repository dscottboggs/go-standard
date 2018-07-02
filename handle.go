package standard

import "log"

// Handle (errors) Check a series of error objects for existence and log them.
// This is designed to provide a central place to override error logging and
// behavior
func Handle(e ...error) {
	for _, err := range e {
		if err != nil {
			log.Fatal(err)
		}
	}
}
