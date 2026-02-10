type Media struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Title     string    `json:"title"`
    URL       string    `json:"url"`
    Duration  int       `json:"duration"`
    CreatedAt time.Time `json:"created_at"`
}
