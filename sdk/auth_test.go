package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
	"gotest.tools/assert"
	"testing"
)

func TestGetTenantAccessTokenInternal(t *testing.T) {
	resp, err := GetTenantAccessTokenInternal(consts.TestAppId, consts.TestAppSecret)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestGetAppAccessTokenInternal(t *testing.T) {
	resp, err := GetAppAccessToken(consts.TestAppId, consts.TestAppSecret, "")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestAppTicketResend(t *testing.T) {
	resp, err := AppTicketResend(consts.TestAppId, consts.TestAppSecret)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTokenLoginValidate(t *testing.T) {
	resp, err := TokenLoginValidate("a-566311d2cf88d054a4fcfc23233a448f2fccba11", "1c3a78be18ac815a")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}