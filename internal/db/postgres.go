func Connect() *gorm.DB {
    dsn := "host=localhost user=postgres password=postgres dbname=media port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&media.Media{})
    return db
}
