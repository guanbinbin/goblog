package models

import "time"

type Category struct {
	Id 				int
	Name 			string
	RouterLink 		string
	LinkArticle 	int
	Status 			int
	Created 		time.Time
	Updated 		time.Time
}
