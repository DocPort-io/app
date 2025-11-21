package controller

import (
	"app/pkg/dto"
	"app/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VersionController struct {
	db *gorm.DB
}

func NewVersionController(db *gorm.DB) *VersionController {
	return &VersionController{db: db}
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
func (c *VersionController) FindAllVersions(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()

	projectId := ginCtx.Query("projectId")

	var versions []model.Version
	var err error

	if projectId != "" {
		versions, err = gorm.G[model.Version](c.db).Where("project_id = ?", projectId).Find(ctx)
	} else {
		versions, err = gorm.G[model.Version](c.db).Find(ctx)
	}

	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, dto.ToListVersionsResponse(versions, int64(len(versions))))
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
func (c *VersionController) GetVersion(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()

	id := ginCtx.Param("id")

	version, err := gorm.G[model.Version](c.db).Where("id = ?", id).First(ctx)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, dto.ToVersionResponse(&version))
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
	var versionDto dto.CreateVersionDto
	if err := ginCtx.ShouldBindJSON(&versionDto); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := ginCtx.Request.Context()

	version := versionDto.ToModel()

	err := gorm.G[model.Version](c.db).Create(ctx, version)
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
func (c *VersionController) UpdateVersion(ginCtx *gin.Context) {
	var updateVersionDto dto.UpdateVersionDto
	if err := ginCtx.ShouldBindJSON(&updateVersionDto); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := ginCtx.Request.Context()

	id := ginCtx.Param("id")
	updateVersion := updateVersionDto.ToModel()

	rowsAffected, err := gorm.G[model.Version](c.db).Where("id = ?", id).Updates(ctx, *updateVersion)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		ginCtx.JSON(http.StatusNotFound, gin.H{
			"error": "version not found",
		})
		return
	}

	updatedVersion, err := gorm.G[model.Version](c.db).Where("id = ?", id).First(ctx)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, dto.ToVersionResponse(&updatedVersion))
}

// DeleteVersion godoc
//
//	@summary	Delete a version
//	@tags		versions
//	@accept		json
//	@param		id	path	uint	true	"Version identifier"
//	@success	204
//	@router		/versions/{id} [delete]
func (c *VersionController) DeleteVersion(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()

	id := ginCtx.Param("id")

	rowsAffected, err := gorm.G[model.Version](c.db).Where("id = ?", id).Delete(ctx)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		ginCtx.JSON(http.StatusNotFound, gin.H{
			"error": "version not found",
		})
		return
	}

	ginCtx.JSON(http.StatusNoContent, nil)
}
