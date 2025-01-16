package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "github.com/username/go_rest_api_crud/model"
    "github.com/username/go_rest_api_crud/service"
)

type ProductHandler struct {
    productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
    return &ProductHandler{
        productService: productService,
    }
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.productService.CreateProduct(&product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
    products, err := h.productService.GetAllProducts()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
    product, err := h.productService.GetProduct(uint(id))
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
    var product models.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    product.ID = uint(id)
    if err := h.productService.UpdateProduct(&product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
    if err := h.productService.DeleteProduct(uint(id)); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
		
        return
    }

    w.WriteHeader(http.StatusNoContent)
}