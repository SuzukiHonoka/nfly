package api

import (
	"net/http"
	api "nfly/common"
	"nfly/internal/user"
	"nfly/internal/utils"
)

const actionRegister = "register"

func UserRegister(w http.ResponseWriter, r *http.Request) {
	// parse account
	if email, pass, err := utils.ParseAccount(w, r); err == nil {
		// user register
		var u *user.User
		u, err = user.Register(email, pass, nil)
		// if register success
		if err == nil {
			utils.WriteReplyNoCheck(w, http.StatusOK, utils.VtoJson(*api.NewReply(actionRegister, true, u)))
			return
		}
		// if register failed
		utils.WriteReplyNoCheck(w, http.StatusInternalServerError, utils.VtoJson(*api.NewReply(actionRegister, false, err.Error())))
	}
}
