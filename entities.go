package main

// Entities were created using https://mholt.github.io/json-to-go/

// groups.json call
type Group struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Locked       bool   `json:"locked"`
	Backup       bool   `json:"backup"`
	ReceiverInfo string `json:"receiver_info"`
	Stamp        int    `json:"stamp"`
	LastMailing  int    `json:"last_mailing"`
	LastChanged  int    `json:"last_changed"`
}

type Mailings struct {
	Finished []Mailing `json:"finished"`
	Draft    []Mailing `json:"draft"`
}

// mailings.json/<id> call
type Mailing struct {
	ID              string `json:"id"`
	CategoryID      string `json:"category_id"`
	Name            string `json:"name"`
	Subject         string `json:"subject"`
	SenderName      string `json:"sender_name"`
	SenderEmail     string `json:"sender_email"`
	BodyHTML        string `json:"body_html"`
	BodyText        string `json:"body_text"`
	Stamp           int    `json:"stamp"`
	LastChanged     int    `json:"last_changed"`
	Started         int    `json:"started"`
	Finished        int    `json:"finished"`
	Ready           int    `json:"ready"`
	MType           int    `json:"m_type"`
	Position        int    `json:"position"`
	IsHTML          bool   `json:"is_html"`
	IsText          bool   `json:"is_text"`
	IsFixed         bool   `json:"is_fixed"`
	IsMailing       bool   `json:"is_mailing"`
	IsCampaign      bool   `json:"is_campaign"`
	IsDynamic       bool   `json:"is_dynamic"`
	IsSplittest     bool   `json:"is_splittest"`
	IsTrackable     bool   `json:"is_trackable"`
	TypeName        string `json:"type_name"`
	ExternalID      string `json:"external_id"`
	UnsubscribeLink string `json:"unsubscribe_link"`
	LastVersionID   string `json:"last_version_id"`
	SendOn          int    `json:"send_on"`
	MailingGroups   struct {
		GroupIds []string `json:"group_ids"`
	} `json:"mailing_groups"`
}

type Editor string

const (
	EditorWizard    = "wizard"
	EditorFreeForm  = "freeform"
	EditorAdvanced  = "advanced"
	EditorPlainText = "plaintext"
)

type LinkTrackingType string

const (
	LinkTrackingTypeGoogle    = "google"
	LinkTrackingTypeIntelliAD = "intelliad"
	LinkTrackingTypeCRConnect = "crconnect"
)

type AddMailing struct {
	Name        string `json:"name"`
	Subject     string `json:"subject"`
	SenderName  string `json:"sender_name"`
	SenderEmail string `json:"sender_email"`
	Content     struct {
		Type string `json:"type"`
		HTML string `json:"html"`
		Text string `json:"text"`
	} `json:"content"`
	Receivers struct {
		Groups []string `json:"groups"`
		Filter string   `json:"filter"`
	} `json:"receivers"`
	Settings struct {
		Editor             Editor `json:"editor"`
		OpenTracking       bool   `json:"open_tracking"`
		ClickTracking      bool   `json:"click_tracking"`
		LinkTrackingURL    string `json:"link_tracking_url"`
		LinkTrackingType   string `json:"link_tracking_type"`
		GoogleCampaignName string `json:"google_campaign_name"`
		UnsubscribeFormID  string `json:"unsubscribe_form_id"`
		CampaignID         string `json:"campaign_id"`
		CategoryID         string `json:"category_id"`
	} `json:"settings"`
}
