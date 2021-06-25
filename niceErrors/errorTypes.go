package niceErrors

//Json
const JsonConvError ErrorType = "JsonConvError"

//SQL
const ConnectionError ErrorType = "ConnectionError"
const SqlError ErrorType = "SqlError"
const UnexpectedResultError ErrorType = "UnexpectedResultError"
const InvalidSqlParameterError ErrorType = "InvalidSqlParameterError"

//Configuration getter
const ConfigurationError ErrorType = "ConfigurationError"

//HTTP
const InvalidActionByUserError ErrorType = "InvalidActionByUserError"
const InvalidHeaderError ErrorType = "InvalidHeaderError"
const InvalidAuthError ErrorType = "InvalidAuthError"
const NoResponseError ErrorType = "NoResponseError"

//File I/O
const FileIOError ErrorType = "FileIOError"

//
