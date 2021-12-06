package models

import(
	"time"
)


// SentinelAlert defines v2 incident body
type SentinelAlert struct {
	ID string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Properties IncidentProperties `json:"properties" binding:"required"`
	// below lines unavailable in docs; collected from blog sample
	ETAG string `json:"etag" binding:"required"`
	Type string `json:"type" binding:"required"`
}

// IncidentProperties details incident event details
type IncidentProperties struct {
	AdditionalData IncidentAdditionalData `json:"additionalData" binding:"-"`
	Classification string `json:"classification" binding:"-"`
	ClassificationComment string `json:"classificationComment" binding:"-"`
	ClassificationReason string `json:"classificationReason" binding:"-"`
	CreatedTimeUTC time.Time `json:"createdTimeUtc" binding:"required"`
	Description string `json:"description" binding:"required"`
	FirstActivityTimeUTC time.Time `json:"firstActivityTimeUtc" binding:"required"`
	IncidentURL string `json:"incidentUrl" binding:"required"`
	// below two lines added from blog sample; unavailable in docs
	ProviderName string `json:"providerName" binding:"required"`
	ProviderIncidentID string `json:"providerIncidentId" binding:"required"`
	IncidentNumber int `json:"incidentNumber" binding:"required"`
	LastActivityTimeUTC time.Time `json:"lastActivityTimeUtc" binding:"required"`
	Severity string `json:"severity" binding:"required"`
	Status string `json:"status" binding:"required"`
	Title string `json:"title" binding:"required"`
	Labels []IncidentLabel `json:"labels" binding:"required"`
	LastModifiedTimeUTC time.Time `json:"lastModifiedTimeUtc" binding:"required"`
	Owner IncidentOwnerInfo `json:"owner" binding:"-"`
	RelatedAnalyticRuleIDs []string `json:"relatedAnalyticRuleIds" binding:"required"`
	Comments []IncidentComment `json:"Comments" binding:"-"`
}

// IncidentAdditionalData additional event details
type IncidentAdditionalData struct {
	AlertsCount int `json:"alertsCount" binding:"required"`
	BookMarksCount int `json:"bookmarksCount" binding:"required"`
	CommentsCount int `json:"commentsCount" binding:"required"`
	AlertProductNames []string `json:"alertProductNames" binding:"required"`
	Tactics []string `json:"tactics" binding:"required"`
}

// AttackTactic item associated with incident

// IncidentComment defines incident comment
type IncidentComment struct {
	ID string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Properties IncidentCommentProperties `json:"properties" binding:"required"`
	// fields below are inconsistently identified in azure docs
	Type string `json:"type" binding:"required"`
	APIVersion string `json:"apiVersion" binding:"required"`
	ETAG string `json:"etag" binding:"required"`
}

// IncidentCommentProperties message string
type IncidentCommentProperties struct {
	Message string `json:"message" binding:"required"`
}

// IncidentLabel define label
type IncidentLabel struct {
	Name string `json:"labelName" binding:"required"`
	Type string `json:"labelType" binding:"required"`
}

// IncidentOwnerInfo define owner
type IncidentOwnerInfo struct {
	Email string `json:"email" binding:"required"`
	AssignedTo string `json:"assignedTo" binding:"required"`
	ObjectID string `json:"objectId" binding:"required"`
	UserPrincipalName string `json:"userPrincipalName" binding:"required"`
}
