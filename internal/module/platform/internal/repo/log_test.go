package repo

import (
	"context"
	"fmt"
	"testing"
	"time"

	logEntity "github.com/perfect-panel/server/internal/module/platform/entity/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestDeleteBeforePreservesFinancialLedgers(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("file:system-log-retention-%d?mode=memory&cache=shared", time.Now().UnixNano())), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&logEntity.SystemLog{}); err != nil {
		t.Fatal(err)
	}

	oldDate := "2025-01-01"
	rows := []*logEntity.SystemLog{
		{Type: logEntity.TypeEmailMessage.Uint8(), Date: oldDate, Content: `{}`},
		{Type: logEntity.TypeLogin.Uint8(), Date: oldDate, Content: `{}`},
		{Type: logEntity.TypeBalance.Uint8(), Date: oldDate, Content: `{}`},
		{Type: logEntity.TypeCommission.Uint8(), Date: oldDate, Content: `{}`},
		{Type: logEntity.TypeGift.Uint8(), Date: oldDate, Content: `{}`},
		{Type: 99, Date: oldDate, Content: `{}`},
	}
	if err := db.Create(&rows).Error; err != nil {
		t.Fatal(err)
	}

	repo := NewLogRepo(db)
	if err := repo.DeleteBefore(context.Background(), time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Fatal(err)
	}

	var remaining []logEntity.SystemLog
	if err := db.Order("type ASC").Find(&remaining).Error; err != nil {
		t.Fatal(err)
	}
	want := []uint8{logEntity.TypeBalance.Uint8(), logEntity.TypeCommission.Uint8(), logEntity.TypeGift.Uint8(), 99}
	if len(remaining) != len(want) {
		t.Fatalf("remaining types = %v, want %v", logTypes(remaining), want)
	}
	for i, typ := range want {
		if remaining[i].Type != typ {
			t.Fatalf("remaining types = %v, want %v", logTypes(remaining), want)
		}
	}
}

func logTypes(rows []logEntity.SystemLog) []uint8 {
	result := make([]uint8, 0, len(rows))
	for _, row := range rows {
		result = append(result, row.Type)
	}
	return result
}
