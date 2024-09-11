package main

import (
    "database/sql"
	"os"
	"log"
	"bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "io/ioutil"

    _ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

// Data structure
type Product struct {
    ID           int64     `json:"id"`
    Title        string    `json:"title"`
    BodyHTML     string    `json:"body_html"`
    Vendor       string    `json:"vendor"`
    ProductType  string    `json:"product_type"`
    Variants     []Variant `json:"variants"`
}

type Variant struct {
    ID                int64  `json:"id"`
    Price             string `json:"price"`
    InventoryQuantity int    `json:"inventory_quantity"`
}

func connectDB() (*sql.DB, error) {
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), 
		os.Getenv("POSTGRES_PORT"), 
		os.Getenv("POSTGRES_USER"), 
		os.Getenv("POSTGRES_PASSWORD"), 
		os.Getenv("POSTGRES_DBNAME"))
    
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("Successfully connected to PostgreSQL!")
    return db, nil
}

func getProducts(db *sql.DB) ([]Product, error) {
    rows, err := db.Query(`
        SELECT p.id, p.title, p.body_html, p.vendor, p.product_type, 
               v.id, v.price, v.inventory_quantity
        FROM products p
        LEFT JOIN variants v ON p.id = v.product_id
    `)
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    products := make(map[int64]*Product)
    for rows.Next() {
        var p Product
        var v Variant

		err := rows.Scan(&p.ID, &p.Title, &p.BodyHTML, &p.Vendor, &p.ProductType, &v.ID, &v.Price, &v.InventoryQuantity)        
		if err!= nil {
            return nil, err
        }

        if _, exists := products[p.ID];!exists {
            products[p.ID] = &p
        }

        products[p.ID].Variants = append(products[p.ID].Variants, v)
    }

    productList := make([]Product, 0, len(products))
    for _, product := range products {
        productList = append(productList, *product)
    }

    return productList, nil
}

func updateShopifyProduct(product Product) error {
	shopName := os.Getenv("SHOP_NAME")
    accessToken := os.Getenv("SHOPIFY_ACCESS_TOKEN")
    shopifyURL := fmt.Sprintf("https://%s.myshopify.com/admin/api/2024-07/products/%d.json", shopName, product.ID)

    productData := map[string]Product{"product": product}
    jsonData, err := json.Marshal(productData)
    if err != nil {
        return err
    }

    req, err := http.NewRequest("PUT", shopifyURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return err
    }

    req.Header.Add("X-Shopify-Access-Token", accessToken)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode!= 200 {
        body, err := ioutil.ReadAll(resp.Body)
        if err!= nil {
            return err
        }
        log.Printf("Error updating product %d: %s", product.ID, string(body))
        return err
    }

    return nil
}

func main() {
	err := godotenv.Load()
    if err!= nil {
        log.Fatal(err)
    }

    db, err := connectDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    products, err := getProducts(db)
    if err != nil {
        log.Fatal(err)
    }

    for _, product := range products {
        err := updateShopifyProduct(product)
        if err != nil {
            log.Printf("Error updating product %d: %v", product.ID, err)
        }
    }

    log.Println("All products updated successfully!")
}
