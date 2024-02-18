package cerror

import "net/http"

// Code represents Error Code
type Code int

func (c Code) String() string {
	return codeMap[c].message
}

const (
	OK Code = iota
	NotFound
	InvalidArgument
	Forbidden
	Unauthorized
	Internal
	AlreadyExists
	PostgreSQL
	Unknown
	Pagination
	NoRows
	InOpportuneTime
	Mail
	Auth0

	ErrorCodeMax // to validate codeMap size
)

type codeDetail struct {
	message    string
	httpStatus int
}

var codeMap = map[Code]codeDetail{ //nolint: gochecknoglobals
	OK:              {"ok", http.StatusOK},
	NotFound:        {"not_found", http.StatusNotFound},
	InvalidArgument: {"invalid_argument", http.StatusBadRequest},
	Forbidden:       {"forbidden", http.StatusForbidden},
	Unauthorized:    {"unauthorized", http.StatusUnauthorized},
	Internal:        {"internal", http.StatusInternalServerError},
	AlreadyExists:   {"already_exists", http.StatusConflict},
	Unknown:         {"unknown", http.StatusInternalServerError},
	PostgreSQL:      {"postgres", http.StatusInternalServerError}, // データベース系のエラー
	Pagination:      {"pagination", http.StatusBadRequest},        // ページネーション系のエラー
	NoRows:          {"no_rows", http.StatusNotFound},             // データベース系のエラー
	InOpportuneTime: {"in_opportune_time", http.StatusBadRequest}, // 適した時間でない時
	Mail:            {"mail", http.StatusInternalServerError},     // メール送信系のエラー
	Auth0:           {"auth0", http.StatusInternalServerError},    // Auth0系のエラー
}

func mapHTTPErrorToCode(httpStatusCode int) Code {
	switch httpStatusCode {
	case http.StatusNotFound:
		return NotFound
	case http.StatusBadRequest:
		return InvalidArgument
	case http.StatusForbidden:
		return Forbidden
	case http.StatusUnauthorized:
		return Unauthorized
	case http.StatusServiceUnavailable:
		return InOpportuneTime
	case http.StatusInternalServerError:
		return Internal
	default:
		return Unknown
	}
}
