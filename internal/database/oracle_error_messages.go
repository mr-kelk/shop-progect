package database

import (
	"example/shop-progect/internal/enum"
	"fmt"
)

var oracleErrorMessages = map[int]string{
	1:   enum.MsgOracleUniqueViolation,
	900: enum.MsgOracleInvalidSQL,
	901: enum.MsgOracleInvalidCreate,
	902: enum.MsgOracleInvalidDataType,
	903: enum.MsgOracleInvalidTable,
	904: enum.MsgOracleInvalidIdentifier,
	905: enum.MsgOracleMisspelledKeyword,
	906: enum.MsgOracleMissingLeftParen,
	907: enum.MsgOracleMissingRightParen,

	1001: enum.MsgOracleInvalidCursor,
	1013: enum.MsgOracleCanceled,

	3135: enum.MsgOracleConnLost,

	12506: enum.MsgOracleListenerACL,
	12514: enum.MsgOracleServiceUnknown,
	12516: enum.MsgOracleNoHandler,
	12564: enum.MsgOracleConnRefused,

	12631: enum.MsgOracleUsernameFailed,

	28041: enum.MsgOracleAuthProtocol,
}

func OracleErrorMessage(code int) string {
	if msg, ok := oracleErrorMessages[code]; ok {
		return msg
	}
	return fmt.Sprintf("Ошибка базы данных Oracle (ORA-%05d)", code)
}
