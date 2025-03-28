package vstruct

import (
	"github.com/jinzhu/copier"

	"github.com/sunshinexcode/gotk/vcode"
	"github.com/sunshinexcode/gotk/verror"
)

func Copy(toValue interface{}, fromValue interface{}) (err error) {
	if err = copier.Copy(toValue, fromValue); err != nil {
		return verror.Wrap(err, vcode.CodeErrDataCopy, toValue, fromValue)
	}

	return
}

func CopyWithOption(toValue interface{}, fromValue interface{}, opt Option) (err error) {
	if err = copier.CopyWithOption(toValue, fromValue, opt); err != nil {
		return verror.Wrap(err, vcode.CodeErrDataCopy, toValue, fromValue, opt)
	}

	return
}
