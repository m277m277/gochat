package minip

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/yiigo"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/wx"
)

func TestUploadMedia(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"type": "image",
		"media_id": "MEDIA_ID",
		"created_at": 1606717010
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(MediaUploadResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadMedia(dest, MediaImage, "../test/test.jpg"))

	assert.Nil(t, err)
	assert.Equal(t, &MediaUploadResult{
		Type:      "image",
		MediaID:   "MEDIA_ID",
		CreatedAt: 1606717010,
	}, dest)
}

func TestUploadMediaByURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=image", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{
		"errcode": 0,
		"errmsg": "ok",
		"type": "image",
		"media_id": "MEDIA_ID",
		"created_at": 1606717010
	  }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(MediaUploadResult)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadMediaByURL(dest, MediaImage, "../test/test.png", "https://golang.google.cn/doc/gopher/pkg.png"))

	assert.Nil(t, err)
	assert.Equal(t, &MediaUploadResult{
		Type:      "image",
		MediaID:   "MEDIA_ID",
		CreatedAt: 1606717010,
	}, dest)
}

func TestGetMedia(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID").Return([]byte("BUFFER"), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	dest := new(Media)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetMedia(dest, "MEDIA_ID"))

	assert.Nil(t, err)
	assert.Equal(t, "BUFFER", string(dest.Buffer))
}