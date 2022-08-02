package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/nadirbasalamah/go-products-api/config"
	"github.com/nadirbasalamah/go-products-api/database"
	"github.com/nadirbasalamah/go-products-api/model"
	"github.com/nadirbasalamah/go-products-api/utils"
	"github.com/steinfletcher/apitest"
)

func newApp() *fiber.App {
	app := newFiberApp()

	database.InitDatabase(config.Config("DB_TEST_NAME"))

	return app
}

func getProduct() model.Product {
	database.InitDatabase(config.Config("DB_TEST_NAME"))
	product, err := database.SeedProduct()
	if err != nil {
		panic(err)
	}

	return product
}

func cleanup(res *http.Response, req *http.Request, apiTest *apitest.APITest) {
	if http.StatusOK == res.StatusCode {
		database.CleanSeeders()
	}
}

func getJWTToken(t *testing.T) string {
	database.InitDatabase(config.Config("DB_TEST_NAME"))
	user, err := database.SeedUser()
	if err != nil {
		panic(err)
	}

	var userRequest *model.UserInput = &model.UserInput{
		Email:    user.Email,
		Password: user.Password,
	}

	var resp *http.Response = apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End().Response

	var response map[string]string = map[string]string{}

	json.NewDecoder(resp.Body).Decode(&response)

	var token string = response["token"]

	var JWT_TOKEN = "Bearer " + token

	return JWT_TOKEN
}

func TestRegister_Success(t *testing.T) {
	userData, err := utils.CreateFakerData[model.User]()

	if err != nil {
		panic(err)
	}

	var userRequest *model.UserInput = &model.UserInput{
		Email:    userData.Email,
		Password: userData.Password,
	}

	apitest.New().
		Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestRegister_ValidationFailed(t *testing.T) {
	var userRequest *model.UserInput = &model.UserInput{
		Email:    "",
		Password: "",
	}

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/register").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Success(t *testing.T) {
	database.InitDatabase(config.Config("DB_TEST_NAME"))
	user, err := database.SeedUser()
	if err != nil {
		panic(err)
	}

	var userRequest *model.UserInput = &model.UserInput{
		Email:    user.Email,
		Password: user.Password,
	}

	apitest.New().
		Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestLogin_ValidationFailed(t *testing.T) {
	var userRequest *model.UserInput = &model.UserInput{
		Email:    "",
		Password: "",
	}

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestLogin_Failed(t *testing.T) {
	var userRequest *model.UserInput = &model.UserInput{
		Email:    "notfound@mail.com",
		Password: "123123",
	}

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/login").
		JSON(userRequest).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

func TestGetProducts_Success(t *testing.T) {
	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Get("/api/products").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetProduct_Success(t *testing.T) {
	var product model.Product = getProduct()

	apitest.New().
		Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Get("/api/products/" + product.ID).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestGetProduct_NotFound(t *testing.T) {
	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Get("/api/products/0").
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestCreateProduct_Success(t *testing.T) {
	productData, err := utils.CreateFakerData[model.Product]()
	if err != nil {
		panic(err)
	}

	var productRequest *model.ProductInput = &model.ProductInput{
		Name:        productData.Name,
		Price:       productData.Price,
		Description: productData.Description,
		Stock:       productData.Stock,
	}

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/products").
		Header("Authorization", token).
		JSON(productRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestCreateProduct_ValidationFailed(t *testing.T) {
	var productRequest *model.ProductInput = &model.ProductInput{
		Name:        "",
		Price:       0,
		Stock:       0,
		Description: "",
	}

	var token string = getJWTToken(t)

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Post("/api/products").
		Header("Authorization", token).
		JSON(productRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdateProduct_Success(t *testing.T) {
	var product model.Product = getProduct()

	var productRequest *model.ProductInput = &model.ProductInput{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Stock:       product.Stock,
	}

	var token string = getJWTToken(t)

	apitest.New().
		Observe(cleanup).
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Put("/api/products/"+product.ID).
		Header("Authorization", token).
		JSON(productRequest).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestUpdateProduct_ValidationFailed(t *testing.T) {
	var product model.Product = getProduct()

	var productRequest *model.ProductInput = &model.ProductInput{
		Name:        "",
		Price:       0,
		Stock:       0,
		Description: "",
	}

	var token string = getJWTToken(t)

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Put("/api/products/"+product.ID).
		Header("Authorization", token).
		JSON(productRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestUpdateProduct_Failed(t *testing.T) {
	var productRequest *model.ProductInput = &model.ProductInput{
		Name:        "changed",
		Price:       10,
		Stock:       10,
		Description: "changed",
	}

	var token string = getJWTToken(t)

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Put("/api/products/0").
		Header("Authorization", token).
		JSON(productRequest).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func TestDeleteProduct_Success(t *testing.T) {
	var product model.Product = getProduct()

	var token string = getJWTToken(t)

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Delete("/api/products/"+product.ID).
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeleteProduct_Failed(t *testing.T) {
	var token string = getJWTToken(t)

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Delete("/api/products/0").
		Header("Authorization", token).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}

		// copy headers
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}
