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
