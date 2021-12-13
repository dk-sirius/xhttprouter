package xrouter

import (
	"github.com/dk-sirius/utilx/xtag"
)

func Path(s interface{}) string {
	tags := xtag.
		NewXTag().
		Search(s).
		ByTag(PATH.String())
	if len(tags) > 0 {
		return tags[0].TagValue
	}
	return ""
}
