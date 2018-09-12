package exit

import (
	"os"
)

const (
	// Success code
	OK      = 0
	Success = 0
	// Failure
	NotOK                       = 1
	Failure                     = 1
	GenericError                = 1
	CannotExecute               = 126
	CommandInvokedCannotExecute = 126
	CommandNotFound             = 127
	NotFound                    = 127

	// From sysexits.h https://gist.githubusercontent.com/bojanrajkovic/831993/raw/79d07934534ba03d1b21c78917b9a8b699d8d6fe/sysexits.h
	Usage = iota + 64
	DataError
	NoInput
	NoSuchUser
	NoSuchHost
	ServiceUnavailable
	InternalSoftwareError
	OSError
	OSFileMissing
	CannotCreate
	IOError
	TemporaryFailure
	Protocol
	PermissionDenied
	ConfigurationError

	// Alternate names
	AddresseeUnknown = NoSuchUser
	HostNameUnknown  = NoSuchHost
	Unavailable      = ServiceUnavailable
	Software         = InternalSoftwareError
	InternalError    = InternalSoftwareError
	NoPermission     = PermissionDenied
	NoPerm           = NoPermission
	Config           = ConfigurationError
	NoSuchFile		= OSFileMissing
)

type ExitCode uint8

// Fatal adds 128 to a given exit code to signify it as a "fatal" error
func Fatal(code ExitCode) ExitCode {
	return 128 + code
}

// With -- call os.Exit with the given ExitCode
func With(code ExitCode) {
	os.Exit(int(code))
}
