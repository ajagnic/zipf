package check

import "log"

func Fatal(err error) {
	if err != nil {
		log.Fatalf("\nFatalError: %v", err)
	}
}
