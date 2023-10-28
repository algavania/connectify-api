package pkg

import (
	"errors"
	"example/connectify/app/constant"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func PanicException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseKey constant.ResponseStatus) {
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		key := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		switch key {
		case
			constant.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse_(key, msg, Null()))
			c.Abort()
		case
			constant.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, BuildResponse_(key, msg, Null()))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, BuildResponse_(key, msg, Null()))
			c.Abort()
		}
	}
}

func ConvertConstraintName(str string) string {
	split := strings.Split(str, "_")
	name := split[len(split)-1]
	caser := cases.Title(language.English)
	titleStr := caser.String(name)
	return titleStr
}

func HandleError(error *pgconn.PgError, c *gin.Context) bool {
	if error.Code == "23505" {
		c.JSON(http.StatusBadRequest, BuildResponse_("DUPLICATED_FIELD", ConvertConstraintName(error.ConstraintName)+" already exist", Null()))
		return true
	} else {
		PanicException(constant.UnknownError)
		return false
	}
}
