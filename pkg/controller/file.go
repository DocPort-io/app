package controller

import (
	"app/pkg/dto"
	"app/pkg/service"
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
	versionId := ctx.Query("versionId")

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
//	@param		id	path	uint	true	"File identifier"
//	@success	200	{object}	dto.FileResponseDto
//	@router		/files/{id} [get]
func (c *FileController) GetFile(ctx *gin.Context) {
	id := ctx.Param("id")

	file, err := c.fileService.FindFileById(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto.ToFileResponse(file))
}
