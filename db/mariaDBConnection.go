/*
 * MIT License
 *
 * Copyright (c) 2025 Junghee Wang
 */

package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	MariaDB *gorm.DB
)

// ConnectMariaDB initializes and connects to the MariaDB database
func ConnectMariaDB(username, password, host, dbName string, port int) {
	// DSN (Data Source Name) 구성
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbName,
	)

	// GORM 설정
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // SQL 쿼리 로깅
	})
	if err != nil {
		log.Fatalf("[DB] Failed to connect to MariaDB: %v", err)
	}

	// 데이터베이스 연결 테스트
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("[DB] Failed to get SQL DB instance: %v", err)
	}

	// CP
	sqlDB.SetMaxOpenConns(100) // 최대 연결 수
	sqlDB.SetMaxIdleConns(10)  // 유휴 상태 최대 연결 수
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	// Ping 테스트
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("[DB] MariaDB Doesn't respond to ping: %v", err)
	}

	log.Println("[DB] Successfully connected to MariaDB!")

	// DB 전역 변수 설정
	MariaDB = db
}
