// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// Cluster defines model for Cluster.
type Cluster struct {
	// Embedded struct due to allOf(#/components/schemas/ClusterId)
	ClusterId
	// Embedded struct due to allOf(#/components/schemas/ClusterProperties)
	ClusterProperties
}

// ClusterFacts defines model for ClusterFacts.
type ClusterFacts map[string]interface{}

// ClusterId defines model for ClusterId.
type ClusterId struct {

	// A unique object identifier string. Automatically generated by the API on creation.
	Id Id `json:"id"`
}

// ClusterProperties defines model for ClusterProperties.
type ClusterProperties struct {

	// Display Name of the cluster
	DisplayName *string `json:"displayName,omitempty"`

	// Facts about a cluster object. Statically configured key/value pairs.
	Facts *ClusterFacts `json:"facts,omitempty"`

	// Configuration Git repository, usually generated by the API
	GitRepo *GitRepo `json:"gitRepo,omitempty"`

	// Id of the tenant this cluster belongs to
	Tenant string `json:"tenant"`
}

// GitRepo defines model for GitRepo.
type GitRepo struct {

	// SSH public key / deploy key for clusterconfiguration catalog Git repository. This property is managed by Steward.
	DeployKey *string `json:"deployKey,omitempty"`

	// SSH known hosts of the git server (multiline possible for multiple keys)
	HostKeys *string `json:"hostKeys,omitempty"`

	// Full URL of the git repo
	Url *string `json:"url,omitempty"`
}

// Id defines model for Id.
type Id string

// Inventory defines model for Inventory.
type Inventory struct {
	Cluster   string                  `json:"cluster"`
	Inventory *map[string]interface{} `json:"inventory,omitempty"`
}

// Reason defines model for Reason.
type Reason struct {

	// The reason message
	Reason string `json:"reason"`
}

// Tenant defines model for Tenant.
type Tenant struct {
	// Embedded struct due to allOf(#/components/schemas/TenantId)
	TenantId
	// Embedded struct due to allOf(#/components/schemas/TenantProperties)
	TenantProperties
}

// TenantId defines model for TenantId.
type TenantId struct {

	// A unique object identifier string. Automatically generated by the API on creation.
	Id Id `json:"id"`
}

// TenantProperties defines model for TenantProperties.
type TenantProperties struct {

	// Display name of the tenant
	DisplayName *string `json:"displayName,omitempty"`

	// Configuration Git repository, usually generated by the API
	GitRepo *GitRepo `json:"gitRepo,omitempty"`

	// The tenant this tenant belongs to
	Tenant *string `json:"tenant,omitempty"`
}

// ClusterIdParameter defines model for ClusterIdParameter.
type ClusterIdParameter Id

// TenantIdParameter defines model for TenantIdParameter.
type TenantIdParameter Id

// Default defines model for Default.
type Default Reason

// ListClustersParams defines parameters for ListClusters.
type ListClustersParams struct {

	// Filter clusters by tenant id
	Tenant *string `json:"tenant,omitempty"`
}

// CreateClusterJSONBody defines parameters for CreateCluster.
type CreateClusterJSONBody ClusterProperties

// InstallStewardParams defines parameters for InstallSteward.
type InstallStewardParams struct {

	// Initial bootstrap token
	Token *string `json:"token,omitempty"`
}

// QueryInventoryParams defines parameters for QueryInventory.
type QueryInventoryParams struct {

	// InfluxQL query string
	Q *string `json:"q,omitempty"`
}

// UpdateInventoryJSONBody defines parameters for UpdateInventory.
type UpdateInventoryJSONBody Inventory

// CreateTenantJSONBody defines parameters for CreateTenant.
type CreateTenantJSONBody TenantProperties

// CreateClusterRequestBody defines body for CreateCluster for application/json ContentType.
type CreateClusterJSONRequestBody CreateClusterJSONBody

// UpdateInventoryRequestBody defines body for UpdateInventory for application/json ContentType.
type UpdateInventoryJSONRequestBody UpdateInventoryJSONBody

// CreateTenantRequestBody defines body for CreateTenant for application/json ContentType.
type CreateTenantJSONRequestBody CreateTenantJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns a list of clusters.
	// (GET /clusters)
	ListClusters(ctx echo.Context, params ListClustersParams) error
	// Creates a new cluster
	// (POST /clusters)
	CreateCluster(ctx echo.Context) error
	// Deletes a cluster
	// (DELETE /clusters/{clusterId})
	DeleteCluster(ctx echo.Context, clusterId ClusterIdParameter) error
	// Returns all values of a cluster
	// (GET /clusters/{clusterId})
	GetCluster(ctx echo.Context, clusterId ClusterIdParameter) error
	// Updates a cluster
	// (PATCH /clusters/{clusterId})
	UpdateCluster(ctx echo.Context, clusterId ClusterIdParameter) error
	// API health check
	// (GET /healthz)
	Healthz(ctx echo.Context) error
	// Returns the Steward JSON installation manifest
	// (GET /install/steward.json)
	InstallSteward(ctx echo.Context, params InstallStewardParams) error
	// Returns inventory data according to query
	// (GET /inventory)
	QueryInventory(ctx echo.Context, params QueryInventoryParams) error
	// Write inventory data
	// (POST /inventory)
	UpdateInventory(ctx echo.Context) error
	// Returns a list of tenants.
	// (GET /tenants)
	ListTenants(ctx echo.Context) error
	// Creates a new tenant
	// (POST /tenants)
	CreateTenant(ctx echo.Context) error
	// Deletes a tenant
	// (DELETE /tenants/{tenantId})
	DeleteTenant(ctx echo.Context, tenantId TenantIdParameter) error
	// Returns all values of a tenant
	// (GET /tenants/{tenantId})
	GetTenant(ctx echo.Context, tenantId TenantIdParameter) error
	// Updates a tenant
	// (PATCH /tenants/{tenantId})
	UpdateTenant(ctx echo.Context, tenantId TenantIdParameter) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListClusters converts echo context to params.
func (w *ServerInterfaceWrapper) ListClusters(ctx echo.Context) error {
	var err error

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListClustersParams
	// ------------- Optional query parameter "tenant" -------------
	if paramValue := ctx.QueryParam("tenant"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "tenant", ctx.QueryParams(), &params.Tenant)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenant: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListClusters(ctx, params)
	return err
}

// CreateCluster converts echo context to params.
func (w *ServerInterfaceWrapper) CreateCluster(ctx echo.Context) error {
	var err error

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateCluster(ctx)
	return err
}

// DeleteCluster converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteCluster(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clusterId" -------------
	var clusterId ClusterIdParameter

	err = runtime.BindStyledParameter("simple", false, "clusterId", ctx.Param("clusterId"), &clusterId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clusterId: %s", err))
	}

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteCluster(ctx, clusterId)
	return err
}

// GetCluster converts echo context to params.
func (w *ServerInterfaceWrapper) GetCluster(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clusterId" -------------
	var clusterId ClusterIdParameter

	err = runtime.BindStyledParameter("simple", false, "clusterId", ctx.Param("clusterId"), &clusterId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clusterId: %s", err))
	}

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCluster(ctx, clusterId)
	return err
}

// UpdateCluster converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateCluster(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "clusterId" -------------
	var clusterId ClusterIdParameter

	err = runtime.BindStyledParameter("simple", false, "clusterId", ctx.Param("clusterId"), &clusterId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter clusterId: %s", err))
	}

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateCluster(ctx, clusterId)
	return err
}

// Healthz converts echo context to params.
func (w *ServerInterfaceWrapper) Healthz(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Healthz(ctx)
	return err
}

// InstallSteward converts echo context to params.
func (w *ServerInterfaceWrapper) InstallSteward(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params InstallStewardParams
	// ------------- Optional query parameter "token" -------------
	if paramValue := ctx.QueryParam("token"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "token", ctx.QueryParams(), &params.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter token: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.InstallSteward(ctx, params)
	return err
}

// QueryInventory converts echo context to params.
func (w *ServerInterfaceWrapper) QueryInventory(ctx echo.Context) error {
	var err error

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params QueryInventoryParams
	// ------------- Optional query parameter "q" -------------
	if paramValue := ctx.QueryParam("q"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "q", ctx.QueryParams(), &params.Q)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter q: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.QueryInventory(ctx, params)
	return err
}

// UpdateInventory converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateInventory(ctx echo.Context) error {
	var err error

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateInventory(ctx)
	return err
}

// ListTenants converts echo context to params.
func (w *ServerInterfaceWrapper) ListTenants(ctx echo.Context) error {
	var err error

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListTenants(ctx)
	return err
}

// CreateTenant converts echo context to params.
func (w *ServerInterfaceWrapper) CreateTenant(ctx echo.Context) error {
	var err error

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateTenant(ctx)
	return err
}

// DeleteTenant converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteTenant(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenantId" -------------
	var tenantId TenantIdParameter

	err = runtime.BindStyledParameter("simple", false, "tenantId", ctx.Param("tenantId"), &tenantId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenantId: %s", err))
	}

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteTenant(ctx, tenantId)
	return err
}

// GetTenant converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenant(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenantId" -------------
	var tenantId TenantIdParameter

	err = runtime.BindStyledParameter("simple", false, "tenantId", ctx.Param("tenantId"), &tenantId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenantId: %s", err))
	}

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTenant(ctx, tenantId)
	return err
}

// UpdateTenant converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTenant(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenantId" -------------
	var tenantId TenantIdParameter

	err = runtime.BindStyledParameter("simple", false, "tenantId", ctx.Param("tenantId"), &tenantId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenantId: %s", err))
	}

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateTenant(ctx, tenantId)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/clusters", wrapper.ListClusters)
	router.POST("/clusters", wrapper.CreateCluster)
	router.DELETE("/clusters/:clusterId", wrapper.DeleteCluster)
	router.GET("/clusters/:clusterId", wrapper.GetCluster)
	router.PATCH("/clusters/:clusterId", wrapper.UpdateCluster)
	router.GET("/healthz", wrapper.Healthz)
	router.GET("/install/steward.json", wrapper.InstallSteward)
	router.GET("/inventory", wrapper.QueryInventory)
	router.POST("/inventory", wrapper.UpdateInventory)
	router.GET("/tenants", wrapper.ListTenants)
	router.POST("/tenants", wrapper.CreateTenant)
	router.DELETE("/tenants/:tenantId", wrapper.DeleteTenant)
	router.GET("/tenants/:tenantId", wrapper.GetTenant)
	router.PATCH("/tenants/:tenantId", wrapper.UpdateTenant)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
