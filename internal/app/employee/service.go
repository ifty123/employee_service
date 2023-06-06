package employee

import (
	"context"
	dto_internal "latihan_service/internal/dto"
	"latihan_service/internal/factory"
	"latihan_service/internal/repository"
	"latihan_service/pkg/dto"
	pkgdto "latihan_service/pkg/dto"
	"latihan_service/pkg/util/response"
)

type Service struct {
	EmployeeRepository repository.EmployeeRepository
}

type EmployeeService interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto_internal.EmployeeResponse], error)
}

func Newservice(f *factory.Factory) Service {
	return Service{
		EmployeeRepository: f.EmployeeRepository,
	}
}

func (s *Service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto_internal.EmployeeResponse], error) {
	employee, paginate, err := s.EmployeeRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	var data []dto_internal.EmployeeResponse

	for _, e := range employee {
		data = append(data, dto_internal.EmployeeResponse{
			ID:       e.ID,
			Fullname: e.Fullname,
			Email:    e.Email,
		})
	}

	res := new(dto.SearchGetResponse[dto_internal.EmployeeResponse])

	res.Data = data
	res.PaginationInfo = *paginate

	return res, nil
}
