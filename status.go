package splunk

const (
	StatusRequestCompletedSuccessfully               = 200
	StatusCreateRequestCompletedSuccessfully         = 201
	StatusRequestError                               = 400
	StatusAuthenticationFailure                      = 401
	StatusSplunkEnterpriseLicenseDisablesThisFeature = 402
	StatusInsufficientPermission                     = 403
	StatusRequestedEndpointDoesNotExist              = 404
	StatusInvalidOperationForThisEndpoint            = 409
	StatusUnspecifiedInternalServerError             = 500
	StatusFeatureIsDisabledInConfigurationFile       = 503
)
