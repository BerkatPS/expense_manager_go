package monthly_reports

import (
	"github.com/BerkatPS/internal/domain/model"
	"gorm.io/gorm"
	"time"
)

type MonthlyReportRepository interface {
	Create(report *model.MonthlyReport) (*model.MonthlyReport, error)
	FindAll() ([]*model.MonthlyReport, error)
	UpdateMonthlyReport(accountID int, amount float64, transactionType int) error
	FindByMonthAndYear(month, year int) ([]*model.MonthlyReport, error)
}

type monthlyReportRepository struct {
	db *gorm.DB
}

func (r *monthlyReportRepository) UpdateMonthlyReport(accountID int, amount float64, transactionType int) error {
	var reports []*model.MonthlyReport

	currentMonth := time.Now().Month()
	currentYear := time.Now().Year()

	err := r.db.Where("account_id = ? AND report_month = ? AND report_year = ?", accountID, currentMonth, currentYear).Find(&reports).Error

	if err != nil {

		report := model.MonthlyReport{
			AccountID:   accountID,
			ReportMonth: int(currentMonth),
			ReportYear:  currentYear,
		}

		if transactionType == 1 {
			report.TotalIncome = amount
		} else if transactionType == 2 {
			report.TotalExpense = amount
		}

		err = r.db.Create(&report).Error
	} else {
		if transactionType == 1 {
			reports[0].TotalIncome = reports[0].TotalIncome + amount
		} else if transactionType == 2 {
			reports[0].TotalExpense = reports[0].TotalExpense + amount
		}
		err = r.db.Save(&reports[0]).Error
	}
	return err
}

func NewMonthlyReportRepository(db *gorm.DB) MonthlyReportRepository {
	return &monthlyReportRepository{db: db}
}

func (r *monthlyReportRepository) Create(report *model.MonthlyReport) (*model.MonthlyReport, error) {
	err := r.db.Create(&report).Error
	return report, err
}
func (r *monthlyReportRepository) FindByMonthAndYear(month, year int) ([]*model.MonthlyReport, error) {
	var reports []*model.MonthlyReport
	err := r.db.Where("report_month = ? AND report_year = ?", month, year).Find(&reports).Error
	return reports, err
}

func (r *monthlyReportRepository) FindAll() ([]*model.MonthlyReport, error) {
	var reports []*model.MonthlyReport
	err := r.db.Find(&reports).Error
	return reports, err
}
