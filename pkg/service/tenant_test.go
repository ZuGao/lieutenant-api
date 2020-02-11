package service

import (
	"net/http"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/labstack/echo/v4"
	"github.com/projectsyn/lieutenant-api/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestListTenants(t *testing.T) {
	e := setupTest(t)

	result := testutil.NewRequest().
		Get(APIBasePath+"/tenants/").
		WithHeader(echo.HeaderAuthorization, bearerToken).
		Go(t, e)
	assert.Equal(t, http.StatusOK, result.Code())
	tenants := &[]api.Tenant{}
	err := result.UnmarshalJsonToObject(tenants)
	assert.NoError(t, err)
	assert.NotNil(t, tenants)
	assert.Greater(t, len(*tenants), 1)
}

func TestCreateTenant(t *testing.T) {
	e := setupTest(t)

	newTenant := api.TenantProperties{
		DisplayName: pointer.ToString("My test Tenant"),
	}
	result := testutil.NewRequest().
		Post(APIBasePath+"/tenants").
		WithJsonBody(newTenant).
		WithHeader(echo.HeaderAuthorization, bearerToken).
		Go(t, e)
	assert.Equal(t, http.StatusCreated, result.Code())
	tenant := &api.Tenant{}
	err := result.UnmarshalJsonToObject(tenant)
	assert.NoError(t, err)
	assert.NotNil(t, tenant)
	assert.Contains(t, tenant.Id, api.TenantIDPrefix)
	assert.Equal(t, newTenant.DisplayName, tenant.DisplayName)
}

func TestCreateTenantFail(t *testing.T) {
	e := setupTest(t)

	result := testutil.NewRequest().
		Post(APIBasePath+"/tenants/").
		WithJsonContentType().
		WithBody([]byte("invalid-body")).
		WithHeader(echo.HeaderAuthorization, bearerToken).
		Go(t, e)
	assert.Equal(t, http.StatusBadRequest, result.Code())
	reason := &api.Reason{}
	err := result.UnmarshalJsonToObject(reason)
	assert.NoError(t, err)
	assert.NotEmpty(t, reason.Reason)
}

func TestCreateTenantEmpty(t *testing.T) {
	e := setupTest(t)

	result := testutil.NewRequest().
		Post(APIBasePath+"/tenants/").
		WithHeader(echo.HeaderAuthorization, bearerToken).
		Go(t, e)
	assert.Equal(t, http.StatusBadRequest, result.Code())
	reason := &api.Reason{}
	err := result.UnmarshalJsonToObject(reason)
	assert.NoError(t, err)
	assert.NotEmpty(t, reason.Reason)
}

func TestTenantDelete(t *testing.T) {
	e := setupTest(t)

	result := testutil.NewRequest().
		Delete(APIBasePath+"/tenants/"+tenantA.Name).
		WithHeader(echo.HeaderAuthorization, bearerToken).
		Go(t, e)
	assert.Equal(t, http.StatusNoContent, result.Code())
}

func TestTenantGet(t *testing.T) {
	e := setupTest(t)

	result := testutil.NewRequest().
		Get(APIBasePath+"/tenants/"+tenantA.Name).
		WithHeader(echo.HeaderAuthorization, bearerToken).
		Go(t, e)
	assert.Equal(t, http.StatusOK, result.Code())
	tenant := &api.Tenant{}
	err := result.UnmarshalJsonToObject(tenant)
	assert.NoError(t, err)
	assert.Equal(t, tenantA.Name, string(tenant.Id))
	assert.Equal(t, tenantA.Spec.DisplayName, *tenant.DisplayName)
	assert.Equal(t, tenantA.Spec.GitRepoURL, *tenant.GitRepo)
}

func TestTenantUpdateEmpty(t *testing.T) {
	e := setupTest(t)

	result := testutil.NewRequest().
		Patch(APIBasePath+"/tenants/1").
		WithHeader(echo.HeaderAuthorization, bearerToken).
		Go(t, e)
	assert.Equal(t, http.StatusBadRequest, result.Code())
	reason := &api.Reason{}
	err := result.UnmarshalJsonToObject(reason)
	assert.NoError(t, err)
	assert.NotEmpty(t, reason.Reason)
}

func TestTenantUpdate(t *testing.T) {
	e := setupTest(t)

	updateTenant := &api.TenantProperties{
		DisplayName: pointer.ToString("New Name"),
	}
	result := testutil.NewRequest().
		Patch(APIBasePath+"/tenants/"+tenantB.Name).
		WithJsonBody(updateTenant).
		WithContentType(api.ContentJSONPatch).
		WithHeader(echo.HeaderAuthorization, bearerToken).
		Go(t, e)
	assert.Equal(t, http.StatusNoContent, result.Code())
}
