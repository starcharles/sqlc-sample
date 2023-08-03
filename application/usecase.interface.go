package application

import "context"

type UseCaseOutput any

type IUseCase[O UseCaseOutput] interface {
	Handle(context.Context) (O, error)
}
