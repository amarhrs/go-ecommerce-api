package helpers

import (
	"amarhrs/ecommerce/models"
	"fmt"
	"regexp"
	"strings"
)

func ValidateProductName(name string) (bool, string) {
	name = strings.TrimSpace(name)
	if name == "" {
		return false, "Field name required"
	}
	if len(name) < 3 {
		return false, "Product name must be at least 3 characters long"
	}
	if len(name) > 50 {
		return false, "Product name cannot be more than 50 characters long"
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9\s\-_]+$`).MatchString(name) {
		return false, "Product name can only contain letters, numbers, spaces, hyphens, and underscores"
	}
	return true, ""
}

func ValidatePrice(price float64) (bool, string) {
	if price <= 0 {
		return false, "Price must be greater than 0"
	}
	if price > 100000000 {
		return false, "Price is too high"
	}

	// Cek maksimal 2 digit desimal
	priceStr := fmt.Sprintf("%.2f", price)
	if matched, _ := regexp.MatchString(`^\d+(\.\d{1,2})?$`, priceStr); !matched {
		return false, "Price must have at most 2 decimal places"
	}
	return true, ""
}

func ValidateStock(stock int) (bool, string) {
	if stock < 0 {
		return false, "Stock cannot be negative"
	}
	return true, ""
}

func ValidateDescription(desc string) (bool, string) {
	desc = strings.TrimSpace(desc)
	if desc == "" {
		return false, "Field description required"
	}
	if len(desc) > 500 {
		return false, "Description cannot exceed 500 characters"
	}
	return true, ""
}

// Validasi format image URL (jpg/jpeg/png)
func ValidateImageFormat(url string) (bool, string) {
	url = strings.ToLower(url)
	if !(strings.HasSuffix(url, ".jpg") || strings.HasSuffix(url, ".jpeg") || strings.HasSuffix(url, ".png")) {
		return false, "Invalid image format, only jpg/jpeg/png allowed"
	}
	return true, ""
}

func ValidateProductImages(images []models.ProductImage) (bool, string) {
	if len(images) == 0 {
		return false, "At least one product image is required"
	}

	for _, img := range images {
		if ok, msg := ValidateImageFormat(img.URL); !ok {
			return false, msg
		}
	}

	return true, ""
}

func ValidateProductCategoryName(name string) (bool, string) {
	name = strings.TrimSpace(name)
	if name == "" {
		return false, "Field name required"
	}
	if len(name) < 3 {
		return false, "Category name must be at least 3 characters long"
	}
	if len(name) > 50 {
		return false, "Category name cannot be more than 50 characters long"
	}
	if !regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(name) {
		return false, "Category name can only contain letters and spaces"
	}
	return true, ""
}

func ValidateProductInput(product models.Product) (bool, []string) {
	var errs []string

	if ok, msg := ValidateProductName(product.Name); !ok {
		errs = append(errs, msg)
	}
	if ok, msg := ValidatePrice(product.Price); !ok {
		errs = append(errs, msg)
	}
	if ok, msg := ValidateStock(product.Stock); !ok {
		errs = append(errs, msg)
	}
	if ok, msg := ValidateDescription(product.Description); !ok {
		errs = append(errs, msg)
	}
	if ok, msg := ValidateProductImages(product.Images); !ok {
		errs = append(errs, msg)
	}

	return len(errs) == 0, errs
}

func ValidateProductCategoryInput(category models.ProductCategory) (bool, []string) {
	var errs []string

	if ok, msg := ValidateProductCategoryName(category.Name); !ok {
		errs = append(errs, msg)
	}

	return len(errs) == 0, errs
}
