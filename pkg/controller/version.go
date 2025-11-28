package controller

import (
	"app/pkg/dto"
	"app/pkg/service"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	projectId := ctx.Query("projectId")

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
	id := ctx.Param("id")

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
func (c *VersionController) CreateVersion(ginCtx *gin.Context) {
	var input dto.CreateVersionDto
	if err := ginCtx.ShouldBindJSON(&input); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	version, err := c.versionService.CreateVersion(ginCtx.Request.Context(), input)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusCreated, dto.ToVersionResponse(version))
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

	id := ctx.Param("id")

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
	id := ctx.Param("id")

	err := c.versionService.DeleteVersion(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// UploadFileToVersion godoc
//
//	@summary	Upload a file to a version
//	@tags		versions
//	@accept		multipart/form-data
//	@produce	json
//	@param		id		path		uint	true	"Version identifier"
//	@param		file	formData	file	true	"File to upload"
//	@success	201		{object}	dto.FileResponseDto
//	@router		/versions/{id}/upload [post]
func (c *VersionController) UploadFileToVersion(ctx *gin.Context) {
	id := ctx.Param("id")

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	tempDir, err := os.MkdirTemp("", "upload")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var filePath = path.Join(tempDir, uuid.NewString())

	err = ctx.SaveUploadedFile(fileHeader, filePath, 0700)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	createFileDto := dto.CreateFileDto{
		Name: fileHeader.Filename,
		Size: fileHeader.Size,
		Path: filePath,
	}

	file, err := c.versionService.UploadFileToVersion(ctx.Request.Context(), id, createFileDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ToFileResponse(file))
}
