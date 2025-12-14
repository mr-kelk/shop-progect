package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	// Формирование строки подключения
	//connStr := go_ora.BuildUrl(
	//	config.Cfg.OracleHost,
	//	config.Cfg.OraclePort,
	//	config.Cfg.OracleService,
	//	config.Cfg.OracleUser,
	//	config.Cfg.OraclePass,
	//	nil,
	//)

	// Открываем соединение
	//db, err := sql.Open("oracle", connStr)
	//if err != nil {
	//	return c.String(http.StatusInternalServerError, "DB connection error: "+err.Error())
	//}
	//defer db.Close()
	//
	//// Проверяем
	//if err = db.Ping(); err != nil {
	//	return c.String(http.StatusInternalServerError, "Oracle ping error: "+err.Error())
	//}
	//
	//// выполняем запрос
	//rows, err := db.Query("SELECT ID, NAME FROM SHOP.ROLES")
	//if err != nil {
	//	return c.String(http.StatusInternalServerError, "Query error: "+err.Error())
	//}
	//defer rows.Close()
	//
	//// читаем результаты
	//var roles []Role
	//
	//for rows.Next() {
	//	var r Role
	//	if err := rows.Scan(&r.ID, &r.Name); err != nil {
	//		return c.String(http.StatusInternalServerError, "Scan error: "+err.Error())
	//	}
	//	roles = append(roles, r)
	//}
	//
	//// проверяем ошибку чтения
	//if err = rows.Err(); err != nil {
	//	return c.String(http.StatusInternalServerError, "Rows error: "+err.Error())
	//}

	return c.JSON(http.StatusOK, "product list")
}

func GetProductByUUID(c echo.Context) error {
	uuid := c.Param("uuid")
	return c.String(http.StatusOK, "product "+uuid)
}
