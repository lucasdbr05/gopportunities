package schemas

import (
  "gorm.io/gorm"
)

type Product struct {
  gorm.Model
  Role string
  Company string
  Location string
  Remote bool
  Link string
  Salary float64
}


