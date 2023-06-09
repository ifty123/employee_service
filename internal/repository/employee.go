package repository

import (
	"context"
	"latihan_service/internal/dto"
	"latihan_service/internal/model"
	pkgdto "latihan_service/pkg/dto"
	"strings"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Employee, *pkgdto.PaginationInfo, error)
	FindById(ctx context.Context, id uint, usePreload bool) (*model.Employee, error)
	FindByEmail(ctx context.Context, email *string) (*model.Employee, error)
	ExistByEmail(ctx context.Context, email string) (bool, error)
	Save(ctx context.Context, employee *dto.RegisterEmployeeReq) (*model.Employee, error)
}

type employee struct {
	Db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *employee {
	return &employee{
		Db: db,
	}
}

func (e *employee) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Employee, *pkgdto.PaginationInfo, error) {
	var users []model.Employee
	var count int64

	query := e.Db.WithContext(ctx).Model(&model.Employee{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(fullname) LIKE ? or lower(email) Like ? ", search, search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&users).Error

	return users, pkgdto.CheckInfoPagination(pagination, count), err
}

func (e *employee) FindById(ctx context.Context, id uint, usePreload bool) (*model.Employee, error) {
	var emp *model.Employee
	q := e.Db.WithContext(ctx).Model(&model.Employee{}).Where("id = ?", id)

	if usePreload {
		q = q.Preload("Division")
	}

	err := q.First(emp).Error
	return emp, err
}

func (e *employee) FindByEmail(ctx context.Context, email *string) (*model.Employee, error) {
	var emp model.Employee

	err := e.Db.WithContext(ctx).Where("email = ?", email).First(&emp).Error

	if err != nil {
		return nil, err
	}

	return &emp, err
}

func (e *employee) ExistByEmail(ctx context.Context, email string) (bool, error) {
	var (
		count   int64
		isExist bool
	)

	if err := e.Db.WithContext(ctx).Model(&model.Employee{}).Where("email = ?", &email).Count(&count).Error; err != nil {
		return isExist, nil
	}

	if count > 0 {
		isExist = true
	}

	return isExist, nil
}

func (e *employee) Save(ctx context.Context, employee *dto.RegisterEmployeeReq) (*model.Employee, error) {
	newEmployee := model.Employee{
		Fullname:   employee.Fullname,
		Email:      employee.Email,
		Password:   employee.Password,
		DivisionId: *employee.DivisionID,
		Role:       employee.Role,
	}

	if err := e.Db.WithContext(ctx).Save(&newEmployee).Error; err != nil {
		return &newEmployee, err
	}

	return &newEmployee, nil
}
