package enrollment_customization_preview

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the enrollment customization preview methods of the Jamf Pro API.
	EnrollmentCustomizationPreview struct {
		client transport.HTTPClient
	}
)

func NewEnrollmentCustomizationPreview(client transport.HTTPClient) *EnrollmentCustomizationPreview {
	return &EnrollmentCustomizationPreview{client: client}
}

// ParseMarkdown parses markdown to HTML.
// URL: POST /api/v1/enrollment-customization/parse-markdown
func (s *EnrollmentCustomizationPreview) ParseMarkdown(ctx context.Context, request *RequestParseMarkdown) (*ResponseParseMarkdown, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/parse-markdown", constants.EndpointJamfProEnrollmentCustomizationPreviewV1)

	var result ResponseParseMarkdown

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetAllPanels returns all panels for an enrollment customization.
// URL: GET /api/v1/enrollment-customization/{id}/all
func (s *EnrollmentCustomizationPreview) GetAllPanels(ctx context.Context, id string) (*PanelListResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/all", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id)

	var result PanelListResponse

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetPanelByID returns a single panel by ID.
// URL: GET /api/v1/enrollment-customization/{id}/all/{panel-id}
func (s *EnrollmentCustomizationPreview) GetPanelByID(ctx context.Context, id, panelID string) (*ResourcePanel, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}
	if panelID == "" {
		return nil, nil, fmt.Errorf("panel ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/all/%s", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourcePanel

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeletePanel deletes a panel from an enrollment customization.
// URL: DELETE /api/v1/enrollment-customization/{id}/all/{panel-id}
func (s *EnrollmentCustomizationPreview) DeletePanel(ctx context.Context, id, panelID string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("enrollment customization ID is required")
	}
	if panelID == "" {
		return nil, fmt.Errorf("panel ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/all/%s", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id, panelID)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// CreateLdapPanel creates an LDAP panel.
// URL: POST /api/v1/enrollment-customization/{id}/ldap
func (s *EnrollmentCustomizationPreview) CreateLdapPanel(ctx context.Context, id string, request *ResourceLdapPanel) (*ResourceLdapPanel, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}
	if request.Title == "" {
		return nil, nil, fmt.Errorf("title is required")
	}

	endpoint := fmt.Sprintf("%s/%s/ldap", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id)

	var result ResourceLdapPanel

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetLdapPanel returns an LDAP panel by ID.
// URL: GET /api/v1/enrollment-customization/{id}/ldap/{panel-id}
func (s *EnrollmentCustomizationPreview) GetLdapPanel(ctx context.Context, id, panelID string) (*ResourceLdapPanel, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}
	if panelID == "" {
		return nil, nil, fmt.Errorf("panel ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/ldap/%s", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourceLdapPanel

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateLdapPanel updates an LDAP panel.
// URL: PUT /api/v1/enrollment-customization/{id}/ldap/{panel-id}
func (s *EnrollmentCustomizationPreview) UpdateLdapPanel(ctx context.Context, id, panelID string, request *ResourceLdapPanel) (*ResourceLdapPanel, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}
	if panelID == "" {
		return nil, nil, fmt.Errorf("panel ID is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}
	if request.Title == "" {
		return nil, nil, fmt.Errorf("title is required")
	}

	endpoint := fmt.Sprintf("%s/%s/ldap/%s", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourceLdapPanel

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteLdapPanel deletes an LDAP panel.
// URL: DELETE /api/v1/enrollment-customization/{id}/ldap/{panel-id}
func (s *EnrollmentCustomizationPreview) DeleteLdapPanel(ctx context.Context, id, panelID string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("enrollment customization ID is required")
	}
	if panelID == "" {
		return nil, fmt.Errorf("panel ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/ldap/%s", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id, panelID)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// CreateSsoPanel creates an SSO panel.
// URL: POST /api/v1/enrollment-customization/{id}/sso
func (s *EnrollmentCustomizationPreview) CreateSsoPanel(ctx context.Context, id string, request *ResourceSsoPanel) (*ResourceSsoPanel, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}

	endpoint := fmt.Sprintf("%s/%s/sso", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id)

	var result ResourceSsoPanel

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetSsoPanel returns an SSO panel by ID.
// URL: GET /api/v1/enrollment-customization/{id}/sso/{panel-id}
func (s *EnrollmentCustomizationPreview) GetSsoPanel(ctx context.Context, id, panelID string) (*ResourceSsoPanel, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}
	if panelID == "" {
		return nil, nil, fmt.Errorf("panel ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/sso/%s", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourceSsoPanel

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateSsoPanel updates an SSO panel.
// URL: PUT /api/v1/enrollment-customization/{id}/sso/{panel-id}
func (s *EnrollmentCustomizationPreview) UpdateSsoPanel(ctx context.Context, id, panelID string, request *ResourceSsoPanel) (*ResourceSsoPanel, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}
	if panelID == "" {
		return nil, nil, fmt.Errorf("panel ID is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}

	endpoint := fmt.Sprintf("%s/%s/sso/%s", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourceSsoPanel

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteSsoPanel deletes an SSO panel.
// URL: DELETE /api/v1/enrollment-customization/{id}/sso/{panel-id}
func (s *EnrollmentCustomizationPreview) DeleteSsoPanel(ctx context.Context, id, panelID string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("enrollment customization ID is required")
	}
	if panelID == "" {
		return nil, fmt.Errorf("panel ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/sso/%s", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id, panelID)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// CreateTextPanel creates a text panel.
// URL: POST /api/v1/enrollment-customization/{id}/text
func (s *EnrollmentCustomizationPreview) CreateTextPanel(ctx context.Context, id string, request *ResourceTextPanel) (*ResourceTextPanel, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}
	if request.Title == "" {
		return nil, nil, fmt.Errorf("title is required")
	}
	if request.Body == "" {
		return nil, nil, fmt.Errorf("body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/text", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id)

	var result ResourceTextPanel

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetTextPanel returns a text panel by ID.
// URL: GET /api/v1/enrollment-customization/{id}/text/{panel-id}
func (s *EnrollmentCustomizationPreview) GetTextPanel(ctx context.Context, id, panelID string) (*ResourceTextPanel, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}
	if panelID == "" {
		return nil, nil, fmt.Errorf("panel ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/text/%s", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourceTextPanel

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateTextPanel updates a text panel.
// URL: PUT /api/v1/enrollment-customization/{id}/text/{panel-id}
func (s *EnrollmentCustomizationPreview) UpdateTextPanel(ctx context.Context, id, panelID string, request *ResourceTextPanel) (*ResourceTextPanel, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}
	if panelID == "" {
		return nil, nil, fmt.Errorf("panel ID is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}
	if request.Title == "" {
		return nil, nil, fmt.Errorf("title is required")
	}
	if request.Body == "" {
		return nil, nil, fmt.Errorf("body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/text/%s", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourceTextPanel

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteTextPanel deletes a text panel.
// URL: DELETE /api/v1/enrollment-customization/{id}/text/{panel-id}
func (s *EnrollmentCustomizationPreview) DeleteTextPanel(ctx context.Context, id, panelID string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("enrollment customization ID is required")
	}
	if panelID == "" {
		return nil, fmt.Errorf("panel ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/text/%s", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id, panelID)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetTextPanelMarkdown returns the markdown for a text panel.
// URL: GET /api/v1/enrollment-customization/{id}/text/{panel-id}/markdown
func (s *EnrollmentCustomizationPreview) GetTextPanelMarkdown(ctx context.Context, id, panelID string) (*ResponseTextPanelMarkdown, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}
	if panelID == "" {
		return nil, nil, fmt.Errorf("panel ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/text/%s/markdown", constants.EndpointJamfProEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResponseTextPanelMarkdown

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
