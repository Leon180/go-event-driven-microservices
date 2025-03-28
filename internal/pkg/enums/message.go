package enums

type MessageDeliveryType int

const (
	MessageDeliveryTypeOutbox   MessageDeliveryType = 1
	MessageDeliveryTypeInbox    MessageDeliveryType = 2
	MessageDeliveryTypeInternal MessageDeliveryType = 4
)

type MessageStatus int

const (
	MessageStatusStored    MessageStatus = 1
	MessageStatusProcessed MessageStatus = 2
)
