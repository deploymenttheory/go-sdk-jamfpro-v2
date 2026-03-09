package enrollment_customization_preview

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// EnrollmentCustomizationPreviewServiceInterface defines the interface for enrollment customization panel operations.
	//
	// Manages enrollment customization panels (LDAP, SSO, Text) and markdown parsing.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/enrollment-customization-preview
	EnrollmentCustomizationPreviewServiceInterface interface {
		// ParseMarkdown parses markdown to HTML.
		ParseMarkdown(ctx context.Context, request *RequestParseMarkdown) (*ResponseParseMarkdown, *resty.Response, error)

		// GetAllPanels returns all panels for an enrollment customization.
		GetAllPanels(ctx context.Context, id string) (*PanelListResponse, *resty.Response, error)

		// GetPanelByID returns a single panel by ID.
		GetPanelByID(ctx context.Context, id, panelID string) (*ResourcePanel, *resty.Response, error)

		// DeletePanel deletes a panel from an enrollment customization.
		DeletePanel(ctx context.Context, id, panelID string) (*resty.Response, error)

		// CreateLdapPanel creates an LDAP panel.
		CreateLdapPanel(ctx context.Context, id string, request *ResourceLdapPanel) (*ResourceLdapPanel, *resty.Response, error)

		// GetLdapPanel returns an LDAP panel by ID.
		GetLdapPanel(ctx context.Context, id, panelID string) (*ResourceLdapPanel, *resty.Response, error)

		// UpdateLdapPanel updates an LDAP panel.
		UpdateLdapPanel(ctx context.Context, id, panelID string, request *ResourceLdapPanel) (*ResourceLdapPanel, *resty.Response, error)

		// DeleteLdapPanel deletes an LDAP panel.
		DeleteLdapPanel(ctx context.Context, id, panelID string) (*resty.Response, error)

		// CreateSsoPanel creates an SSO panel.
		CreateSsoPanel(ctx context.Context, id string, request *ResourceSsoPanel) (*ResourceSsoPanel, *resty.Response, error)

		// GetSsoPanel returns an SSO panel by ID.
		GetSsoPanel(ctx context.Context, id, panelID string) (*ResourceSsoPanel, *resty.Response, error)

		// UpdateSsoPanel updates an SSO panel.
		UpdateSsoPanel(ctx context.Context, id, panelID string, request *ResourceSsoPanel) (*ResourceSsoPanel, *resty.Response, error)

		// DeleteSsoPanel deletes an SSO panel.
		DeleteSsoPanel(ctx context.Context, id, panelID string) (*resty.Response, error)

		// CreateTextPanel creates a text panel.
		CreateTextPanel(ctx context.Context, id string, request *ResourceTextPanel) (*ResourceTextPanel, *resty.Response, error)

		// GetTextPanel returns a text panel by ID.
		GetTextPanel(ctx context.Context, id, panelID string) (*ResourceTextPanel, *resty.Response, error)

		// UpdateTextPanel updates a text panel.
		UpdateTextPanel(ctx context.Context, id, panelID string, request *ResourceTextPanel) (*ResourceTextPanel, *resty.Response, error)

		// DeleteTextPanel deletes a text panel.
		DeleteTextPanel(ctx context.Context, id, panelID string) (*resty.Response, error)

		// GetTextPanelMarkdown returns the markdown for a text panel.
		GetTextPanelMarkdown(ctx context.Context, id, panelID string) (*ResponseTextPanelMarkdown, *resty.Response, error)
	}

	// Service handles communication with the enrollment customization preview methods of the Jamf Pro API.
	EnrollmentCustomizationPreview struct {
		client transport.HTTPClient
	}
)

var _ EnrollmentCustomizationPreviewServiceInterface = (*EnrollmentCustomizationPreview)(nil)

func NewEnrollmentCustomizationPreview(client transport.HTTPClient) *EnrollmentCustomizationPreview {
	return &EnrollmentCustomizationPreview{client: client}
}

// ParseMarkdown parses markdown to HTML.
// URL: POST /api/v1/enrollment-customization/parse-markdown
func (s *EnrollmentCustomizationPreview) ParseMarkdown(ctx context.Context, request *RequestParseMarkdown) (*ResponseParseMarkdown, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/parse-markdown", EndpointEnrollmentCustomizationPreviewV1)

	var result ResponseParseMarkdown

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/all", EndpointEnrollmentCustomizationPreviewV1, id)

	var result PanelListResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/all/%s", EndpointEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourcePanel

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/all/%s", EndpointEnrollmentCustomizationPreviewV1, id, panelID)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/ldap", EndpointEnrollmentCustomizationPreviewV1, id)

	var result ResourceLdapPanel

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/ldap/%s", EndpointEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourceLdapPanel

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/ldap/%s", EndpointEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourceLdapPanel

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/ldap/%s", EndpointEnrollmentCustomizationPreviewV1, id, panelID)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/sso", EndpointEnrollmentCustomizationPreviewV1, id)

	var result ResourceSsoPanel

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/sso/%s", EndpointEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourceSsoPanel

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/sso/%s", EndpointEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourceSsoPanel

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/sso/%s", EndpointEnrollmentCustomizationPreviewV1, id, panelID)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/text", EndpointEnrollmentCustomizationPreviewV1, id)

	var result ResourceTextPanel

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/text/%s", EndpointEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourceTextPanel

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/text/%s", EndpointEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResourceTextPanel

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/text/%s", EndpointEnrollmentCustomizationPreviewV1, id, panelID)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/text/%s/markdown", EndpointEnrollmentCustomizationPreviewV1, id, panelID)

	var result ResponseTextPanelMarkdown

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
