func CreateMedia(db *gorm.DB, media *Media) error {
    return db.Create(media).Error
}

func GetMediaByID(db *gorm.DB, id uint) (*Media, error) {
    var m Media
    if err := db.First(&m, id).Error; err != nil {
        return nil, err
    }
    return &m, nil
}
