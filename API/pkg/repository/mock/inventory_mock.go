// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interfaces/Inventory.go

// Package mock is a generated GoMock package.
package mock

import (
	domain "CrocsClub/pkg/domain"
	models "CrocsClub/pkg/utils/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInventoryRepository is a mock of InventoryRepository interface.
type MockInventoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryRepositoryMockRecorder
}

// MockInventoryRepositoryMockRecorder is the mock recorder for MockInventoryRepository.
type MockInventoryRepositoryMockRecorder struct {
	mock *MockInventoryRepository
}

// NewMockInventoryRepository creates a new mock instance.
func NewMockInventoryRepository(ctrl *gomock.Controller) *MockInventoryRepository {
	mock := &MockInventoryRepository{ctrl: ctrl}
	mock.recorder = &MockInventoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInventoryRepository) EXPECT() *MockInventoryRepositoryMockRecorder {
	return m.recorder
}

// AddInventory mocks base method.
func (m *MockInventoryRepository) AddInventory(inventory models.AddInventories, url string) (models.ProductsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInventory", inventory, url)
	ret0, _ := ret[0].(models.ProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddInventory indicates an expected call of AddInventory.
func (mr *MockInventoryRepositoryMockRecorder) AddInventory(inventory, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInventory", reflect.TypeOf((*MockInventoryRepository)(nil).AddInventory), inventory, url)
}

// CheckInventory mocks base method.
func (m *MockInventoryRepository) CheckInventory(pid int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckInventory", pid)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckInventory indicates an expected call of CheckInventory.
func (mr *MockInventoryRepositoryMockRecorder) CheckInventory(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckInventory", reflect.TypeOf((*MockInventoryRepository)(nil).CheckInventory), pid)
}

// CheckInventoryByCatAndName mocks base method.
func (m *MockInventoryRepository) CheckInventoryByCatAndName(cat int, prdct string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckInventoryByCatAndName", cat, prdct)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckInventoryByCatAndName indicates an expected call of CheckInventoryByCatAndName.
func (mr *MockInventoryRepositoryMockRecorder) CheckInventoryByCatAndName(cat, prdct interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckInventoryByCatAndName", reflect.TypeOf((*MockInventoryRepository)(nil).CheckInventoryByCatAndName), cat, prdct)
}

// CheckStock mocks base method.
func (m *MockInventoryRepository) CheckStock(inventory_id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckStock", inventory_id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckStock indicates an expected call of CheckStock.
func (mr *MockInventoryRepositoryMockRecorder) CheckStock(inventory_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckStock", reflect.TypeOf((*MockInventoryRepository)(nil).CheckStock), inventory_id)
}

// DeleteInventory mocks base method.
func (m *MockInventoryRepository) DeleteInventory(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInventory", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInventory indicates an expected call of DeleteInventory.
func (mr *MockInventoryRepositoryMockRecorder) DeleteInventory(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInventory", reflect.TypeOf((*MockInventoryRepository)(nil).DeleteInventory), id)
}

// EditInventory mocks base method.
func (m *MockInventoryRepository) EditInventory(arg0 domain.Inventories, arg1 int) (domain.Inventories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditInventory", arg0, arg1)
	ret0, _ := ret[0].(domain.Inventories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditInventory indicates an expected call of EditInventory.
func (mr *MockInventoryRepositoryMockRecorder) EditInventory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditInventory", reflect.TypeOf((*MockInventoryRepository)(nil).EditInventory), arg0, arg1)
}

// FetchProductDetails mocks base method.
func (m *MockInventoryRepository) FetchProductDetails(productId uint) (models.Inventories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchProductDetails", productId)
	ret0, _ := ret[0].(models.Inventories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchProductDetails indicates an expected call of FetchProductDetails.
func (mr *MockInventoryRepositoryMockRecorder) FetchProductDetails(productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchProductDetails", reflect.TypeOf((*MockInventoryRepository)(nil).FetchProductDetails), productId)
}

// FilterByCategory mocks base method.
func (m *MockInventoryRepository) FilterByCategory(CategoryIdInt int) ([]models.ProductsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterByCategory", CategoryIdInt)
	ret0, _ := ret[0].([]models.ProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterByCategory indicates an expected call of FilterByCategory.
func (mr *MockInventoryRepositoryMockRecorder) FilterByCategory(CategoryIdInt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterByCategory", reflect.TypeOf((*MockInventoryRepository)(nil).FilterByCategory), CategoryIdInt)
}

// GetInventory mocks base method.
func (m *MockInventoryRepository) GetInventory(prefix string) ([]models.ProductsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInventory", prefix)
	ret0, _ := ret[0].([]models.ProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInventory indicates an expected call of GetInventory.
func (mr *MockInventoryRepositoryMockRecorder) GetInventory(prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInventory", reflect.TypeOf((*MockInventoryRepository)(nil).GetInventory), prefix)
}

// ImageUploader mocks base method.
func (m *MockInventoryRepository) ImageUploader(inventoryID int, url string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImageUploader", inventoryID, url)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImageUploader indicates an expected call of ImageUploader.
func (mr *MockInventoryRepositoryMockRecorder) ImageUploader(inventoryID, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageUploader", reflect.TypeOf((*MockInventoryRepository)(nil).ImageUploader), inventoryID, url)
}

// ListProducts mocks base method.
func (m *MockInventoryRepository) ListProducts(arg0, arg1 int) ([]models.ProductsResponseDisp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0, arg1)
	ret0, _ := ret[0].([]models.ProductsResponseDisp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockInventoryRepositoryMockRecorder) ListProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockInventoryRepository)(nil).ListProducts), arg0, arg1)
}

// ShowIndividualProducts mocks base method.
func (m *MockInventoryRepository) ShowIndividualProducts(id string) (models.ProductsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowIndividualProducts", id)
	ret0, _ := ret[0].(models.ProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShowIndividualProducts indicates an expected call of ShowIndividualProducts.
func (mr *MockInventoryRepositoryMockRecorder) ShowIndividualProducts(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowIndividualProducts", reflect.TypeOf((*MockInventoryRepository)(nil).ShowIndividualProducts), id)
}

// UpdateInventory mocks base method.
func (m *MockInventoryRepository) UpdateInventory(pid, stock int) (models.ProductsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInventory", pid, stock)
	ret0, _ := ret[0].(models.ProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInventory indicates an expected call of UpdateInventory.
func (mr *MockInventoryRepositoryMockRecorder) UpdateInventory(pid, stock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInventory", reflect.TypeOf((*MockInventoryRepository)(nil).UpdateInventory), pid, stock)
}
