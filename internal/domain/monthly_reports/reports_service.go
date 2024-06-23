package monthly_reports

import "github.com/BerkatPS/internal/domain/model"

type MonthlyReportService interface {
	CreateMonthlyReport(report *model.MonthlyReport) (*model.MonthlyReport, error)
	GetAllMonthlyReports() ([]*model.MonthlyReport, error)
	UpdateMonthlyReport(accountID int, amount float64, transactionType int) error
	GetMonthlyReportsByMonthAndYear(month, year int) ([]*model.MonthlyReport, error)
}

type monthlyReportService struct {
	repository MonthlyReportRepository
}

func (s *monthlyReportService) GetMonthlyReportsByMonthAndYear(month, year int) ([]*model.MonthlyReport, error) {
	return s.repository.FindByMonthAndYear(month, year)
}

func NewMonthlyReportService(repo MonthlyReportRepository) MonthlyReportService {
	return &monthlyReportService{repository: repo}
}

func (s *monthlyReportService) CreateMonthlyReport(report *model.MonthlyReport) (*model.MonthlyReport, error) {
	return s.repository.Create(report)
}

func (s *monthlyReportService) UpdateMonthlyReport(accountID int, amount float64, transactionType int) error {
	return s.repository.UpdateMonthlyReport(accountID, amount, transactionType)
}
func (s *monthlyReportService) GetAllMonthlyReports() ([]*model.MonthlyReport, error) {
	return s.repository.FindAll()
}
