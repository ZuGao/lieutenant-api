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

	// Cluster configuration catalog Git repository, usually generated by the API
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

	// Tenant configuration Git repository, usually generated by the API
	GitRepo *string `json:"gitRepo,omitempty"`

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
