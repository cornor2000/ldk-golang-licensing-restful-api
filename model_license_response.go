/*
 * Sentinel LDK Runtime RESTful API  documentation
 *
 * This pages documents Sentinel LDK Runtime RESTful API Definition
 */

package ldklicensingretfulapi

type LicenseResponse struct {
	ErrorCode     int          `json:"errorCode"`
	SessionId     string       `json:"sessionId,omitempty"`
	SessionInfo   *SessionInfo `json:"sessionInfo,omitempty"`
	FeatureId     int32        `json:"featureId,omitempty"`
	ProductId     int32        `json:"productId,omitempty"`
	KeyId         int64        `json:"keyId,omitempty"`
	Message       string       `json:"message"`
	LmAccessToken string       `json:"lmAccessToken"`
}
