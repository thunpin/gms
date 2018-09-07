package filters

import (
	"github.com/thunpin/gms"
	"github.com/thunpin/gms/db/gorm"
)

const TX_KEY = "TX_FILTER_KEY"

var TX = func(context *gms.Context, chain *gms.Chain) (interface{}, error) {
	tx := gorm.DB().Begin()
	context.Params[TX_KEY] = tx
	defer tx.Rollback()

	result, err := chain.Next(context)
	if err == nil {
		tx.Commit()
	}

	return result, err
}
