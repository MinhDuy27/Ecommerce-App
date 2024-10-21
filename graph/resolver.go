package graph

import (
	"github.com/MinhDuy27/Ecommerce-App/internal/service"
	
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Usv service.UserService
	Psv service.ProductService
}
