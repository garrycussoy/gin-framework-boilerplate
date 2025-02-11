// This package provides defined custom errors used accross the code
package custom_errors

// ===== Incoming-request-related error =====
func AuthorizationFailed(detail string) CustomError {
	return CreateCustomError(401, "0001", "Authorization failed", detail)
}

func ValidationFailed(detail string) CustomError {
	return CreateCustomError(400, "0002", "Payload validation failed", detail)
}
