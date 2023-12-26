package types

type DeliveryStatus string

const (
	DeliveryStatusNotSent DeliveryStatus = "not_sent"
	DeliveryStatusSent    DeliveryStatus = "sent"
	DeliveryStatusFailed  DeliveryStatus = "failed"
)
