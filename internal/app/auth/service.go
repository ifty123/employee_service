package auth

import (
	"context"
	"errors"
	"latihan_service/internal/dto"
	"latihan_service/internal/factory"
	"latihan_service/internal/repository"
	"latihan_service/pkg/constant"
	"latihan_service/pkg/util"
	"latihan_service/pkg/util/response"
	"log"
)

type Service struct {
	EmployeeRepository repository.EmployeeRepository
}

type AuthService interface {
	LoginByEmailAndPassword(ctx context.Context, payload *dto.EmailAndPasswordReq) (*dto.EmployeeWithJWTResponse, error)
	RegisterByEmailAndPassword(ctx context.Context, payload *dto.RegisterEmployeeReq) (*dto.EmployeeResponse, error)
}

func NewService(f *factory.Factory) AuthService {
	return &Service{
		EmployeeRepository: f.EmployeeRepository,
	}
}

func (s *Service) LoginByEmailAndPassword(ctx context.Context, payload *dto.EmailAndPasswordReq) (*dto.EmployeeWithJWTResponse, error) {
	var res *dto.EmployeeWithJWTResponse

	data, err := s.EmployeeRepository.FindByEmail(ctx, &payload.Email)
	if err != nil {
		log.Println("err find email :", err)
		if err == constant.RECORD_NOT_FOUND {
			return res, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	if !(util.CompareHashPassword(payload.Password, data.Password)) {
		return res, response.ErrorBuilder(
			&response.ErrorConstant.EmailOrPasswordIncorrect, errors.New(response.ErrorConstant.EmailOrPasswordIncorrect.Response.Message),
		)
	}

	claims := util.CreateJWTClaims(data.Email, data.ID, data.DivisionId)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		log.Println("err generate :", err)
		return res, response.ErrorBuilder(
			&response.ErrorConstant.InternalServerError,
			errors.New("Error when generate token"),
		)
	}

	res = &dto.EmployeeWithJWTResponse{
		EmployeeResponse: dto.EmployeeResponse{
			ID:       data.ID,
			Fullname: data.Fullname,
			Email:    data.Email,
		},
		JWT: token,
	}

	return res, nil
}

func (s *Service) RegisterByEmailAndPassword(ctx context.Context, payload *dto.RegisterEmployeeReq) (*dto.EmployeeResponse, error) {
	var res *dto.EmployeeResponse

	isExist, err := s.EmployeeRepository.ExistByEmail(ctx, payload.Email)
	if err != nil {
		return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	if isExist {
		return res, response.ErrorBuilder(&response.ErrorConstant.Duplicate, errors.New("Employee already exist"))
	}

	hashPw, err := util.HashPassword(payload.Password)
	if err != nil {
		return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	payload.Password = hashPw

	data, err := s.EmployeeRepository.Save(ctx, payload)
	if err != nil {
		return res, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	res = &dto.EmployeeResponse{
		ID:       data.ID,
		Fullname: data.Fullname,
		Email:    data.Email,
	}

	return res, nil
}
