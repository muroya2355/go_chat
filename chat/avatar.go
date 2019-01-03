package main

import (
	"errors"
)

// ErrNoAvatar は Avatar インスタンスがアバターのURLを返すことができない場合に発生するエラーです
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません")

// Avatar はユーザのプロフィール画像を表す型です
type Avatar interface {
	// GetAvatarURL は指定されたクライアントのアバターのURLを返します
	// 問題が発生した場合はエラーを返します。特に、URLを取得できなかった場合にはErrNoAvatarURLを返します
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}
