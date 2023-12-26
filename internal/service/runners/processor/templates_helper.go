package processor

import (
	"encoding/json"

	"github.com/apodeixis/notifications-router-svc/internal/types"

	"github.com/pkg/errors"

	"github.com/apodeixis/notifications-router-svc/internal/config"
	"github.com/apodeixis/notifications-router-svc/internal/data"
	"github.com/apodeixis/notifications-router-svc/internal/providers/templates"
)

type templatesHelper struct {
	templatesProvider templates.TemplatesProvider
	notificatorCfg    *config.NotificatorConfig
}

func (h *templatesHelper) buildMessage(channel string, notification data.Notification) (data.Message, error) {
	if notification.Message.Type != types.MessageTypeTemplate {
		return notification.Message, nil
	}

	templateAttrs, err := unmarshalTemplateMessageAttributes(notification.Message.Attributes)
	if err != nil {
		return data.Message{}, errors.Wrap(err, "failed to unmarshal template attributes")
	}

	locale := h.notificatorCfg.DefaultLocale
	if templateAttrs.Locale != nil {
		locale = *templateAttrs.Locale
	}

	rawMes, err := h.templatesProvider.GetTemplate(notification.Topic, channel, locale)
	if err != nil {
		return data.Message{}, errors.Wrap(err, "failed to download template")
	}
	if rawMes == nil {
		return data.Message{}, errors.New("template not found")
	}

	if templateAttrs.Payload != nil {
		rawAttrs, err := interpolate(string(rawMes), *templateAttrs.Payload)
		if err != nil {
			return data.Message{}, errors.Wrap(err, "failed to interpolate template")
		}
		rawMes = rawAttrs
	}

	var result data.Message
	if err = json.Unmarshal(rawMes, &result); err != nil {
		return data.Message{}, errors.Wrap(err, "failed to marshal template to message")
	}

	if len(templateAttrs.Files) > 0 {
		result.Attributes["files"] = templateAttrs.Files
	}

	return result, nil
}

func (h *templatesHelper) getLocale(templateAttrs *data.TemplateMessageAttributes) string {
	if templateAttrs.Locale != nil {
		return *templateAttrs.Locale
	}
	return h.notificatorCfg.DefaultLocale
}

func unmarshalTemplateMessageAttributes(attrs map[string]interface{}) (*data.TemplateMessageAttributes, error) {
	templateAttrs := new(data.TemplateMessageAttributes)

	if payload, ok := attrs["payload"]; ok {
		payloadJSON, err := json.Marshal(payload)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal payload")
		}
		templateAttrs.Payload = new(json.RawMessage)
		*templateAttrs.Payload = payloadJSON
	}

	if locale, ok := attrs["locale"].(string); ok {
		templateAttrs.Locale = &locale
	}

	if files, ok := attrs["files"].([]interface{}); ok {
		for _, file := range files {
			if fileStr, ok := file.(string); ok {
				templateAttrs.Files = append(templateAttrs.Files, fileStr)
			}
		}
	}

	return templateAttrs, nil
}
