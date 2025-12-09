package controller

import (
	"app/pkg/dto"
	"app/pkg/service"
	"app/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
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
//	@param		projectId	query		uint	false	"Project identifier"
//	@success	200			{object}	dto.ListVersionsResponseDto
//	@router		/versions [get]
func (c *VersionController) FindAllVersions(ctx *gin.Context) {
	projectId, err := util.GetQueryParameterAsInt64(ctx, "projectId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	versions, err := c.versionService.FindAllVersions(ctx.Request.Context(), projectId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToListVersionsResponse(versions, int64(len(versions))))
}

// GetVersion godoc
//
//	@summary	Get a version
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		id	path		uint	true	"Version identifier"
//	@success	200	{object}	dto.VersionResponseDto
//	@router		/versions/{id} [get]
func (c *VersionController) GetVersion(ctx *gin.Context) {
	id, err := util.GetPathParameterAsInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	version, err := c.versionService.FindVersionById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToVersionResponse(version))
}

// CreateVersion godoc
//
//	@summary	Create a version
//	@tags		versions
//	@accept		json
//	@produce	json
//	@param		request	body		dto.CreateVersionDto	true	"Create a version"
//	@success	201		{object}	dto.VersionResponseDto
//	@router		/versions [post]
func (c *VersionController) CreateVersion(ctx *gin.Context) {
	var input dto.CreateVersionDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	version, err := c.versionService.CreateVersion(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ToVersionResponse(version))
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
//	@router		/versions/{id} [put]
func (c *VersionController) UpdateVersion(ctx *gin.Context) {
	var input dto.UpdateVersionDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := util.GetPathParameterAsInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	version, err := c.versionService.UpdateVersion(ctx.Request.Context(), id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToVersionResponse(version))
}

// DeleteVersion godoc
//
//	@summary	Delete a version
//	@tags		versions
//	@accept		json
//	@param		id	path	uint	true	"Version identifier"
//	@success	204
//	@router		/versions/{id} [delete]
func (c *VersionController) DeleteVersion(ctx *gin.Context) {
	id, err := util.GetPathParameterAsInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	err = c.versionService.DeleteVersion(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// AttachFileToVersion godoc
//
//	@summary	Attaches a file to a version
//	@tags		versions
//	@accept		json
//	@param		id		path	uint						true	"Version identifier"
//	@param		request	body	dto.AttachFileToVersionDto	true	"File to attach"
//	@success	204
//	@router		/versions/{id}/attach-file [post]
func (c *VersionController) AttachFileToVersion(ctx *gin.Context) {
	id, err := util.GetPathParameterAsInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	var input dto.AttachFileToVersionDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.versionService.AttachFileToVersion(ctx.Request.Context(), id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// DetachFileFromVersion godoc
//
//	@summary	Detach a file from a version
//	@tags		versions
//	@accept		json
//	@param		id		path	uint							true	"Version identifier"
//	@param		request	body	dto.DetachFileFromVersionDto	true	"File to detach"
//	@success	204
//	@router		/versions/{id}/detach-file [post]
func (c *VersionController) DetachFileFromVersion(ctx *gin.Context) {
	id, err := util.GetPathParameterAsInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	var input dto.DetachFileFromVersionDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = c.versionService.DetachFileFromVersion(ctx.Request.Context(), id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
