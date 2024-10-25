package utils

import (
	"errors"
	"log"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestFile struct {
	File    multipart.File
	Headers *multipart.FileHeader
}

type RequestWrapper struct {
	C *gin.Context

	ID uint // ID from the Path

	Page    int // ?page=#
	PerPage int // ?perPage=#

	DefaultPage    int
	DefaultPerPage int

	// ?q= query will be here
	Q string

	Files []RequestFile
}

// Returns the RequestWrapper for some common "?q=" type query params
//
// If there is a cached data, `CreateRequestWrapper()` middleware sends the cache
// and the request doesn't reach controllers.
func GetRequestWrapper(ctx *gin.Context) *RequestWrapper {
	val := ctx.Value("request-wrapper")

	if u, ok := val.(RequestWrapper); ok {
		return &u
	}

	log.Printf("[ERROR] request wrapper can not be parsed!")
	return nil
}

func (r *RequestWrapper) GetOffset() int {
	return r.Page * r.PerPage
}

func (r *RequestWrapper) GetIntQuery(name string, dflt int) int {
	return AtoiDefault(r.C.Query(name), dflt)
}

func (r *RequestWrapper) GetParam(name string) string {
	return r.C.Param(name)
}

func (r *RequestWrapper) GetBoolQuery(name string) bool {
	val, err := strconv.ParseBool(r.C.Query(name))
	if err != nil {
		return false
	}
	return val
}

func (r *RequestWrapper) ParseDefaultQueryParams() {
	r.Page = AtoiDefault(r.C.Query("page"), r.DefaultPage)
	r.PerPage = AtoiDefault(r.C.Query("perPage"), r.DefaultPerPage)

	replacer := strings.NewReplacer("<", "", ">", "", "#", "", "\"", "")
	r.Q = replacer.Replace(strings.ToLower(strings.TrimSpace(r.C.Query("q"))))
}

func (r *RequestWrapper) ParseDefaultPathParams() {
	r.ID = uint(AtoiDefault(r.C.Param("id"), 0))
}

func (r *RequestWrapper) GetSliceQuery(name, delim string) []string {
	return strings.Split(r.C.Query(name), delim)
}

func (r *RequestWrapper) GetPathParam(name string) string {
	return r.C.Param(name)
}

func (r *RequestWrapper) GetPathParamInt(name string) (int, error) {
	plainValue := r.C.Param(name)
	if len(plainValue) == 0 {
		return 0, errors.New("empty parameter received")
	}
	return strconv.Atoi(plainValue)
}

func (r *RequestWrapper) GetQuery(name string) string {
	return r.C.Query(name)
}

func (r *RequestWrapper) ParseFormFile(fileNames ...string) error {
	var files []RequestFile

	for _, name := range fileNames {
		r.C.Request.ParseMultipartForm(10 << 20) // max 10MB

		file, headers, err := r.C.Request.FormFile(name)
		if err != nil {
			return err
		}

		files = append(files, RequestFile{
			File:    file,
			Headers: headers,
		})
	}
	if len(fileNames) != len(files) {
		return errors.New("failed to read all files")
	}
	r.Files = files
	return nil
}
