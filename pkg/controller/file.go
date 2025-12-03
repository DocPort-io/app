package controller

import (
	"app/pkg/dto"
	"app/pkg/service"
	"app/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	fileService *service.FileService
}

func NewFileController(fileService *service.FileService) *FileController {
	return &FileController{fileService: fileService}
}

// FindAllFiles godoc
//
//	@summary	Find all files
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		versionId	query		uint	false	"Version identifier"
//	@success	200			{object}	dto.ListFilesResponseDto
//	@router		/files [get]
func (c *FileController) FindAllFiles(ctx *gin.Context) {
	versionId, err := util.GetQueryParameterAsInt64(ctx, "versionId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid versionId parameter"})
		return
	}

	files, err := c.fileService.FindAllFiles(ctx.Request.Context(), versionId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToListFilesResponse(files, int64(len(files))))
}

// GetFile godoc
//
//	@summary	Get a file
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		id	path		uint	true	"File identifier"
//	@success	200	{object}	dto.FileResponseDto
//	@router		/files/{id} [get]
func (c *FileController) GetFile(ctx *gin.Context) {
	id, err := util.GetPathParameterAsInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	file, err := c.fileService.FindFileById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToFileResponse(file))
}

// CreateFile godoc
//
//	@summary	Create a file
//	@tags		files
//	@accept		multipart/form-data
//	@produce	json
//	@param		file	formData	file	true	"File to upload"
//	@success	201		{object}	dto.FileResponseDto
//	@router		/files [post]
func (c *FileController) CreateFile(ctx *gin.Context) {
	multipartFile, err := util.SaveMultipartFileToTemp(ctx, "file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	createFileDto := dto.ToCreateFileDto(multipartFile)

	file, err := c.fileService.CreateFile(ctx.Request.Context(), *createFileDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ToFileResponse(file))
}

// DeleteFile godoc
//
//	@summary	Delete a file
//	@tags		files
//	@accept		json
//	@produce	json
//	@param		id	path	uint	true	"File identifier"
//	@success	204
//	@router		/files/{id} [delete]
func (c *FileController) DeleteFile(ctx *gin.Context) {
	id, err := util.GetPathParameterAsInt64(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	err = c.fileService.DeleteFile(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
