package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AtoiDefault(s string, def int) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return v
}

func AtoiUint64Default(s string, def uint64) uint64 {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return def
	}
	return v
}

func AtoiFloat64Default(s string, def float64) float64 {
	if s == "" {
		return def
	}
	v, _ := strconv.ParseFloat(strings.TrimSpace(s), 64)
	return v
}

func AtoiFloat64(s string) (float64, error) {
	if s == "" {
		return 0, errors.New("empty string received")
	}
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}

var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)

func ClearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func ParsePerPage(ctx *gin.Context) int {
	perPage := AtoiFloat64Default(ctx.Query("perPage"), 10)
	return int(perPage)
}

func ParsePage(ctx *gin.Context) int {
	page := AtoiFloat64Default(ctx.Query("page"), 0)
	return int(page)
}

func Offset(page, perPage int) int {
	return page * perPage
}
