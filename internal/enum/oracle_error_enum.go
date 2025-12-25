package enum

const (
	// Common
	MsgOracleUniqueViolation   = "The uniqueness constraint has been violated. Such a record already exists"
	MsgOracleInvalidSQL        = "Invalid SQL query"
	MsgOracleInvalidCreate     = "Invalid CREATE command"
	MsgOracleInvalidDataType   = "Invalid data type"
	MsgOracleInvalidTable      = "Invalid table name"
	MsgOracleInvalidIdentifier = "Invalid field name"
	MsgOracleMisspelledKeyword = "SQL keyword error"
	MsgOracleMissingLeftParen  = "Missing left parenthesis"
	MsgOracleMissingRightParen = "The right parenthesis is missing"

	// Cursor / runtime
	MsgOracleInvalidCursor = "Invalid cursor"
	MsgOracleCanceled      = "The operation was canceled by the user"
	MsgOracleConnLost      = "The database connection was lost"

	// Network / listener
	MsgOracleListenerACL    = "Listener rejected connection (ACL filtering)"
	MsgOracleServiceUnknown = "Listener does not know the specified service"
	MsgOracleNoHandler      = "Listener could not find an available connection handler"
	MsgOracleConnRefused    = "Connection with Oracle rejected"
	MsgOracleUsernameFailed = "Error getting username"

	// Auth
	MsgOracleAuthProtocol = "Authentication protocol error"
)
