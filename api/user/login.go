package api

import (
	api "common_notify_server/common"
	"common_notify_server/internal/errors"
	"common_notify_server/internal/session"
	"common_notify_server/internal/user"
	"common_notify_server/internal/utils"
	"net/http"
)

const actionLogin = "login"

func UserLogin(w http.ResponseWriter, r *http.Request) {
	// parse account
	if email, pass, err := utils.ParseAccount(w, r); err == nil {
		// user login
		var u *user.User
		u, err = user.Login(email, pass)
		// if login success
		if err == nil {
			// alloc session
			if s := session.NewSession(utils.ParseIP(r), u); s != nil {
				// set header session value
				http.SetCookie(w, &http.Cookie{
					Name:    "session",
					Value:   s.UUID.String(),
					Expires: s.ExpDate,
				})
				utils.WriteReplyNoCheck(w, http.StatusOK, utils.VtoJson(*api.NewReply(actionLogin, true, u)))
				return
			}
			// if alloc failed
			utils.WriteReplyNoCheck(w, http.StatusLocked, utils.VtoJson(*api.NewReply(actionLogin, false, errors.SessionPoolMaxReached.Error())))
			return
		}
		// if login failed
		utils.WriteReplyNoCheck(w, http.StatusUnauthorized, utils.VtoJson(*api.NewReply(actionLogin, false, err.Error())))
	}
}
