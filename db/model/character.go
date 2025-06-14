/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package model

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	AccountID uint
	Name      string `gorm:"unique;uniqueIndex:idx_name"`

	Face   int
	Hair   int
	Top    int
	Bottom int
	Shoes  int
	Weapon int

	Str uint16 `gorm:"default:4"`
	Dex uint16 `gorm:"default:4"`
	Int uint16 `gorm:"default:4"`
	Luk uint16 `gorm:"default:4"`

	MaxHp uint16 `gorm:"default:50"`
	MaxMp uint16 `gorm:"default:25"`
	Hp    uint16 `gorm:"default:50"`
	Mp    uint16 `gorm:"default:25"`
	Exp   uint32 `gorm:"default:0"`

	Level int `gorm:"default:1"`
	Job   int `gorm:"default:0"`
	Map   int `gorm:"default:1"`
}
