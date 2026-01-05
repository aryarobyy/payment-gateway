package helper

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(c *gin.Context) (int, int, error) {
	limitStr := c.Query("limit")
	limit, limErr := strconv.Atoi(limitStr)
	if limErr != nil {
		return 0, 0, fmt.Errorf("failed getting limit: %w", limErr)
	}
	if limit == 0 {
		limit = 20
	}

	offsetStr := c.Query("offset")
	offset, offsErr := strconv.Atoi(offsetStr)
	if offsErr != nil {
		return 0, 0, fmt.Errorf("failed getting offset: %w", limErr)
	}

	return limit, offset, nil
}
