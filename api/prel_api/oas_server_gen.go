// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// APIIamRolesGet implements GET /api/iam-roles operation.
	//
	// Return iam roles in project id.
	//
	// GET /api/iam-roles
	APIIamRolesGet(ctx context.Context, params APIIamRolesGetParams) (APIIamRolesGetRes, error)
	// APIInvitationsPost implements POST /api/invitations operation.
	//
	// Create user invitation.
	//
	// POST /api/invitations
	APIInvitationsPost(ctx context.Context, req *APIInvitationsPostReq) (APIInvitationsPostRes, error)
	// APIRequestsGet implements GET /api/requests operation.
	//
	// Return admin request with paging.
	//
	// GET /api/requests
	APIRequestsGet(ctx context.Context, params APIRequestsGetParams) (APIRequestsGetRes, error)
	// APIRequestsPost implements POST /api/requests operation.
	//
	// Post request.
	//
	// POST /api/requests
	APIRequestsPost(ctx context.Context, req *APIRequestsPostReq) (APIRequestsPostRes, error)
	// APIRequestsRequestIDDelete implements DELETE /api/requests/{requestID} operation.
	//
	// Delete request.
	//
	// DELETE /api/requests/{requestID}
	APIRequestsRequestIDDelete(ctx context.Context, params APIRequestsRequestIDDeleteParams) (APIRequestsRequestIDDeleteRes, error)
	// APIRequestsRequestIDPatch implements PATCH /api/requests/{requestID} operation.
	//
	// Update request.
	//
	// PATCH /api/requests/{requestID}
	APIRequestsRequestIDPatch(ctx context.Context, req *APIRequestsRequestIDPatchReq, params APIRequestsRequestIDPatchParams) (APIRequestsRequestIDPatchRes, error)
	// APIUsersGet implements GET /api/users operation.
	//
	// Return admin user with paging.
	//
	// GET /api/users
	APIUsersGet(ctx context.Context, params APIUsersGetParams) (APIUsersGetRes, error)
	// APIUsersUserIDPatch implements PATCH /api/users/{userID} operation.
	//
	// Update user.
	//
	// PATCH /api/users/{userID}
	APIUsersUserIDPatch(ctx context.Context, req *APIUsersUserIDPatchReq, params APIUsersUserIDPatchParams) (APIUsersUserIDPatchRes, error)
	// AdminRequestGet implements GET /admin/request operation.
	//
	// Return admin request page.
	//
	// GET /admin/request
	AdminRequestGet(ctx context.Context) (AdminRequestGetRes, error)
	// AdminUserGet implements GET /admin/user operation.
	//
	// Return admin user page.
	//
	// GET /admin/user
	AdminUserGet(ctx context.Context) (AdminUserGetRes, error)
	// AuthGoogleCallbackGet implements GET /auth/google/callback operation.
	//
	// Google callback endpoint.
	//
	// GET /auth/google/callback
	AuthGoogleCallbackGet(ctx context.Context, params AuthGoogleCallbackGetParams) (AuthGoogleCallbackGetRes, error)
	// Get implements GET / operation.
	//
	// Display top page.
	//
	// GET /
	Get(ctx context.Context, params GetParams) (GetRes, error)
	// HealthGet implements GET /health operation.
	//
	// Healthcheck.
	//
	// GET /health
	HealthGet(ctx context.Context) (HealthGetRes, error)
	// RequestFormGet implements GET /request-form operation.
	//
	// Get request form.
	//
	// GET /request-form
	RequestFormGet(ctx context.Context) (RequestFormGetRes, error)
	// RequestGet implements GET /request operation.
	//
	// Return request list page.
	//
	// GET /request
	RequestGet(ctx context.Context) (RequestGetRes, error)
	// RequestRequestIDGet implements GET /request/{requestID} operation.
	//
	// Get request page.
	//
	// GET /request/{requestID}
	RequestRequestIDGet(ctx context.Context, params RequestRequestIDGetParams) (RequestRequestIDGetRes, error)
	// SigninPost implements POST /signin operation.
	//
	// Sign in.
	//
	// POST /signin
	SigninPost(ctx context.Context) (SigninPostRes, error)
	// SignoutPost implements POST /signout operation.
	//
	// Sign out.
	//
	// POST /signout
	SignoutPost(ctx context.Context) (SignoutPostRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
