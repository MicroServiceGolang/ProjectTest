package status

var (
	NoError                       = 0
	ErrorCodeValidation           = -10
	ErrorCodeValidationDateFormat = -12

	ErrorCodeDB             = -30
	ErrorCodeEntityNotFound = -31

	ErrorCodeRemoteDBO               = -40
	ErrorCodeRemoteDBOEntityNotFound = -41

	ErrorCodeRemoteESB               = -45
	ErrorCodeRemoteESBEntityNotFound = -46

	ErrorCodeRemoteCRM               = -50
	ErrorCodeRemoteCRMEntityNotFound = -51

	ErrorCodeRemoteOther               = -55
	ErrorCodeRemoteOtherEntityNotFound = -56

	ErrorCodeRMQ = -60
)

var (
	//Success - ...
	Success = "Success"
	//Failure - ...
	Failure = "Failure"
)
