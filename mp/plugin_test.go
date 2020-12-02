package mp

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestApplyPlugin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/plugin?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", ApplyPlugin("aaaa", "hello"))

	assert.Nil(t, err)
}

func TestGetPluginDevApplyList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/devplugin?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"apply_list": [{
			"appid": "xxxxxxxxxxxxx",
			"status": 1,
			"nickname": "名称",
			"headimgurl": "**********",
			"reason": "polo has gone",
			"apply_url": "*******",
			"create_time": "1536305096",
			"categories": [{
				"first": "IT科技",
				"second": "硬件与设备"
			}]
		}]
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := make([]PluginDevApplyInfo, 0)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetPluginDevApplyList(1, 10, &dest))

	assert.Nil(t, err)
	assert.Equal(t, []PluginDevApplyInfo{
		{
			AppID:      "xxxxxxxxxxxxx",
			Status:     1,
			Nickname:   "名称",
			HeadImgURL: "**********",
			Categories: []wx.X{
				{
					"first":  "IT科技",
					"second": "硬件与设备",
				},
			},
			CreateTime: "1536305096",
			ApplyURL:   "*******",
			Reason:     "polo has gone",
		},
	}, dest)
}

func TestGetPluginList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/plugin?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"plugin_list": [{
			"appid": "aaaa",
			"status": 1,
			"nickname": "插件昵称",
			"headimgurl": "http://plugin.qq.com"
		}]
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := make([]PluginInfo, 0)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetPluginList(&dest))

	assert.Nil(t, err)
	assert.Equal(t, []PluginInfo{
		{
			AppID:      "aaaa",
			Status:     1,
			Nickname:   "插件昵称",
			HeadImgURL: "http://plugin.qq.com",
		},
	}, dest)
}

func TestSetDevPluginApplyStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/devplugin?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SetDevPluginApplyStatus("dev_agree", "aaaa", ""))

	assert.Nil(t, err)
}

func TestUnbindPlugin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/wxa/plugin?access_token=ACCESS_TOKEN", gomock.AssignableToTypeOf(postBody)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UnbindPlugin("aaaa"))

	assert.Nil(t, err)
}