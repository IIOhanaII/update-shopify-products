# Update Shopify Products

This Go project is designed to update product data in a Shopify store after it has been modified in bulk within a database. It works in tandem with the companion project [Store Shopify Products](https://github.com/IIOhanaII/store-shopify-products), which must be used first to fetch and store the Shopify product data into a database.

## Important Note

### **Before using this project, you must first run the companion project [Store Shopify Products](https://github.com/IIOhanaII/store-shopify-products).**

The companion project is responsible for fetching product data from Shopify and storing it in a database. Once the data has been modified in bulk in the database, this project can be used to push the changes back to Shopify. This ensures the integrity and consistency of product updates across both the database and the Shopify store.

## Features

- **Update Shopify product data**: Push updated product data from a database back to your Shopify store.
- **Database-driven modifications**: Allows you to handle bulk modifications via direct database updates.
- **API-driven synchronization**: Integrates with Shopify’s API to ensure all modifications are reflected in the store.

## Prerequisites

1. **Go**: Ensure that you have Go installed. You can download it [here](https://golang.org/dl/).
2. **PostgreSQL Database**: The database used in the companion project should remain connected.
3. **Shopify Store**: Ensure that your Shopify store is set up with the appropriate API credentials.
4. **Companion Project**: You must have already used the [Store Shopify Products](https://github.com/IIOhanaII/store-shopify-products) project to fetch and modify your Shopify product data in the database.
5. **.env File**: Create a `.env` file in the root directory of your project to store sensitive variables like API keys, database credentials, etc.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/IIOhanaII/update-shopify-products.git
    cd update-shopify-products
    ```

2. Install dependencies:

    ```bash
    go mod download
    ```

3. Set up your `.env` file:

    Create a `.env` file in the root directory of the project and include the following sensitive variables:

    ```
    SHOP_NAME=your_shopify_store_name
    SHOPIFY_ACCESS_TOKEN=your_shopify_access_token
    POSTGRES_DBNAME=your_database_name
    POSTGRES_USER=your_database_user
    POSTGRES_PASSWORD=your_database_password
    POSTGRES_HOST=your_database_host
    POSTGRES_PORT=your_database_port
    ```

   **Explanation of variables**:
   - `SHOP_NAME`: Your Shopify store name (e.g., `yourstore` within `yourstore.myshopify.com`).
   - `SHOPIFY_ACCESS_TOKEN`: The access token for authenticating with the Shopify API.
   - `POSTGRES_DBNAME`: Name of the PostgreSQL database containing the modified product data.
   - `POSTGRES_USER`: PostgreSQL database user.
   - `POSTGRES_PASSWORD`: Password for the PostgreSQL database user.
   - `POSTGRES_HOST`: Host address of your PostgreSQL database (e.g., `localhost`).
   - `POSTGRES_PORT`: Port on which the PostgreSQL database is running (default is `5432`).

4. Run the project:

    ```bash
    go run main.go
    ```

## Usage

1. **Pre-Requisite**: Before using this project, ensure that the Shopify product data has been fetched and stored in the database using the companion project [Store Shopify Products](https://github.com/IIOhanaII/store-shopify-products). This project will read the modified product data from the same database.

2. **Pushing updated data to Shopify**:
   After you have modified the product data in the database, run this project to push the updates back to Shopify. The project will connect to the database, read the modified data, and use the Shopify API to update the products in your store.

## Project Structure

```bash
├── .env              # Contains sensitive environment variables
├── main.go           # Main entry point of the application
├── db                # Database connection and product retrieval logic
├── shopify           # Shopify API interaction logic
├── go.mod            # Go module dependencies
└── README.md         # Project documentation
```

## Environment Variables

Sensitive information such as Shopify API credentials and database connection details are stored in the `.env` file. The project uses the [godotenv](https://github.com/joho/godotenv) package to load environment variables from the `.env` file into the application.

## Database Configuration

Ensure that the PostgreSQL database containing the product data (fetched and modified using the companion project) is properly set up and accessible.

Example PostgreSQL connection string in the `.env`:

```
POSTGRES_DBNAME=shopify_db
POSTGRES_USER=postgres
POSTGRES_PASSWORD=yourpassword
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
```

## Contributing

If you'd like to contribute to this project, feel free to submit a pull request or file an issue. Contributions like bug fixes, new features, or documentation improvements are welcome.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Related Projects

- [Store Shopify Products](https://github.com/IIOhanaII/store-shopify-products): Use this project to fetch Shopify product data and store it in a database for bulk modifications before using the current project to update the store.
