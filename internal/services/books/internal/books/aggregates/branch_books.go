package aggregates

import (
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/books/internal/books/entities"
	"github.com/samber/lo"
)

type BranchBooksBuilder interface {
	SetBranch(branch *Branch) BranchBooksBuilder
	BuildBoooksByTimePeriod(startDate time.Time, endDate time.Time) BranchBooksBuilder
	GetDateBooks() []DateBooks
}

func NewBranchBooksBuilder(uuidGenerator uuid.UUIDGenerator) BranchBooksBuilder {
	return &branchBooksBuilder{
		uuidGenerator: uuidGenerator,
		dateBooks:     []DateBooks{},
	}
}

type branchBooksBuilder struct {
	uuidGenerator uuid.UUIDGenerator
	branch        *Branch
	dateBooks     []DateBooks // YYYY-MM-DD,  []Book
}

type DateBooks struct {
	Date  string
	Books []Book
}

func (b *branchBooksBuilder) SetBranch(branch *Branch) BranchBooksBuilder {
	b.branch = branch
	return b
}

func (b *branchBooksBuilder) BuildBoooksByTimePeriod(startDate time.Time, endDate time.Time) BranchBooksBuilder {
	if b.branch == nil {
		return b
	}
	t := time.Now()
	tables := b.getTables()
	avaliableDates := b.getAvailableDateTimeByTimePeriod(startDate, endDate)
	dateBooks := make([]DateBooks, len(avaliableDates))
	for i, avaliableDate := range avaliableDates {
		dateBooks[i] = DateBooks{
			Date:  avaliableDate.Date,
			Books: make([]Book, len(tables)*len(avaliableDate.AvaliableTimes)),
		}
		for j, avaliableTime := range avaliableDate.AvaliableTimes {
			for k, table := range tables {
				dateBooks[i].Books[j*len(tables)+k] = Book{
					ID:        b.uuidGenerator.GenerateUUID(),
					BranchID:  b.branch.ID,
					Date:      avaliableDate.Date,
					StartTime: avaliableTime.StartTime,
					EndTime:   avaliableTime.EndTime,
					Capacity:  table.Capacity,
					CommonCQRSHistoryModel: entities.CommonCQRSHistoryModel{
						ActiveStatus: true,
						CreatedAt:    t,
						UpdatedAt:    t,
					},
				}
			}
		}
	}
	b.dateBooks = append(b.dateBooks, dateBooks...)
	return b
}

func (b *branchBooksBuilder) GetDateBooks() []DateBooks {
	dateBooks := b.dateBooks
	b.dateBooks = []DateBooks{}
	return dateBooks
}

func (b *branchBooksBuilder) getTables() []Table {
	if b.branch == nil {
		return nil
	}
	return lo.Map(b.branch.Tables, func(table Table, _ int) Table {
		return table
	})
}

type availableDateTime struct {
	Date           string
	Weekday        time.Weekday
	AvaliableTimes []avaliableTime
}

type avaliableWeekTime map[time.Weekday][]avaliableTime

type avaliableTime struct {
	StartTime string
	EndTime   string
}

func (b *branchBooksBuilder) getAvailableDateTimeByTimePeriod(startDate time.Time, endDate time.Time) []availableDateTime {
	if b.branch == nil {
		return nil
	}
	// use startDate, endDate  to get all date in between and then use avaliable's weekday to filter date
	awt := b.getAvailableWeekTime()
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, time.UTC)

	adt := []availableDateTime{}
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		if avaliableTimes, ok := awt[d.Weekday()]; ok {
			adt = append(adt, availableDateTime{
				Date:           d.Format(time.DateOnly),
				Weekday:        d.Weekday(),
				AvaliableTimes: avaliableTimes,
			})
		}
	}
	return adt
}

func (b *branchBooksBuilder) getAvailableWeekTime() avaliableWeekTime {
	if b.branch == nil {
		return nil
	}
	awt := make(avaliableWeekTime)
	for _, avaliable := range b.branch.Availables {
		awt[avaliable.Weekday] = append(awt[avaliable.Weekday], avaliableTime{
			StartTime: avaliable.StartTime,
			EndTime:   avaliable.EndTime,
		})
	}
	return awt
}
