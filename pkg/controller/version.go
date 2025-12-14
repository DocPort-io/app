package controller

import (
	"app/pkg/apperrors"
	"app/pkg/dto"
	"app/pkg/httputil"
	"app/pkg/service"
	"net/http"

	"github.com/go-chi/render"
)

type VersionController struct {
	versionService *service.VersionService
}

func NewVersionController(versionService *service.VersionService) *VersionController {
	return &VersionController{versionService: versionService}
}

// FindAllVersions godoc
//
//	@summary	Find all versions
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		projectId	query		uint	true	"Project identifier"
//	@success	200			{object}	dto.ListVersionsResponseDto
//	@router		/api/v1/versions [get]
func (c *VersionController) FindAllVersions(w http.ResponseWriter, r *http.Request) {
	projectId, err := httputil.QueryParamInt64(r, "projectId", true)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	versions, total, err := c.versionService.FindAllVersions(r.Context(), projectId)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	httputil.Render(w, r, dto.ToListVersionsResponse(versions, total))
}

// GetVersion godoc
//
//	@summary	Get a version
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		id	path		uint	true	"Version identifier"
//	@success	200	{object}	dto.VersionResponseDto
//	@router		/api/v1/versions/{id} [get]
func (c *VersionController) GetVersion(w http.ResponseWriter, r *http.Request) {
	versionId, err := httputil.URLParamInt64(r, "versionId")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	version, err := c.versionService.FindVersionById(r.Context(), versionId)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	httputil.Render(w, r, dto.ToVersionResponse(version))
}

// CreateVersion godoc
//
//	@summary	Create a version
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		request	body		dto.CreateVersionDto	true	"Create a version"
//	@success	201		{object}	dto.VersionResponseDto
//	@router		/api/v1/versions [post]
func (c *VersionController) CreateVersion(w http.ResponseWriter, r *http.Request) {
	input := &dto.CreateVersionDto{}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	version, err := c.versionService.CreateVersion(r.Context(), input)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	httputil.Render(w, r, dto.ToVersionResponse(version))
}

// UpdateVersion godoc
//
//	@summary	Update a version
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		id		path		uint					true	"Version identifier"
//	@param		request	body		dto.UpdateVersionDto	true	"Update a version"
//	@success	200		{object}	dto.VersionResponseDto
//	@router		/api/v1/versions/{id} [put]
func (c *VersionController) UpdateVersion(w http.ResponseWriter, r *http.Request) {
	input := &dto.UpdateVersionDto{}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	versionId, err := httputil.URLParamInt64(r, "versionId")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	version, err := c.versionService.UpdateVersion(r.Context(), versionId, input)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	httputil.Render(w, r, dto.ToVersionResponse(version))
}

// DeleteVersion godoc
//
//	@summary	Delete a version
//	@tags		versions
//	@accept		json
//	@param		id	path	uint	true	"Version identifier"
//	@success	204
//	@router		/api/v1/versions/{id} [delete]
func (c *VersionController) DeleteVersion(w http.ResponseWriter, r *http.Request) {
	versionId, err := httputil.URLParamInt64(r, "versionId")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	err = c.versionService.DeleteVersion(r.Context(), versionId)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.NoContent(w, r)
}

// AttachFileToVersion godoc
//
//	@summary	Attaches a file to a version
//	@tags		versions
//	@accept		json
//	@param		id		path	uint						true	"Version identifier"
//	@param		request	body	dto.AttachFileToVersionDto	true	"File to attach"
//	@success	204
//	@router		/api/v1/versions/{id}/attach-file [post]
func (c *VersionController) AttachFileToVersion(w http.ResponseWriter, r *http.Request) {
	versionId, err := httputil.URLParamInt64(r, "versionId")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	input := &dto.AttachFileToVersionDto{}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	err = c.versionService.AttachFileToVersion(r.Context(), versionId, input)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.NoContent(w, r)
}

// DetachFileFromVersion godoc
//
//	@summary	Detach a file from a version
//	@tags		versions
//	@accept		json
//	@param		id		path	uint							true	"Version identifier"
//	@param		request	body	dto.DetachFileFromVersionDto	true	"File to detach"
//	@success	204
//	@router		/api/v1/versions/{id}/detach-file [post]
func (c *VersionController) DetachFileFromVersion(w http.ResponseWriter, r *http.Request) {
	versionId, err := httputil.URLParamInt64(r, "versionId")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	input := &dto.DetachFileFromVersionDto{}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	err = c.versionService.DetachFileFromVersion(r.Context(), versionId, input)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.NoContent(w, r)
}
