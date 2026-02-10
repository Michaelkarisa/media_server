func main() {
    db := db.Connect()
    r := gin.Default()
    media.RegisterRoutes(r, db)
    r.Run(":8080")
}
