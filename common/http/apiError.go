package common

type Error struct {
	Code        int    `json:"errCode,omitempty"`
	Type        string `json:"errType,omitempty"`
	ShortError  string `json:"shortError,omitempty"`
	Description string `json:"description,omitempty"`
}

type apiError struct {
	DefaultError        *Error
	InternalServerError *Error
	JsonMarshalError    *Error
	JsonUnmarshalError  *Error
	NoDataFoundError    *Error
	UnauthorizedError   *Error
	ForbiddenError      *Error
	InvalidsessError    *Error
	GetCacheError       *Error
	SetCacheError       *Error
	AddDBError          *Error
	GetDBError          *Error
	DelDBError          *Error
	OracleDBConnError   *Error
	BadRequestError     *Error
}

var ApiErrors = newApiErrorsRegistry()

func newApiErrorsRegistry() *apiError {
	return &apiError{
		DefaultError:        &Error{Code: 1000, Type: "DefaultError", ShortError: "DefaultError", Description: "DefaultError"},
		InternalServerError: &Error{Code: 1001, Type: "InternalServerError", ShortError: "InternalServerError", Description: "InternalServerError"},
		JsonMarshalError:    &Error{Code: 1002, Type: "JsonMarshalError", ShortError: "JsonMarshalError", Description: "JsonMarshalError"},
		JsonUnmarshalError:  &Error{Code: 1003, Type: "JsonUnmarshalError", ShortError: "JsonUnmarshalError", Description: "JsonUnmarshalError"},
		NoDataFoundError:    &Error{Code: 1004, Type: "NoDataFoundError", ShortError: "NoDataFoundError", Description: "NoDataFoundError"},
		UnauthorizedError:   &Error{Code: 1005, Type: "UnauthorizedError", ShortError: "UnauthorizedError", Description: "UnauthorizedError"},
		ForbiddenError:      &Error{Code: 1006, Type: "ForbiddenError", ShortError: "ForbiddenError", Description: "ForbiddenError"},
		InvalidsessError:    &Error{Code: 1007, Type: "InvalidsessError", ShortError: "InvalidsessError", Description: "InvalidsessError"},
		GetCacheError:       &Error{Code: 1008, Type: "GetCacheError", ShortError: "GetCacheError", Description: "GetCacheError"},
		SetCacheError:       &Error{Code: 1009, Type: "SetCacheError", ShortError: "SetCacheError", Description: "SetCacheError"},
		AddDBError:          &Error{Code: 1010, Type: "AddDBError", ShortError: "AddDBError", Description: "AddDBError"},
		GetDBError:          &Error{Code: 1011, Type: "GetDBError", ShortError: "GetDBError", Description: "GetDBError"},
		DelDBError:          &Error{Code: 1012, Type: "DelDBError", ShortError: "DelDBError", Description: "DelDBError"},
		OracleDBConnError:   &Error{Code: 1013, Type: "OracleDBConnError", ShortError: "OracleDBConnError", Description: "OracleDBConnError"},
		BadRequestError:     &Error{Code: 1014, Type: "BadRequestError", ShortError: "BadRequestError", Description: "BadRequestError"},
	}
}

func (e *Error) Error() string {
	if e == nil {
		return "UndefinedError"
	}
	return e.Type + ": " + e.Description
}

func (e *Error) WithErrorDescription(description string) Error {
	var err Error

	if e == nil {
		return *ApiErrors.DefaultError
	} else {
		err = *e
	}

	e.Description = description
	return err
}
