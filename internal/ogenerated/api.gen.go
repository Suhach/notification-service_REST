// Package notification provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// CreateNotificationDTOReq defines model for CreateNotificationDTOReq.
type CreateNotificationDTOReq struct {
	Message string             `json:"message"`
	UserId  openapi_types.UUID `json:"user_id"`
}

// CreateNotificationDTORes defines model for CreateNotificationDTORes.
type CreateNotificationDTORes struct {
	NotificationUuid *openapi_types.UUID `json:"notification_uuid,omitempty"`
}

// Error400 defines model for Error400.
type Error400 struct {
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

// Error401 defines model for Error401.
type Error401 struct {
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

// Error403 defines model for Error403.
type Error403 struct {
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

// Error422 defines model for Error422.
type Error422 struct {
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

// Error500 defines model for Error500.
type Error500 struct {
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

// GetNotificationByUUIDDTORes defines model for GetNotificationByUUIDDTORes.
type GetNotificationByUUIDDTORes struct {
	CreatedAt        *time.Time          `json:"created_at,omitempty"`
	Message          *string             `json:"message,omitempty"`
	NotificationUuid *openapi_types.UUID `json:"notification_uuid,omitempty"`
	Status           *string             `json:"status,omitempty"`
	UserId           *openapi_types.UUID `json:"user_id,omitempty"`
}

// GetNotificationsResponseDTO defines model for GetNotificationsResponseDTO.
type GetNotificationsResponseDTO struct {
	CreatedAt        *time.Time          `json:"created_at,omitempty"`
	Message          *string             `json:"message,omitempty"`
	NotificationUuid *openapi_types.UUID `json:"notification_uuid,omitempty"`
	Status           *string             `json:"status,omitempty"`
	UserId           *openapi_types.UUID `json:"user_id,omitempty"`
}

// UpdateNotificationDTOReq defines model for UpdateNotificationDTOReq.
type UpdateNotificationDTOReq struct {
	Uuid *openapi_types.UUID `json:"uuid,omitempty"`
}

// UpdateNotificationDTORes defines model for UpdateNotificationDTORes.
type UpdateNotificationDTORes struct {
	Status *string `json:"status,omitempty"`
}

// PatchNotificationUuidJSONRequestBody defines body for PatchNotificationUuid for application/json ContentType.
type PatchNotificationUuidJSONRequestBody = UpdateNotificationDTOReq

// PostNotificationsJSONRequestBody defines body for PostNotifications for application/json ContentType.
type PostNotificationsJSONRequestBody = CreateNotificationDTOReq

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get notification
	// (GET /notification/{uuid})
	GetNotificationUuid(c *gin.Context, uuid openapi_types.UUID)
	// Update notification
	// (PATCH /notification/{uuid})
	PatchNotificationUuid(c *gin.Context, uuid openapi_types.UUID)
	// List notifications
	// (GET /notifications)
	GetNotifications(c *gin.Context)
	// Create notification
	// (POST /notifications)
	PostNotifications(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetNotificationUuid operation middleware
func (siw *ServerInterfaceWrapper) GetNotificationUuid(c *gin.Context) {

	var err error

	// ------------- Path parameter "uuid" -------------
	var uuid openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", c.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uuid: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetNotificationUuid(c, uuid)
}

// PatchNotificationUuid operation middleware
func (siw *ServerInterfaceWrapper) PatchNotificationUuid(c *gin.Context) {

	var err error

	// ------------- Path parameter "uuid" -------------
	var uuid openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "uuid", c.Param("uuid"), &uuid, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uuid: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchNotificationUuid(c, uuid)
}

// GetNotifications operation middleware
func (siw *ServerInterfaceWrapper) GetNotifications(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetNotifications(c)
}

// PostNotifications operation middleware
func (siw *ServerInterfaceWrapper) PostNotifications(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostNotifications(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/notification/:uuid", wrapper.GetNotificationUuid)
	router.PATCH(options.BaseURL+"/notification/:uuid", wrapper.PatchNotificationUuid)
	router.GET(options.BaseURL+"/notifications", wrapper.GetNotifications)
	router.POST(options.BaseURL+"/notifications", wrapper.PostNotifications)
}

type GetNotificationUuidRequestObject struct {
	Uuid openapi_types.UUID `json:"uuid"`
}

type GetNotificationUuidResponseObject interface {
	VisitGetNotificationUuidResponse(w http.ResponseWriter) error
}

type GetNotificationUuid200JSONResponse GetNotificationByUUIDDTORes

func (response GetNotificationUuid200JSONResponse) VisitGetNotificationUuidResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetNotificationUuid400JSONResponse Error400

func (response GetNotificationUuid400JSONResponse) VisitGetNotificationUuidResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetNotificationUuid401JSONResponse Error401

func (response GetNotificationUuid401JSONResponse) VisitGetNotificationUuidResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type GetNotificationUuid403JSONResponse Error403

func (response GetNotificationUuid403JSONResponse) VisitGetNotificationUuidResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type GetNotificationUuid422JSONResponse Error422

func (response GetNotificationUuid422JSONResponse) VisitGetNotificationUuidResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(422)

	return json.NewEncoder(w).Encode(response)
}

type GetNotificationUuid500JSONResponse Error500

func (response GetNotificationUuid500JSONResponse) VisitGetNotificationUuidResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PatchNotificationUuidRequestObject struct {
	Uuid openapi_types.UUID `json:"uuid"`
	Body *PatchNotificationUuidJSONRequestBody
}

type PatchNotificationUuidResponseObject interface {
	VisitPatchNotificationUuidResponse(w http.ResponseWriter) error
}

type PatchNotificationUuid200JSONResponse UpdateNotificationDTORes

func (response PatchNotificationUuid200JSONResponse) VisitPatchNotificationUuidResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchNotificationUuid400JSONResponse Error400

func (response PatchNotificationUuid400JSONResponse) VisitPatchNotificationUuidResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PatchNotificationUuid401JSONResponse Error401

func (response PatchNotificationUuid401JSONResponse) VisitPatchNotificationUuidResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type PatchNotificationUuid403JSONResponse Error403

func (response PatchNotificationUuid403JSONResponse) VisitPatchNotificationUuidResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type PatchNotificationUuid422JSONResponse Error422

func (response PatchNotificationUuid422JSONResponse) VisitPatchNotificationUuidResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(422)

	return json.NewEncoder(w).Encode(response)
}

type PatchNotificationUuid500JSONResponse Error500

func (response PatchNotificationUuid500JSONResponse) VisitPatchNotificationUuidResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetNotificationsRequestObject struct {
}

type GetNotificationsResponseObject interface {
	VisitGetNotificationsResponse(w http.ResponseWriter) error
}

type GetNotifications200JSONResponse []GetNotificationsResponseDTO

func (response GetNotifications200JSONResponse) VisitGetNotificationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetNotifications400JSONResponse Error400

func (response GetNotifications400JSONResponse) VisitGetNotificationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetNotifications401JSONResponse Error401

func (response GetNotifications401JSONResponse) VisitGetNotificationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type GetNotifications403JSONResponse Error403

func (response GetNotifications403JSONResponse) VisitGetNotificationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type GetNotifications422JSONResponse Error422

func (response GetNotifications422JSONResponse) VisitGetNotificationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(422)

	return json.NewEncoder(w).Encode(response)
}

type GetNotifications500JSONResponse Error500

func (response GetNotifications500JSONResponse) VisitGetNotificationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type PostNotificationsRequestObject struct {
	Body *PostNotificationsJSONRequestBody
}

type PostNotificationsResponseObject interface {
	VisitPostNotificationsResponse(w http.ResponseWriter) error
}

type PostNotifications201JSONResponse CreateNotificationDTORes

func (response PostNotifications201JSONResponse) VisitPostNotificationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostNotifications400JSONResponse Error400

func (response PostNotifications400JSONResponse) VisitPostNotificationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostNotifications401JSONResponse Error401

func (response PostNotifications401JSONResponse) VisitPostNotificationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(401)

	return json.NewEncoder(w).Encode(response)
}

type PostNotifications403JSONResponse Error403

func (response PostNotifications403JSONResponse) VisitPostNotificationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type PostNotifications422JSONResponse Error422

func (response PostNotifications422JSONResponse) VisitPostNotificationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(422)

	return json.NewEncoder(w).Encode(response)
}

type PostNotifications500JSONResponse Error500

func (response PostNotifications500JSONResponse) VisitPostNotificationsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get notification
	// (GET /notification/{uuid})
	GetNotificationUuid(ctx context.Context, request GetNotificationUuidRequestObject) (GetNotificationUuidResponseObject, error)
	// Update notification
	// (PATCH /notification/{uuid})
	PatchNotificationUuid(ctx context.Context, request PatchNotificationUuidRequestObject) (PatchNotificationUuidResponseObject, error)
	// List notifications
	// (GET /notifications)
	GetNotifications(ctx context.Context, request GetNotificationsRequestObject) (GetNotificationsResponseObject, error)
	// Create notification
	// (POST /notifications)
	PostNotifications(ctx context.Context, request PostNotificationsRequestObject) (PostNotificationsResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetNotificationUuid operation middleware
func (sh *strictHandler) GetNotificationUuid(ctx *gin.Context, uuid openapi_types.UUID) {
	var request GetNotificationUuidRequestObject

	request.Uuid = uuid

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetNotificationUuid(ctx, request.(GetNotificationUuidRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetNotificationUuid")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetNotificationUuidResponseObject); ok {
		if err := validResponse.VisitGetNotificationUuidResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PatchNotificationUuid operation middleware
func (sh *strictHandler) PatchNotificationUuid(ctx *gin.Context, uuid openapi_types.UUID) {
	var request PatchNotificationUuidRequestObject

	request.Uuid = uuid

	var body PatchNotificationUuidJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchNotificationUuid(ctx, request.(PatchNotificationUuidRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchNotificationUuid")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PatchNotificationUuidResponseObject); ok {
		if err := validResponse.VisitPatchNotificationUuidResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetNotifications operation middleware
func (sh *strictHandler) GetNotifications(ctx *gin.Context) {
	var request GetNotificationsRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetNotifications(ctx, request.(GetNotificationsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetNotifications")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetNotificationsResponseObject); ok {
		if err := validResponse.VisitGetNotificationsResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostNotifications operation middleware
func (sh *strictHandler) PostNotifications(ctx *gin.Context) {
	var request PostNotificationsRequestObject

	var body PostNotificationsJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostNotifications(ctx, request.(PostNotificationsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostNotifications")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostNotificationsResponseObject); ok {
		if err := validResponse.VisitPostNotificationsResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
