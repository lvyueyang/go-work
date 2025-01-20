package service

import (
	"bytes"
	"server/utils"
	"text/template"
)

type NotifyService struct{}

var notifyService = new(NotifyService)

func NewNotifyService() *NotifyService {
	return notifyService
}

// SendCaptchaEmail 发送验证码邮件
func (s *NotifyService) SendCaptchaEmail(subject, title, target, code string) error {
	tmpl, err := template.ParseFiles("resource/template/email_captcha.html")
	if err != nil {
		return err
	}

	var w bytes.Buffer
	if err := tmpl.Execute(&w, map[string]string{
		"title": title,
		"code":  code,
	}); err != nil {
		return err
	}

	c := utils.SendEmail{
		To:      []string{target},
		Title:   subject,
		Content: w.String(),
	}
	return c.Send()
}
