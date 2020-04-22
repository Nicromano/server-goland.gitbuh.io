package model

import "time"
/* Estructura de usuario */
type User struct {
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Email     string         `json:"email"`
	Image     string         `json:"image"`
	Links     []Link        `json:"links"`
	Followers []string `json:"followers"`
	Follow    []string  `json:"follow"`
}

/* Estructura de link */
type Link struct {
	Name        string     `json:"name"`
	Url         string     `json:"url"`
	Description string     `json:"description"`
	Comments    []Comment `json:"comments"`
	Like        uint32     `json:"like"`
	Timestamp 	time.Time  `json:"timestamp"`
}

/* Estructura para comentario */
type Comment struct {
	IdUser  string `json:"iduser"`
	Content string `json:"content"`
	Like    uint32 `json:"like"`
}
