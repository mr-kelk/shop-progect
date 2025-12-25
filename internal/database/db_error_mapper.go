package database

//TODO: Доделать когда будет свободное время на выходных
import (
	"database/sql"
	"errors"
	"example/shop-progect/internal/enum"

	"github.com/sijms/go-ora/v2/network"
)

var defaultDbErors = "database errors"

func MapDBError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return errors.New(enum.ErrNoRows)
	}

	var oraErr *network.OracleError
	if errors.As(err, &oraErr) {
		return errors.New(OracleErrorMessage(oraErr.ErrCode))
	}

	return errors.New(defaultDbErors)
}
