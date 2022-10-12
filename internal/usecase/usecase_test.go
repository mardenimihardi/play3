package usecase

import (
	types "play3/internal/model"
	repository "play3/internal/repository"
	mockRepository "play3/testfile/repository"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUsecase_GetAllDishes(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockRepository := mockRepository.NewMockRepositoryInterface(mockController)

	tests := []struct {
		name string
		u    Usecase
		want []types.Dishes
		mock func()
	}{
		{
			name: "TestUsecase_GetAllDishes: success",
			u: Usecase{
				Repository: mockRepository,
			},
			want: []types.Dishes{{Dish: "Dish"}},
			mock: func() {
				mockRepository.EXPECT().GetAllDishes().Return([]types.Dishes{{Dish: "Dish"}})
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.GetAllDishes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.GetAllDishes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitUsecase(t *testing.T) {
	type args struct {
		repository repository.Repository
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "TestInitUsecase: success",
			args: args{
				repository: repository.Repository{},
			},
			want: Usecase{
				Repository: repository.Repository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitUsecase(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
