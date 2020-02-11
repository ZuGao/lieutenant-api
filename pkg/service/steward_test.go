package service

import (
	"net/http"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/testutil"
	"github.com/projectsyn/lieutenant-api/pkg/api"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
)

func TestInstallSteward(t *testing.T) {
	e := setupTest(t)

	result := testutil.NewRequest().
		Get(APIBasePath+"/install/steward.json?token="+clusterA.Status.BootstrapToken.Token).
		Go(t, e)
	assert.Equal(t, http.StatusOK, result.Code())
	manifests := &corev1.List{}
	err := result.UnmarshalJsonToObject(&manifests)
	assert.NoError(t, err)
	assert.Len(t, manifests.Items, 6)
	decoder := json.NewSerializer(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme, true)
	foundSecret := false
	for _, item := range manifests.Items {
		obj, err := runtime.Decode(decoder, item.Raw)
		assert.NoError(t, err)
		if secret, ok := obj.(*corev1.Secret); ok {
			foundSecret = true
			assert.Equal(t, secret.StringData["token"], string(clusterASecret.Data["token"]))
		}
	}
	assert.True(t, foundSecret, "Could not find secret with steward token")
}

func TestInstallStewardNoToken(t *testing.T) {
	e := setupTest(t)

	result := testutil.NewRequest().
		Get(APIBasePath+"/install/steward.json").
		Go(t, e)
	assert.Equal(t, http.StatusBadRequest, result.Code())
	reason := &api.Reason{}
	err := result.UnmarshalJsonToObject(reason)
	assert.NoError(t, err)
	assert.NotEmpty(t, reason.Reason)
}

func TestInstallStewardInvalidToken(t *testing.T) {
	e := setupTest(t)

	result := testutil.NewRequest().
		Get(APIBasePath+"/install/steward.json?token=NonExistentToken").
		Go(t, e)
	assert.Equal(t, http.StatusUnauthorized, result.Code())
	reason := &api.Reason{}
	err := result.UnmarshalJsonToObject(reason)
	assert.NoError(t, err)
	assert.NotEmpty(t, reason.Reason)
}

func TestInstallStewardUsedToken(t *testing.T) {
	e := setupTest(t)

	result := testutil.NewRequest().
		Get(APIBasePath+"/install/steward.json?token="+clusterB.Status.BootstrapToken.Token).
		Go(t, e)
	assert.Equal(t, http.StatusUnauthorized, result.Code())
	reason := &api.Reason{}
	err := result.UnmarshalJsonToObject(reason)
	assert.NoError(t, err)
	assert.NotEmpty(t, reason.Reason)
}
