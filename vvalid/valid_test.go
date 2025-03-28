package vvalid_test

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/sunshinexcode/gotk/vconv"
	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vtest"
	"github.com/sunshinexcode/gotk/vvalid"
)

func TestNew(t *testing.T) {
	err := vvalid.New().Rules("min:18").Data(16).Messages("Age does not meet requirements").Run(context.TODO())

	vtest.Equal(t, "Age does not meet requirements", err.Error())

	err = vvalid.New().Rules("min:18").Data(16).Run(context.TODO())

	vtest.Equal(t, "The value `16` must be equal or greater than 18", err.Error())

	// Struct
	type User struct {
		Name string `v:"required#Please enter user name"`
		Type int    `v:"required#Please select user type"`
	}
	user := User{}
	data := vmap.Map{
		"name": "john",
	}
	err2 := vconv.Struct(data, &user)

	vtest.Nil(t, err2)

	err = vvalid.New().Assoc(data).Data(user).Run(context.TODO())

	vtest.NotNil(t, err)
	vtest.Equal(t, "Please select user type", err.Error())

	// Map
	params := map[string]interface{}{
		"passport":  "",
		"password":  "123456",
		"password2": "1234567",
	}
	rules := map[string]string{
		"passport":  "required|length:6,16",
		"password":  "required|length:6,16|same:password2",
		"password2": "required|length:6,16",
	}
	messages := map[string]interface{}{
		"passport": "Account cannot be empty|Account length should be between {min} and {max}",
		"password": map[string]string{
			"required": "Password cannot be empty",
			"same":     "Two passwords entered are not equal",
		},
	}
	err = vvalid.New().Messages(messages).Rules(rules).Data(params).Run(gctx.New())

	vtest.NotNil(t, err)
	vtest.Equal(t, 2, len(err.Items()))
}
