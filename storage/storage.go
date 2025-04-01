package storage

import "Sinekod/models"

var Books = make(map[int]models.Book)
var Users = make(map[int]models.User)
var IdUser int = 0
var IdBook int = 0