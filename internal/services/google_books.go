package services

import (
	"elibrary/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type GoogleBooksService struct {
	apiKey string
}

func NewGoogleBooksService(apiKey string) *GoogleBooksService {
	return &GoogleBooksService{apiKey: apiKey}
}

func (s *GoogleBooksService) GetBookByISBN(isbn string) (*models.Book, error) {
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=isbn:%s&key=%s", isbn, s.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Items []struct {
			VolumeInfo struct {
				Title               string   `json:"title"`
				Authors             []string `json:"authors"`
				IndustryIdentifiers []struct {
					Type string `json:"type"`
					ID   string `json:"identifier"`
				} `json:"industryIdentifiers"`
			} `json:"volumeInfo"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Items) == 0 {
		return nil, fmt.Errorf("book not found")
	}

	book := &models.Book{
		Title:  result.Items[0].VolumeInfo.Title,
		Author: result.Items[0].VolumeInfo.Authors[0],
		ISBN:   isbn,
	}

	return book, nil
}
