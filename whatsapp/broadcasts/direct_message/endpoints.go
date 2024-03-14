package direct_message

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Parameter struct {
	Header  HeaderParameter   `json:"header"`
	Body    []BodyParameter   `json:"body"`
	Buttons []ButtonParameter `json:"buttons,omitempty"`
}

type SendDirectMessageRequest struct {
	ToName               string     `json:"to_name"`                //Name of the recipient
	ToNumber             string     `json:"to_number"`              //Whatsapp number of the recipient
	MessageTemplateId    uuid.UUID  `json:"message_template_id"`    //MessageTemplateId in Qontak Platform
	ChannelIntegrationId uuid.UUID  `json:"channel_integration_id"` //ChannelIntegrationId in Qontak Platform
	Language             *Language  `json:"language,omitempty"`
	Parameters           *Parameter `json:"parameters,omitempty"`
}

func (a *SendDirectMessageRequest) ToJSON() []byte {
	j, _ := json.Marshal(a)
	return j
}

type MessageStatusStatistic struct {
	Failed    int `json:"failed"`
	Delivered int `json:"delivered"`
	Read      int `json:"read"`
	Pending   int `json:"pending"`
	Sent      int `json:"sent"`
}

type SendDirectMessageResponseSuccess struct {
	Id                   uuid.UUID              `json:"id"`
	Name                 string                 `json:"name"`
	OrganizationId       uuid.UUID              `json:"organization_id"`
	ChannelIntegrationId uuid.UUID              `json:"channel_integration_id"`
	ContactListId        uuid.NullUUID          `json:"contact_list_id,omitempty"`
	ContactId            uuid.UUID              `json:"contact_id"`
	TargetChannel        string                 `json:"target_channel"`
	SendAt               time.Time              `json:"send_at"`
	ExecuteStatus        string                 `json:"execute_status"`
	ExecuteType          string                 `json:"execute_type"`
	CreatedAt            time.Time              `json:"created_at"`
	Statistic            MessageStatusStatistic `json:"message_status_count"`
}

type SendDirectMessageResponse struct {
	Status string                           `json:"status,omitempty"`
	Data   SendDirectMessageResponseSuccess `json:"data"`
}
