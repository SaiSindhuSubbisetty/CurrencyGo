
## Currency Converter Service

This **Currency Converter Service** is built using **Go** and **gRPC**. It interacts with a **Java-based Wallet App** to provide currency conversion functionalities. The service retrieves conversion rates from a PostgreSQL database and performs conversions between various currencies. The service exposes a gRPC API to the Wallet App, enabling seamless currency conversion between different currencies in the wallet.

## Features

- **Currency Conversion**: Converts an amount from one currency to another using conversion rates stored in a PostgreSQL database.
- **gRPC Service**: Exposes a gRPC API for efficient and low-latency communication with the Java Wallet App.
- **Database Integration**: Retrieves conversion rates from a PostgreSQL database, ensuring accurate and up-to-date conversion rates.
- **Security**: Ensures secure communication and data exchange with the Java Wallet App.

## Tech Stack

- **Go**: For backend logic and service implementation.
- **gRPC**: For fast and efficient communication between the Currency Converter service and the Java Wallet App.
- **PostgreSQL**: For storing and retrieving currency conversion rates.
- **Protocol Buffers (proto)**: For defining the gRPC service and message types.
- **Java**: Wallet App interacts with the service via gRPC to fetch conversion rates and perform transactions.

## Prerequisites

- **Go** 1.16 or higher
- **PostgreSQL** database with a table for currency conversion rates
- **Protocol Buffers** for defining and compiling `.proto` files
- **gRPC** libraries for Go and Java
- **Java** (for Wallet App integration)

## Database Setup

### 1. Create PostgreSQL Database

Run the following SQL commands to create the necessary table for storing currency conversion rates:

```sql
CREATE DATABASE currencydb;

\c currencydb;

CREATE TABLE conversion_rates (
    currency VARCHAR(10) PRIMARY KEY,
    rate FLOAT NOT NULL
);

-- Example data for conversion rates (rates relative to INR)
INSERT INTO conversion_rates (currency, rate) VALUES ('USD', 75.0);
INSERT INTO conversion_rates (currency, rate) VALUES ('EUR', 85.0);
INSERT INTO conversion_rates (currency, rate) VALUES ('GBP', 95.0);
```

## Installation and Setup

### 1. Clone the Repository

```bash
git clone https://github.com/SaiSindhuSubbisetty/CurrencyConverter.git
cd CurrencyConverter
```

### 2. Install Dependencies

Make sure you have **Go** and **gRPC** installed. If you haven't installed **gRPC** for Go, use the following command:

```bash
go get google.golang.org/grpc
go get github.com/lib/pq
```

### 3. Compile the `.proto` Files

You need to define the **gRPC service** and message types using **Protocol Buffers**. Assuming you already have a `proto` directory with a `currency_converter.proto` file, run the following command to generate Go code from the `.proto` file:

```bash
protoc --go_out=. --go-grpc_out=. proto/currency_converter.proto
```

### 4. Update Database Connection

In the code (`main.go`), update the PostgreSQL connection string if needed:

```go
connStr := "user=postgres password=1234 dbname=currencydb sslmode=disable"
```

### 5. Run the Service

Start the server by running the following command:

```bash
go run main.go
```

The server will start and listen on **port 50051**.

## gRPC Service

### 1. Service Method

The service exposes the following gRPC method:

#### `Convert` (Currency Conversion)

- **RPC**: `Convert`
- **Request**: A request that includes the amount to be converted, source currency, and target currency.
- **Response**: The converted amount.

#### Request Structure

```proto
message ConvertRequest {
  float amount = 1;          // Amount to convert
  string sourceCurrency = 2; // Source currency code (e.g., "USD")
  string targetCurrency = 3; // Target currency code (e.g., "INR")
}

message ConvertResponse {
  float convertedAmount = 1; // Converted amount
}
```

### 2. Example gRPC Client (Java Integration)

The **Java Wallet App** can integrate with this service using **gRPC**. Here's an example of how you can set up a Java client to interact with this service.

```java
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import CurrencyConverterGrpc;
import CurrencyConverterProto.ConvertRequest;
import CurrencyConverterProto.ConvertResponse;

public class CurrencyConverterClient {

    public static void main(String[] args) {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 50051)
                .usePlaintext()  // Use plain-text communication (for development, not for production)
                .build();

        CurrencyConverterGrpc.CurrencyConverterBlockingStub stub = CurrencyConverterGrpc.newBlockingStub(channel);

        ConvertRequest request = ConvertRequest.newBuilder()
                .setAmount(100.0)
                .setSourceCurrency("USD")
                .setTargetCurrency("INR")
                .build();

        ConvertResponse response = stub.convert(request);

        System.out.println("Converted Amount: " + response.getConvertedAmount());

        channel.shutdown();
    }
}
```

### 3. Example gRPC Client in Other Languages

You can generate gRPC client code for other languages like Python, Node.js, etc., by using the corresponding **protoc** compiler plugin.

## Running the Service

1. **Start the server**:

   ```bash
   go run main.go
   ```

2. **Use a gRPC client** (like the Java Wallet App) to connect to the service and perform currency conversion.

### Example Workflow in Java Wallet App

- A user requests to convert 100 USD to INR.
- The Wallet App sends a gRPC request to the Currency Converter service.
- The service retrieves the conversion rate from the PostgreSQL database and returns the converted amount.
- The Wallet App processes the response and displays the result to the user.

## Testing

To test the functionality, you can create unit tests for both the Go service and the Java client. For the Go service, you can use the `testing` package, and for Java, you can use **JUnit**.

## Security

For production, secure the gRPC communication by enabling **TLS** or using other security mechanisms like **OAuth2**. It's important to ensure sensitive data like currency rates are protected during communication.

## License

This project is licensed under the MIT License.

