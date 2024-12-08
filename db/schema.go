package db

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID  `gorm:"primaryKey" json:"id"`
	UserName   string     `gorm:"not null" json:"user_name"`
	Password   string     `json:"password"`
	Email      string     `gorm:"unique;not null" json:"email"`
	Tag        string     `gorm:"unique;index" json:"tag"`
	Bio        string     `json:"bio"`
	Image_ID   *uuid.UUID `gorm:"default:null" json:"image_id"`
	Banner_ID  *uuid.UUID `gorm:"default:null" json:"banner_id"`
	GitHub_ID  *uuid.UUID `gorm:"unique" json:"github_id"`
	Google_ID  *uuid.UUID `gorm:"unique" json:"google_id"`
	Discord_ID *uuid.UUID `gorm:"unique" json:"discord_id"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Post struct {
	ID             uuid.UUID  `gorm:"primaryKey" json:"id"`
	Author_ID      uuid.UUID  `gorm:"not null;foreignKey:Author_ID;references:ID" json:"author_id"`
	Likes_count    int        `gorm:"type:integer;default:0" json:"likes_count"`
	Saved_count    int        `gorm:"type:integer;default:0" json:"saves_count"`
	Comments_count int        `gorm:"type:integer;default:0" json:"comments_count"`
	CreatedAt      time.Time  `gorm:"autoCreateTime"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime"`
	DeletedBy      *uuid.UUID `gorm:"default:null" json:"deleted_by"`
	DeletedAt      time.Time  `json:"deleted_at"`
	Target_ID      *uuid.UUID `gorm:"default:null" json:"target_id"`
	Content        string     `json:"content"`
}

type Likes struct {
	User_ID uuid.UUID `gorm:"primaryKey;foreignKey:User_ID;references:ID" json:"user_id"`
	Post_ID uuid.UUID `gorm:"primaryKey;foreignKey:Post_ID;references:ID" json:"post_id"`
}

type Saves struct {
	User_ID uuid.UUID `gorm:"primaryKey;foreignKey:User_ID;references:ID" json:"user_id"`
	Post_ID uuid.UUID `gorm:"primaryKey;foreignKey:Post_ID;references:ID" json:"post_id"`
}

type Comments struct {
	User_ID  string `gorm:"primaryKey;foreignKey:User_ID;references:ID" json:"user_id"`
	Post_ID  string `gorm:"primaryKey;foreignKey:Post_ID;references:ID" json:"post_id"`
	Content  string `json:"content"`
	Image_ID string `json:"image_id"`
}

type Media struct {
	ID        uuid.UUID  `gorm:"primaryKey" json:"id"`
	Target_ID *uuid.UUID `gorm:"default:null" json:"post_id"`
	Image     string     `json:"image"`
	Type      string     `gorm:"default:image" json:"type"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

type Relation struct {
	User_ID   uuid.UUID `gorm:"primaryKey;foreignKey:User_ID;references:ID" json:"user_id"`
	Target_ID uuid.UUID `gorm:"primaryKey;foreignKey:Target_ID;references:ID" json:"target_id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Badge struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Type      string    `gorm:"default:image" json:"type"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Notification struct {
	ID        uuid.UUID  `gorm:"primaryKey" json:"id"`
	User_ID   uuid.UUID  `gorm:"not null;foreignKey:User_ID;references:ID" json:"user_id"`
	Post_ID   *uuid.UUID `gorm:"foreignKey:Post_ID;references:ID;default:null" json:"post_id"`
	Target_ID uuid.UUID  `gorm:"not null;foreignKey:Target_ID;references:ID" json:"target_id"`
	Type      string     `json:"type"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
}

type UserBadge struct {
	User_ID  uuid.UUID `gorm:"primaryKey;foreignKey:User_ID;references:ID" json:"user_id"`
	Badge_ID uuid.UUID `gorm:"primaryKey;foreignKey:Badge_ID;references:ID" json:"badge_id"`
}

type Community struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Type      string    `gorm:"default:image" json:"type"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type CommunityMember struct {
	User_ID      uuid.UUID `gorm:"primaryKey;foreignKey:User_ID;references:ID" json:"user_id"`
	Community_ID uuid.UUID `gorm:"primaryKey;foreignKey:Community_ID;references:ID" json:"community"`
	Role         string    `gorm:"default:member" json:"role"`
}
