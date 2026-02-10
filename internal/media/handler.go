func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
    mediaGroup := r.Group("/media")
    {
        mediaGroup.GET("/", func(c *gin.Context) {
            // List media
        })
        mediaGroup.POST("/", func(c *gin.Context) {
            // Create media
        })
        mediaGroup.GET("/:id", func(c *gin.Context) {
            // Get by ID
        })
        mediaGroup.DELETE("/:id", func(c *gin.Context) {
            // Delete
        })
    }
}
