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
//	@success	200	{object}	dto.ListVersionsResponseDto
//	@router		/versions [get]
func (c *VersionController) FindAllVersions(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()

	versions, err := gorm.G[model.Version](c.db).Find(ctx)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, dto.ToListVersionsResponse(versions, int64(len(versions))))
}

// FindAllProjectVersions godoc
//
//	@summary	Find all versions for a project
//	@tags		projects,versions
//	@accept		json
//	@produce	json
//	@param		id	path		string	true	"Project ID"
//	@success	200	{object}	dto.ListVersionsResponseDto
//	@router		/projects/{id}/versions [get]
func (c *VersionController) FindAllProjectVersions(ginCtx *gin.Context) {
	type Params struct {
		ProjectID string `uri:"id" binding:"required"`
	}

	var params Params
	if err := ginCtx.ShouldBindUri(&params); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := ginCtx.Request.Context()

	versions, err := gorm.G[model.Version](c.db).Where("project_id = ?", params.ProjectID).Find(ctx)
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
//	@param		id	path		uint	true	"Version ID"
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
