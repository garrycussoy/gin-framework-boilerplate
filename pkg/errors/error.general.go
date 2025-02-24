// This package provides defined custom errors used accross the code
package custom_errors

import "fmt"

// ===== Incoming-request-related error =====
func AuthorizationFailed(detail string) CustomError {
	return CreateCustomError(401, "0001", "Authorization failed", detail)
}

func ValidationFailed(detail string) CustomError {
	return CreateCustomError(400, "0002", "Payload validation failed", detail)
}

func TimeoutLimitExceeded() CustomError {
	return CreateCustomError(408, "0003", "Timeout limit exceeded", nil)
}

// ===== Domain-related error =====
// Auth domain
func AuthDomainError(statusCode int, detail string) CustomError {
	return CreateCustomError(statusCode, "0101", "Auth domain error", detail)
}

// ===== Repository-related error =====
// User repository
func UserRepositoryError(statusCode int, detail string) CustomError {
	return CreateCustomError(statusCode, "0201", "User repository error", detail)
}

// ===== Client-related error =====
func ESBClientError(serviceName, detail string) CustomError {
	return CreateCustomError(500, "0301", "ESB client error", fmt.Sprintf("Service : %s. %s", serviceName, detail))
}
