# tictacgo

## layering
domain: contains spesific logic for every functional of the domain

usecase: contains set of domain logic that perform bussiness logic/action

transport: contains presentation logic for interact with UI

## app flow
main <> transport <> usecase <> domain

## unit test
domain/tictactoe_test.go