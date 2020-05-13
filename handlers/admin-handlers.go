package handlers

import (
	"net/http"
)

// adminAPIHandlers provides HTTP handlers for admin API.
type AdminAPIHandlers struct {
}

// VersionHandler - GET /admin/version
// -----------
// Returns Administration API version
func (a AdminAPIHandlers) VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("v1"))
}

// ServiceInfoHandler - GET /admin/v1/service
// ----------
// Returns server version and uptime.
func (a AdminAPIHandlers) ServiceInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}


// ServiceStopOrRestartHandler - POST /admin/v1/service
// ----------
// Returns server version and uptime.
func (a AdminAPIHandlers) ServiceStopOrRestartHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// GetConfigHandler - GET /admin/v1/config
// Get config.json of this os setup.
func (a AdminAPIHandlers) GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	//config, err := readServerConfig(ctx, objectAPI)
}

// GetConfigHandler - GET /admin/v1/config
// Get config.json of this os setup.
func (a AdminAPIHandlers) SetConfigHandler(w http.ResponseWriter, r *http.Request) {

}

// AddUser - PUT /admin/v1/add-user?accessKey=<access_key>
func (a AdminAPIHandlers) AddUser(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//accessKey := vars["accessKey"]

	//var uinfo models.UserInfo

	//if err = json.Unmarshal(configBytes, &uinfo); err != nil {
	//	logger.LogIf(ctx, err)
	//	writeErrorResponseJSON(ctx, w, errorCodes.ToAPIErr(ErrAdminConfigBadJSON), r.URL)
	//	return
	//}
	//
	//if err = globalIAMSys.SetUser(accessKey, uinfo); err != nil {
	//	writeErrorResponseJSON(ctx, w, toAdminAPIErr(ctx, err), r.URL)
	//	return
	//}
	//
	//// Notify all other Minio peers to reload user
	//for _, nerr := range globalNotificationSys.LoadUser(accessKey, false) {
	//	if nerr.Err != nil {
	//		logger.GetReqInfo(ctx).SetTags("peerAddress", nerr.Host.String())
	//		logger.LogIf(ctx, nerr.Err)
	//	}
	//}
}

// AddUser - GET /admin/v1/add-user?accessKey=<access_key>
func (a AdminAPIHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//accessKey := vars["accessKey"]

	//var uinfo models.UserInfo
}

// RemoveUser - GET /admin/v1/remove-user?accessKey=<access_key>
func (a AdminAPIHandlers) RemoveUser(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//accessKey := vars["accessKey"]

	//var uinfo models.UserInfo
}

// ListUser - GET /admin/v1/list-users
func (a AdminAPIHandlers) ListUsers(w http.ResponseWriter, r *http.Request) {

}

