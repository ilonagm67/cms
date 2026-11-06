package service

import (
	"cms/internal/entity"
	"cms/internal/usecase"
	"io"
	"net/http"
	"os"
	"path"
)

type ProductService struct {
	ProductRepo usecase.ProductRepository
}

func NewProductService(ProductRepo usecase.ProductRepository) *ProductService {
	return &ProductService{ProductRepo: ProductRepo}
}

func (Service *ProductService) Delete(name string, weight int) error {
	err := Service.ProductRepo.Delete(name, weight)
	if err != nil {
		return err
	}
	return nil
}

func (Service *ProductService) Get(name string, weight int) (*entity.Product, error) {
	product, err := Service.ProductRepo.Get(name, weight)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (Service *ProductService) List() (map[string]map[int]*entity.Product, error) {
	list, err := Service.ProductRepo.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (Service *ProductService) DownloadImage(URL, Folder, Unique string) (string, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return "", err
	}

	_, err = os.ReadDir(Folder)
	if err != nil {
		err = os.Mkdir(Folder, 0755)
		if err != nil {
			return "", err
		}
	}

	filepath := path.Join(Folder, Unique+".jpg")
	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	return filepath, nil
}
