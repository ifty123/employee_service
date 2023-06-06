package repository

import (
	"context"
	"latihan_service/internal/model"
	pkgdto "latihan_service/pkg/dto"
	"strings"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Employee, *pkgdto.PaginationInfo, error)
}

type Employee struct {
	Db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *Employee {
	return &Employee{
		Db: db,
	}
}

func (e *Employee) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Employee, *pkgdto.PaginationInfo, error) {
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
