package controller

import (
	"app/pkg/apperrors"
	"app/pkg/database"
	"app/pkg/dto"
	"app/pkg/httputil"
	"app/pkg/service"
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

const VersionCtxKey = "version"

type VersionController struct {
	versionService service.VersionService
}

func NewVersionController(versionService service.VersionService) *VersionController {
	return &VersionController{versionService: versionService}
}

func (c *VersionController) VersionCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		versionId, err := httputil.URLParamInt64(r, "versionId")
		if err != nil {
			httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
			return
		}

		result, err := c.versionService.FindVersionById(r.Context(), &dto.FindVersionByIdParams{ID: versionId})
		if err != nil {
			if errors.Is(err, apperrors.ErrNotFound) {
				httputil.Render(w, r, apperrors.ErrHTTPNotFoundError())
				return
			}

			httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
			return
		}

		ctx := context.WithValue(r.Context(), VersionCtxKey, result.Version)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getVersion(ctx context.Context) *database.Version {
	return ctx.Value(VersionCtxKey).(*database.Version)
}

// FindAllVersions godoc
//
//	@summary	Find all versions
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		projectId	query		uint	true	"Project identifier"
//	@success	200			{object}	dto.ListVersionsResponse
//	@failure	400			{object}	apperrors.ErrResponse
//	@failure	500			{object}	apperrors.ErrResponse
//	@router		/api/v1/versions [get]
func (c *VersionController) FindAllVersions(w http.ResponseWriter, r *http.Request) {
	projectId, err := httputil.QueryParamInt64(r, "projectId")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	result, err := c.versionService.FindAllVersions(r.Context(), &dto.FindAllVersionsParams{ProjectID: projectId})
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	httputil.Render(w, r, dto.ToListVersionsResponse(result.Versions, result.Total))
}

// GetVersion godoc
//
//	@summary	Get a version
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		versionId	path		uint	true	"Version identifier"
//	@success	200			{object}	dto.VersionResponse
//	@failure	400			{object}	apperrors.ErrResponse
//	@failure	404			{object}	apperrors.ErrResponse
//	@failure	500			{object}	apperrors.ErrResponse
//	@router		/api/v1/versions/{versionId} [get]
func (c *VersionController) GetVersion(w http.ResponseWriter, r *http.Request) {
	version := getVersion(r.Context())
	httputil.Render(w, r, dto.ToVersionResponse(version))
}

// CreateVersion godoc
//
//	@summary	Create a version
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		request	body		dto.CreateVersionRequest	true	"Create a version"
//	@success	201		{object}	dto.VersionResponse
//	@failure	400		{object}	apperrors.ErrResponse
//	@failure	500		{object}	apperrors.ErrResponse
//	@router		/api/v1/versions [post]
func (c *VersionController) CreateVersion(w http.ResponseWriter, r *http.Request) {
	input := &dto.CreateVersionRequest{}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	// Map controller DTO to service params
	createParams := &dto.CreateVersionParams{
		Name:        input.Name,
		Description: input.Description,
		ProjectID:   input.ProjectId,
	}
	createResult, err := c.versionService.CreateVersion(r.Context(), createParams)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	httputil.Render(w, r, dto.ToVersionResponse(createResult.Version))
}

// UpdateVersion godoc
//
//	@summary	Update a version
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		versionId	path		uint					true	"Version identifier"
//	@param		request		body		dto.UpdateVersionRequest	true	"Update a version"
//	@success	200			{object}	dto.VersionResponse
//	@failure	400			{object}	apperrors.ErrResponse
//	@failure	404			{object}	apperrors.ErrResponse
//	@failure	500			{object}	apperrors.ErrResponse
//	@router		/api/v1/versions/{versionId} [put]
func (c *VersionController) UpdateVersion(w http.ResponseWriter, r *http.Request) {
	version := getVersion(r.Context())

	input := &dto.UpdateVersionRequest{
		Name:        version.Name,
		Description: version.Description,
	}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	versionId, err := httputil.URLParamInt64(r, "versionId")
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	updateParams := &dto.UpdateVersionParams{
		ID:          versionId,
		Name:        input.Name,
		Description: input.Description,
	}
	updateResult, err := c.versionService.UpdateVersion(r.Context(), updateParams)
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	httputil.Render(w, r, dto.ToVersionResponse(updateResult.Version))
}

// DeleteVersion godoc
//
//	@summary	Delete a version
//	@tags		versions
//	@accept		json
//	@param		versionId	path	uint	true	"Version identifier"
//	@success	204
//	@failure	404	{object}	apperrors.ErrResponse
//	@failure	500	{object}	apperrors.ErrResponse
//	@router		/api/v1/versions/{versionId} [delete]
func (c *VersionController) DeleteVersion(w http.ResponseWriter, r *http.Request) {
	version := getVersion(r.Context())

	err := c.versionService.DeleteVersion(r.Context(), &dto.DeleteVersionParams{ID: version.ID})
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
//	@param		versionId	path	uint						true	"Version identifier"
//	@param		request		body	dto.AttachFileToVersionRequest	true	"File to attach"
//	@success	204
//	@failure	400	{object}	apperrors.ErrResponse
//	@failure	404	{object}	apperrors.ErrResponse
//	@failure	500	{object}	apperrors.ErrResponse
//	@router		/api/v1/versions/{versionId}/attach-file [post]
func (c *VersionController) AttachFileToVersion(w http.ResponseWriter, r *http.Request) {
	version := getVersion(r.Context())

	input := &dto.AttachFileToVersionRequest{}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	err := c.versionService.AttachFileToVersion(r.Context(), &dto.AttachFileToVersionParams{VersionID: version.ID, FileID: input.FileId})
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
//	@param		versionId	path	uint							true	"Version identifier"
//	@param		request		body	dto.DetachFileFromVersionRequest	true	"File to detach"
//	@success	204
//	@failure	400	{object}	apperrors.ErrResponse
//	@failure	404	{object}	apperrors.ErrResponse
//	@failure	500	{object}	apperrors.ErrResponse
//	@router		/api/v1/versions/{versionId}/detach-file [post]
func (c *VersionController) DetachFileFromVersion(w http.ResponseWriter, r *http.Request) {
	version := getVersion(r.Context())

	input := &dto.DetachFileFromVersionRequest{}
	if err := render.Bind(r, input); err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPBadRequestError(err))
		return
	}

	err := c.versionService.DetachFileFromVersion(r.Context(), &dto.DetachFileFromVersionParams{VersionID: version.ID, FileID: input.FileId})
	if err != nil {
		httputil.Render(w, r, apperrors.ErrHTTPInternalServerError(err))
		return
	}

	render.NoContent(w, r)
}
