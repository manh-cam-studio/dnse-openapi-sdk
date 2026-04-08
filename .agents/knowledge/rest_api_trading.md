# DNSE API Knowledge Base

This document contains important details on extracting market data correctly from DNSE endpoints, discovered and tested during our quantitative backtesting/replay implementation.

## 1. REST API: Historical OHLC (Nến)

*   **Endpoint:** `GET https://api.dnse.com.vn/chart-api/v2/ohlcs/derivative`
*   **Params:**
    *   `symbol`: Symbol name, e.g., `VN30F1M`
    *   `resolution`: Timeframe e.g., `1` (1m), `5` (5m), `15`, `30`, `60`, `D`, `W`, `M`
    *   `from`: Start Unix timestamp in seconds
    *   `to`: End Unix timestamp in seconds
*   **Limits:** Max query window is 24 hours (86400 seconds) per request. To pull deep historical data, you must paginate by chunking into multiple 24h spans.
*   **Authentication:** Not required.

## 2. GraphQL API: Historical Ticks / Order Match (Khớp Lệnh)

*   **Endpoint:** `POST https://api.dnse.com.vn/price-api/query`
*   **Headers Required:**
    *   `Content-Type: application/json`
    *   `Origin: https://entradex.dnse.com.vn`
*   **Payload Example:**
    ```json
    {
      "query": "query GetKrxTicksBySymbols { GetKrxTicksBySymbols(symbols: \"41I1G4000\", date: \"2026-04-03\", limit: 500000, board: 2) { ticks { symbol matchPrice matchQtty sendingTime side } } }"
    }
    ```
*   **Key Nuances:**
    *   **Internal Symbol IDs:** The GraphQL query *does not* accept literal string tickers like `VN30F1M`. You must use their internal KRX object identification string.
        *   These strings correspond to the futures expiration (đáo hạn).
        *   Examples: `41I1G4000`, `41I1G3000`, `41I1G2000`.
    *   **Limit:** There are roughly 15,000 to 50,000 matches per session. Setting `limit: 500000` is plenty to retrieve an entire day’s tick execution log in one single query.
    *   **Paginating History:** To collect long-term historical ticks, loop through days backwards. You can just decrement the `date` parameter (`YYYY-MM-DD`) and swap the internal `symbols` ID whenever rolling back over contract expiration boundaries.
*   **Side Semantics:** `side: 1` often represents BUYS into the Ask or specific aggressive sides. Need correlating logic or testing to be perfectly mapped based on the delta.
*   **Sending Time:** ISO 8601 string containing microsecond/millisecond precision (e.g. `"2026-04-03T07:45:00.706Z"`).

## Usage in System
*   **CLI Data Exporter:** We use the `vnfx-cli fetch-ticks <symbol_id> <date_str> <out.parquet>` command to parse the JSON and map it to a flat `TickData` Polars structure compatible with Rust `tick_replay_core` zero-copy buffers.


---
sidebar_position: 2
---


# Rate Limits

DNSE OpenAPI áp dụng rate limit theo từng APIKey và từng Endpoint.

Rate limit được định nghĩa bởi:

- Rate: tổng số request trong 1 giờ
- Quota: tổng số request trong 24 giờ (1 ngày)

### Normal Rate Limits

## Rate Limits

| Tên API                         | Endpoint                                                                                                                                                 | Rate / giờ | Quota / ngày |
|---------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------|------------|--------------|
| Thông tin tiền                  | <a href="https://developers.dnse.com.vn/docs/dnse/get-account-balances">/Get Account Balance</a>                              | 10,000     | 100,000      |
| Tài khoản giao dịch             | <a href="https://developers.dnse.com.vn/docs/dnse/get-accounts">/Get Accounts</a>                                             | 100        | 1,000        |
| Danh sách gói vay               | <a href="https://developers.dnse.com.vn/docs/dnse/get-loan-packages">/Get Loan Packages</a>                                   | 100        | 1,000        |
| Sức mua, sức bán                | <a href="https://developers.dnse.com.vn/docs/dnse/get-ppse">/Get PPSE</a>                                                     | 10,000     | 100,000      |
| Sổ lệnh                         | <a href="https://developers.dnse.com.vn/docs/dnse/get-orders">/Get Orders</a>                                                 | 10,000     | 100,000      |
| Lịch sử lệnh                    | <a href="https://developers.dnse.com.vn/docs/dnse/get-orders-history">/Get Order History</a>                                  | 10,000     | 100,000      |
| Chi tiết lệnh theo ID           | <a href="https://developers.dnse.com.vn/docs/dnse/get-order-detail">/Get Order Detail</a>                                     | 10,000     | 100,000      |
| Vị thế nắm giữ                  | <a href="https://developers.dnse.com.vn/docs/dnse/get-accounts-account-no-positions">/Get Positions</a>                       | 10,000     | 100,000      |
| Chi tiết vị thế theo ID         | <a href="https://developers.dnse.com.vn/docs/dnse/get-accounts-positions-position-id">/Get Position Detail by ID</a>          | 10,000     | 100,000      |
| Gửi Email OTP                   | <a href="https://developers.dnse.com.vn/docs/dnse/send-email-otp">/Send Email OTP</a>                                         | 100        | 1,000        |
| Xác thực OTP                    | <a href="https://developers.dnse.com.vn/docs/dnse/2-fa-verification">/Create Trading Token</a>                                | 100        | 1,000        |
| Đặt lệnh                        | <a href="https://developers.dnse.com.vn/docs/dnse/place-order">/Place Order</a>                                               | 50,000     | 100,000      |
| Sửa lệnh                        | <a href="https://developers.dnse.com.vn/docs/dnse/replace-order">/Replace Order</a>                                           | 50,000     | 100,000      |
| Hủy lệnh                        | <a href="https://developers.dnse.com.vn/docs/dnse/cancel-order">/Cancel Order</a>                                             | 50,000     | 100,000      |
| Đóng vị thế                     | <a href="https://developers.dnse.com.vn/docs/dnse/post-accounts-positions-position-id-close">/Close Position</a>              | 50,000     | 100,000      |
| Thông tin giao dịch chứng khoán | <a href="https://developers.dnse.com.vn/docs/dnse/get-secdef">/Get Security Definition</a>                                    | 1,000      | 10,000       |
| Chi tiết mã chứng khoán         | <a href="https://developers.dnse.com.vn/docs/dnse/get-instruments">/Get Instruments</a>                                       | 1,000      | 10,000       |
| Lịch sử OHLC                    | <a href="https://developers.dnse.com.vn/docs/dnse/get-ohlc-history">/Get OHLC</a>                                             | 1,000      | 10,000       |
| Lịch sử khớp lệnh               | <a href="https://developers.dnse.com.vn/docs/dnse/get-price-symbol-trades">/Get Trades</a>                                    | 1,000      | 10,000       |
| Dữ liệu khớp gần nhất           | <a href="https://developers.dnse.com.vn/docs/dnse/get-price-symbol-trades-latest">/Get Latest Trades</a>                      | 1,000      | 10,000       |

### Notes

DNSE có thể cung cấp thông tin về giới hạn sử dụng API thông qua các header trong response:

| Header                | Ý nghĩa                                  |
| --------------------- | ---------------------------------------- |
| X-RateLimit-Limit     | Số lượng request tối đa được phép        |
| X-RateLimit-Remaining | Số request còn lại trong chu kỳ hiện tại |
| X-RateLimit-Reset     | Thời điểm giới hạn được làm mới          |


Khi vượt quá giới hạn, hệ thống sẽ trả về **HTTP 429 (Too Many Requests)**

```json lines
429
Too Many Requests
{
    "error": "Rate Limit Exceeded"
}
```

### Khuyến nghị

- Phân bổ tần suất gọi API hợp lý trong từng khoảng thời gian
- Hạn chế gọi lặp lại các dữ liệu ít thay đổi bằng cách cache
- Ưu tiên sử dụng các API hỗ trợ xử lý nhiều dữ liệu trong một lần gọi (nếu có)
- Theo dõi số lượng request còn lại để chủ động điều chỉnh hành vi gọi API


***

---
sidebar_position: 2
---

#  Tài khoản giao dịch
---

Mỗi nhà đầu tư khi mở tài khoản tại DNSE sẽ có các thông tin định danh duy nhất trên hệ thống. Một tài khoản có thể sở hữu một hoặc nhiều tiểu khoản giao dịch. Việc nắm rõ cấu trúc này rất quan trọng để người dùng tích hợp và sử dụng API đúng cách.

Các thông tin này được trả ra trong response của <a
href="https://developers.dnse.com.vn/docs/dnse/get-accounts">/get-accounts</a>


```json lines
{
  "name": "Nguyen Hoang A",         // Họ tên khách hàng
  "custodyCode": "064CAA8386",      // Số tài khoản lưu ký tại VSD
  "investorId": "1002003456",       // Mã định danh khách hàng tại DNSE
  "accounts": [                     // Danh sách tiểu khoản thuộc tài khoản
    {
      "id": "0001009212",           // Số tiểu khoản giao dịch
      "dealAccount": true,          // Tiểu khoản theo Deal hoặc không (true/ false)
      "derivativeAccount": true,    // Tiểu khoản được phép giao dịch phái sinh hoặc không (true/ false)
      "derivative": {
        "status": "ACTIVE"          // Trạng thái tiểu khoản phái sinh (ACTIVE: Hoạt động/ INACTIVE: Ngưng hoạt động)
      }
    },
    {
      "id": "0001177757",           // Số tiểu khoản giao dịch            
      "dealAccount": true,          // Tiểu khoản theo Deal hoặc không (true/ false)
      "derivativeAccount": false,    // Tiểu khoản được phép giao dịch phái sinh hoặc không (true/ false)
      "derivative": {}
    }
  ]
}
```


---
sidebar_position: 4
---


# Lệnh giao dịch
---

### Vòng đời lệnh

Vòng đời lệnh mô tả các trạng thái mà một lệnh có thể đi qua kể từ lúc bạn gửi yêu cầu đến khi kết thúc.

<div className="guideImg">

[![Locale Dropdown](https://cdn.dnse.com.vn/dnse-openapi/doc/img/sd3.png)](https://cdn.dnse.com.vn/dnse-openapi/doc/img/sd3.png)
</div>

### Trạng thái lệnh

| Trạng thái          | Giải nghĩa            | Chú thích                                                                                 |
|---------------------|-----------------------|-------------------------------------------------------------------------------------------|
| **Pending**         | Lệnh mới được tạo     | Lệnh vừa gửi lên hệ thống, đang được kiểm tra và xử lý nội bộ                             |
| **PendingNew**      | Lệnh chờ gửi lên Sở   | Lệnh hợp lệ và đang chờ gửi lên hệ thống Sở giao dịch                                     |
| **New**             | Lệnh chờ khớp         | Lệnh được Sở ghi nhận và đang chờ khớp theo điều kiện thị trường                          |
| **PartiallyFilled** | Lệnh đã khớp một phần | Một phần khối lượng đã khớp, phần còn lại tiếp tục chờ khớp                               |
| **Filled**          | Lệnh khớp toàn bộ     | Toàn bộ khối lượng lệnh đã được khớp thành công                                           |
| **PendingReplace**  | Lệnh chờ sửa          | Yêu cầu sửa lệnh được ghi nhận, đang chờ hệ thống/Sở xử lý thay đổi                       |
| **PendingCancel**   | Lệnh chờ hủy          | Yêu cầu hủy lệnh đang chờ hệ thống/Sở xử lý                                               |
| **Canceled**        | Lệnh hủy thành công   | Lệnh đã được hủy thành công và không còn hiệu lực giao dịch                               |
| **Rejected**        | Lệnh bị từ chối       | Lệnh không được chấp nhận do không đáp ứng điều kiện (gói vay, sức mua, hạn mức cho vay…) |
| **Expired**         | Lệnh hết hạn          | Lệnh hết hiệu lực do kết thúc phiên hoặc quá thời gian hiệu lực mà chưa được khớp         |
| **DoneForDay**      | Lệnh đã được giải tỏa | Lệnh kết thúc vòng đời trong ngày giao dịch                                               |

### Đặt lệnh

Dưới đây là các thông tin bắt buộc cần gửi đối với một yêu cầu (Request) đặt lệnh.

- **`marketType`**: Phân loại giao dịch
    - `STOCK`: giao dịch cơ sở
    - `DERIVATIVE`: giao dịch phái sinh
- **`orderCategory`**: Loại lệnh thường trong ngày (NORMAL)
- **`accountNo`:** Tiểu khoản giao dịch, được trả trong response Endpoint <a
  href="https://developers.dnse.com.vn/docs/dnse/get-accounts">Tài khoản giao dịch.</a>
- **`symbol`**: Mã chứng khoán giao dịch
- **`loanPackageId`**: Gói vay giao dịch, xem thêm thông tin về gói vay <a
  href="https://developers.dnse.com.vn/docs/guide/trading-api/dnse_margin#gói-vay-loan-packages">
  tại đây.</a>
- **`side`**: Chiều mua (NB) hoặc bán (NS)
- **`orderType`**: Loại lệnh tương ứng với sàn giao dịch
    - Sàn HOSE: ATO, ATC, LO, MTL
    - Sàn HNX: LO, MTL, MOK, MAK, ATC, PLO
    - Sàn Upcom: LO
- **`quantity`**: Khối lượng đặt

    - Khối lượng đặt không vượt quá khối lượng tối đa có thể mua (`qmaxBuy`)hoặc có thể bán (`qmaxSell`) trên tiểu khoản
      giao dịch, người dùng truy vấn thông tin đối với từng mã chứng khoán qua Endpoint <a
      href="https://developers.dnse.com.vn/docs/dnse/get-ppse">/Sức mua, sức bán.</a>
    - Với giao dịch cơ sở, khối lượng đặt là lô chẵn (100,200,...) hoặc lô lẻ (1,2,..99). Khối lượng lẻ lô (101,102,...)
      là không hợp lệ.
- **`price`**: Giá đặt
    - Nếu loại lệnh là LO, giá đặt phải > 0 và phải nằm trong khoảng giá trần sàn của mã chứng khoán tại phiên giao dịch
      đó.
    - Nếu loại lệnh khác LO, giá đặt truyền lên luôn = 0.

<details>
  <summary>VD Yêu cầu đặt lệnh</summary>

```json lines
{
  "method": "POST",
  "path": "/accounts/orders",
  "query": {
    "marketType": "STOCK",
    "orderCategory": "NORMAL"
  },
  "headers": {
    "x-api-key": "lB58g6iWzyrNx2EhwwQXeYeoAnkzlaXkJWi",
    // APIkey được cấp khi đăng ký dịch vụ
    "x-signature": "fjsdhfryt6aaa6c91a8f88b472c9721fde161e0d89df8c",
    // Chữ ký số theo thuật toán HMAC SHA256
    "trading-token": "7ceef658-9f01-414e-8b3e-faa77bb9061e",
    // Token đặt lệnh         
    "date": "Fri, 16 Jan 2026 07:11:30 +0000"
    // Thời gian tạo yêu cầu (UTC)
  },
  "body": {
    "accountNo": "0003979888",
    // Số tiểu khoản giao dịch
    "symbol": "HPG",
    // Mã chứng khoán đặt lệnh
    "side": "NB",
    // Chiều lệnh giao dịch 
    "orderType": "LO",
    // Loại lệnh giao dịch
    "price": 25950,
    // Giá đặt
    "quantity": 100,
    // Khối lượng đặt
    "loanPackageId": 5757
    // Mã gói vay 
  }
}
```

</details>

Khi lệnh khớp mua, hệ thống hình thành các Deals (hay còn gọi là danh mục tài sản) theo cặp `symbol` - `loanPackage`.
Nếu mua cùng mã nhưng khác gói vay → tạo Deal tách biệt (rủi ro được quản trị riêng).

### Sửa lệnh

**Điều kiện chung:**

- Chỉ được sửa lệnh LO trong phiên giao dịch liên tục và áp dụng cho lệnh ở trạng thái Chờ khớp (New) hoặc Đã khớp một
  phần (PartiallyFilled)
- Giá sửa phải nằm trong biên độ trần sàn của mã chứng khoán vào phiên giao dịch đó.
- Nếu giá hoặc khối lượng sau khi sửa vượt quá sức mua /sức bán cho phép, yêu cầu sửa lệnh sẽ bị từ chối.

**Sửa lệnh cơ sở:**

- Khi sửa lệnh thành công, hệ thống hủy lệnh hiện tại và đặt lại một lệnh mới với thông tin đã chỉnh sửa.
- Cho phép sửa đồng thời giá và khối lượng.
- Thứ tự ưu tiên của lệnh sau khi sửa sẽ được xác định lại theo thời điểm ghi nhận sửa lệnh thành công.

**Sửa lệnh phái sinh:**

- Người dùng chỉ được phép sửa hoặc giá hoặc khối lượng trong một yêu cầu.
- Khối lượng sửa phải lớn hơn khối lượng đã khớp (nếu lệnh đã khớp một phần).
- Nếu khối lượng sửa lớn hơn khối lượng ban đầu, thứ tự ưu tiên của lệnh sẽ được thay đổi.

<details>
  <summary>VD Yêu cầu hủy lệnh</summary>

```json lines
{
  "method": "PUT",
  "path": "/accounts/{account_no}/orders/{order_id}",
  "query": {
    "marketType": "STOCK",
    "orderCategory": "NORMAL"
  },
  "headers": {
    "x-api-key": "lB58g6iWzyrNx2EhwwQXeYeoAnkzlaXkJWi",
    // APIkey được cấp khi đăng ký dịch vụ
    "x-signature": "fjsdhfryt6aaa6c91a8f88b472c9721fde161e0d89df8c",
    // Chữ ký số theo thuật toán HMAC SHA256
    "trading-token": "7ceef658-9f01-414e-8b3e-faa77bb9061e",
    // Token đặt lệnh         
    "date": "Fri, 16 Jan 2026 07:11:30 +0000"
    // Thời gian tạo yêu cầu (UTC)
  },
  "body": {
    "price": 25950,
    // Giá sửa
    "quantity": 100
    // Khối lượng sửa
  }
}
```

</details>


---
sidebar_position: 3
---

# Margin tại DNSE
---

Khác với cách quản trị rủi ro trên tài khoản tổng, với Margin tại DNSE - mỗi một Deal (bao gồm một mã chứng khoán và một gói vay ký quỹ) khác nhau của khách hàng sẽ được quản trị tách bạch:

- Các Deals khác nhau về tỷ lệ vay được quản lý tách biệt, danh mục của khách hàng có thể gồm nhiều Deals vay khác nhau.
- Cách tính giá trung bình, giá hòa vốn của mỗi Deal (đã bao gồm lãi vay và các chi phí khác) rõ ràng, chính xác hơn so với giá vốn truyền thống.
- Tỷ lệ ký quỹ của mỗi Deal được kiểm soát độc lập. DNSE sẽ chỉ yêu cầu ký quỹ bổ sung hoặc bán giải chấp Deal có tỷ lệ xuống dưới mức cảnh báo mà không ảnh hưởng tới các Deal an toàn khác.

Đây là sự khác biệt mà DNSE xây dựng để khách hàng của mình quản lý tài sản minh bạch hơn (mô hình Isolated Margin)

### Gói vay (Loan packages)

Gói vay là khái niệm đại diện cho chính sách sản phẩm tại DNSE. Mỗi gói vay quy định các điều kiện áp dụng cho giao dịch, bao gồm: phí giao dịch, tỷ lệ vay, tỷ lệ ký quỹ, lãi suất vay, kỳ hạn và một số thông tin khác.

Các gói vay được áp dụng cho từng mã chứng khoán được trả về trong response Endpoint <a href="https://developers.dnse.com.vn/docs/dnse/get-loan-packages">/Danh sách gói vay</a>

**Danh sách gói vay cơ sở**

VD: Danh sách gói vay cho mã HPG

```json lines
{
  "symbol": "HPG",         // Mã chứng khoán
  "marketType": "STOCK",   // Loại thị trường (STOCK: cơ sở / DERIVATIVE: phái sinh)
  "loanPackages": [       // Danh sách gói vay 
    {
      "id": 1775,         // Mã gói vay
      "name": "Mana RocketX LS 5.99% HPG - KQ 100%",    // Tên gói vay 
      "initialRate": 1,             // Tỷ lệ ký quỹ ban đầu 
      "interestRate": 0.0599,       // Tỷ lệ lãi vay (nếu phát sinh ứng sức mua, nợ margin)
      "liquidRate": 0.3,            // Tỷ lệ xử lý (force sell)
      "maintenanceRate": 0.4,       // Tỷ lệ duy trì (call margin)
      "type": "M",              // Loại gói vay (M: gói vay margin/ N: gói tiền mặt)
      "brokerFirmBuyingFeeRate": 0,     //  Phí mua chứng khoán cơ sở DNSE thu
      "brokerFirmSellingFeeRate": 0     //  Phí bán chứng khoán cơ sở DNSE thu
    },
    {
      "id": 1769,       // Mã gói vay
      "name": "Rocket X LS 5.99% HPG - KQ 50%",    // Tên gói vay 
      "initialRate": 0.5,           // Tỷ lệ ký quỹ ban đầu 
      "interestRate": 0.0599,       // Tỷ lệ lãi vay (nếu phát sinh ứng sức mua, nợ margin)
      "liquidRate": 0.3,            // Tỷ lệ xử lý (force sell)
      "maintenanceRate": 0.4,       // Tỷ lệ duy trì (call margin)
      "type": "M",                  // Loại gói vay (M: gói vay margin/ N: gói tiền mặt)
      "brokerFirmBuyingFeeRate": 0.00045,     //  Phí mua chứng khoán cơ sở DNSE thu
      "brokerFirmSellingFeeRate": 0.00045     //  Phí bán chứng khoán cơ sở DNSE thu
    }  
  ]
}
```

VD: Danh sách gói vay cho mã VGI

```json lines
{
  "symbol": "VGI",      // Mã chứng khoán
  "marketType": "STOCK",      // Loại thị trường (STOCK: cơ sở / DERIVATIVE: phái sinh)
  "loanPackages": [     // Danh sách gói vay 
    {
      "id": 1036,   // Mã gói vay
      "name": "GD tiền mặt",    // Tên gói vay 
      "type": "N",      // Loại gói vay (M: gói vay margin/ N: gói tiền mặt)
      "brokerFirmBuyingFeeRate": 0,     //  Phí mua chứng khoán cơ sở DNSE thu
      "brokerFirmSellingFeeRate": 0     //  Phí bán chứng khoán cơ sở DNSE thu
    }
  ]
}
```

Đối với giao dịch chứng khoán cơ sở, hệ thống sẽ trả về tối đa 2 gói vay mà người dùng có thể sử dụng để đặt lệnh cho mã chứng khoán truy vấn, bao gồm:

- Gói vay giao dịch tiền mặt:

    + Tỷ lệ ký quỹ tiền mặt 100% (`initialRate` = 1)
    + Dành cho giao dịch không sử dụng đòn bẩy tiền vay
- Gói vay ký quỹ (margin) cơ bản:

    + Dành cho giao dịch có sử dụng đòn bẩy tiền vay (`initialRate` ≠ 1)

**Danh sách gói vay phái sinh**

VD: Danh sách gói vay cho mã 41I1G1000

```json lines
{
  "symbolType": "VN30F1M",           // Mã giao dịch Hợp đồng tương lai
  "marketType": "DERIVATIVE",       // Loại thị trường 
  "loanPackages": [                 // Danh sách gói vay 
    {
      "id": 1306,                   // Mã gói vay
      "name": "Gói giao dịch 01",   // Tên gói vay 
      "initialRate": 0.1848,        // Tỷ lệ ký quỹ ban đầu
      "maintenanceRate": 0.1771,    // Tỷ lệ duy trì (call margin)
      "liquidRate": 0.1735,         // Tỷ lệ xử lý (force sell)
      "tradingFee": {               // Chính sách phí giao dịch (dành cho phái sinh)
        "id": 1304,                 // ID chính sách phí    
        "name": "Miễn phí",         // Tên phí
        "scope": "PRODUCT",         // Phạm vi áp dụng chính sách
        "channel": "ALL",           // Kênh giao dịch áp dụng
        "schemaType": "FIXED",      // Loại phí cố định
        "createdDate": "2023-02-02T04:22:56.199278Z",       // Thời điểm tạo chính sách
        "modifiedDate": "2023-02-02T04:22:56.199278Z",      // Thời điểm cập nhật chính sách
        "fixedTradingFee": 2000,        // Phí giao dịch 1 hợp đồng
        "fixedDailyCloseTradingFee": 2000      // Phí giao dịch 1 hợp đồng đóng luôn trong ngày 
      }
    }
  ]
}
```

Với sản phẩm phái sinh, thông thường tài khoản giao dịch của khách hàng chỉ được gắn một gói vay với một bộ tỷ lệ ký quỹ, duy trì và xử lý duy nhất áp dụng cho tất cả các mã phái sinh.

### Deal

Deals hay còn có thể hiểu là danh mục tài sản của khách hàng. Một Deal được hình thành bởi 1 mã chứng khoán và 1 gói vay:

- Với cùng một mã có thể có nhiều Deals độc lập nếu bạn mua cùng mã nhưng chọn gói vay khác nhau
- Việc cho vay, thu nợ, quản trị rủi ro được thực hiện trên từng Deal

Ví dụ:

- Lần 1: Mua 100cp HPG với gói vay “GD Tiền mặt", hệ thống sẽ tạo 1 Deal HPG Tiền mặt, tỷ lệ Ký quỹ 100%
- Lần 2: Mua 500cp HPG với gói vay margin “Tiền mặt 50%”, hệ thống sẽ tạo Deal mới HPG Tiền mặt 50% (khác với Deal bên trên), được hiểu là khách hàng ký quỹ 50% và sử dụng tiền vay 50% tính trên tổng số tiền khớp lệnh mua.
- Lần 3: Mua 200cp HPG với gói vay “Tiền mặt 100%”, do cùng gói vay với lần thứ 1, nên hệ thống gộp khối lượng mua thêm vào Deal HPG Tiền mặt 100%; tổng khối lượng sau mua là 300cp.

<div className="guideImg">

[![Locale Dropdown](https://cdn.dnse.com.vn/dnse-openapi/doc/img/deal.png)](https://cdn.dnse.com.vn/dnse-openapi/doc/img/deal.png)
</div>

Khách hàng có thể tìm hiểu thêm về sản phẩm giao dịch ký quỹ theo DEAL [tại đây.](https://hdsd.dnse.com.vn/san-pham-dich-vu/sp-giao-dich-ky-quy-theo-deal/thong-tin-chung)



