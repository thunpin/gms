package filters

import (
	"net/http"

	"github.com/thunpin/gerrors"
	"github.com/thunpin/gms"
)

type APIResult struct {
	Code   int
	Result interface{}
	Errors []gerrors.Error
}

var ToJSON = func(context *gms.Context, chain *gms.Chain) (interface{}, error) {
	result, err := chain.Next(context)
	apiResult := buildAPIResult(result, err)
	context.JSON(apiResult.Code, apiResult)
	return result, err
}

func buildAPIResult(result interface{}, err error) *APIResult {
	errs := gerrors.New(err).(gerrors.Errors)
	if hasErrors(errs) {
		return buildAPIResultError(errs)
	} else {
		return &APIResult{http.StatusOK, result, nil}
	}
}

func hasErrors(errs gerrors.Errors) bool {
	return errs != nil && len(errs) > 0
}

func buildAPIResultError(errs gerrors.Errors) *APIResult {
	newErrs := make([]gerrors.Error, 0)
	var code int
	for _, err := range errs {
		var currCode int
		var currError gerrors.Error

		switch err.(type) {
		case gerrors.Error:
			currCode, currError = buildError(err.(gerrors.Error))
		default:
			currCode = http.StatusInternalServerError
			currError = gerrors.InternalServerError(nil)
		}

		code = chooseCorrectCode(code, currCode)
		newErrs = append(newErrs, currError)
	}
	return &APIResult{code, nil, newErrs}
}

func buildError(err gerrors.Error) (int, gerrors.Error) {
	if err.Code == http.StatusInternalServerError {
		return err.Code, gerrors.InternalServerError(nil)
	} else {
		return err.Code, err
	}
}

func chooseCorrectCode(oldCode int, newCode int) int {
	if newCode == http.StatusUnauthorized || oldCode == http.StatusUnauthorized {
		return http.StatusUnauthorized
	} else if newCode == http.StatusForbidden || oldCode == http.StatusForbidden {
		return http.StatusForbidden
	} else if newCode > oldCode {
		return newCode
	} else {
		return oldCode
	}
}
