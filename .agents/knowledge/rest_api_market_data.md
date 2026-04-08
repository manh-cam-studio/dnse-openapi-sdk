## Thông tin giao dịch chứng khoán

### Base URLs:
- **https://openapi.dnse.com.vn**

<span id="getSymbolSecdef"></span>

### `GET /price/{symbol}/secdef`

<h3 id="getsymbolsecdef-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|boardId|query|string|false|Mã bảng giao dịch|
|X-API-Key|header|string|true|API Key được cấp khi đăng ký dịch vụ|
|X-Aux-Date|header|string|true|Thời gian thực hiện yêu cầu|
|X-Signature|header|string|true|Chữ ký xác thực yêu cầu|
|symbol|path|string|true|Mã chứng khoán|

#### Detailed descriptions

**boardId**: Mã bảng giao dịch
- G1: Lô chẵn
- G4: Lô lẻ
- T1: Thỏa thuận trong giờ (9h - 14h45)
- T3: Thỏa thuận sau giờ (14h45 - 15h)
- T4: Thỏa thuận lô lẻ trong giờ (9h - 14h45)
- T6: Thỏa thuận lô lẻ sau giờ  (14h45 - 15h)

> Code samples

```shell
# You can also use wget
curl -X GET https://openapi.dnse.com.vn/price/{symbol}/secdef \
  -H 'Accept: application/json' \
  -H 'X-API-Key: eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==' \
  -H 'X-Aux-Date: Mon, 19 Jan 2026 07:45:23 +0000' \
  -H 'X-Signature: your_signature'

```

```http
GET https://openapi.dnse.com.vn/price/{symbol}/secdef HTTP/1.1
Host: openapi.dnse.com.vn
Accept: application/json
X-API-Key: eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==
X-Aux-Date: Mon, 19 Jan 2026 07:45:23 +0000
X-Signature: your_signature

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "X-API-Key": []string{"eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ=="},
        "X-Aux-Date": []string{"Mon, 19 Jan 2026 07:45:23 +0000"},
        "X-Signature": []string{"your_signature"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "https://openapi.dnse.com.vn/price/{symbol}/secdef", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

```javascript

const headers = {
  'Accept':'application/json',
  'X-API-Key':'eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==',
  'X-Aux-Date':'Mon, 19 Jan 2026 07:45:23 +0000',
  'X-Signature':'your_signature'
};

fetch('https://openapi.dnse.com.vn/price/{symbol}/secdef',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'X-API-Key': 'eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==',
  'X-Aux-Date': 'Mon, 19 Jan 2026 07:45:23 +0000',
  'X-Signature': 'your_signature'
}

r = requests.get('https://openapi.dnse.com.vn/price/{symbol}/secdef', headers = headers)

print(r.json())

```

```java
URL obj = new URL("https://openapi.dnse.com.vn/price/{symbol}/secdef");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

> Example responses

> 200 Response

```json
[
  {
    "marketId": "STO",
    "boardId": "G1",
    "isin": "VN000000HPG4",
    "symbol": "HPG",
    "productGrpId": "STO",
    "securityGroupId": "ST",
    "basicPrice": 26.8,
    "ceilingPrice": 28.65,
    "floorPrice": 24.95,
    "securityStatus": "UNSPECIFIED",
    "symbolAdminStatusCode": "NRM",
    "symbolTradingMethodStatusCode": "NRM",
    "symbolTradingSanctionStatusCode": "NRM"
  }
]
```

<h3 id="getsymbolsecdef-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» marketId|string|false|none|Mã thị trường niêm yết mã chứng khoán<br>- DVX: Phái sinh sàn HNX<br>- HCX: Trái phiếu doanh nghiệp HNX<br>- STO: Cổ phiếu sàn HOSE<br>- STX: Cổ phiếu sàn HNX<br>- UPX: Cổ phiếu sàn Upcom|
|» boardId|string|false|none|Mã bảng giao dịch<br>- G1: Lô chẵn<br>- G4: Lô lẻ<br>- T1: Thỏa thuận trong giờ (9h - 14h45)<br>- T3: Thỏa thuận sau giờ (14h45 - 15h)<br>- T4: Thỏa thuận lô lẻ trong giờ (9h - 14h45)<br>- T6: Thỏa thuận lô lẻ sau giờ (14h45 - 15h)|
|» isin|string|false|none|Mã định danh quốc tế|
|» symbol|string|false|none|Mã chứng khoán|
|» productGrpId|string|false|none|Nhóm sản phẩm theo thị trường<br>-FBX: Hợp đồng tương lai Trái phiếu<br>-FIO: Hợp đồng tương lai Chỉ số<br>-HCX: Trái phiếu Doanh nghiệp HNX<br>-STO: Cổ phiếu sàn HOSE<br>-STX: Cổ phiếu sàn HNX<br>-UPX: Cổ phiếu sàn Upcom|
|» securityGroupId|string|false|none|Nhóm chứng khoán<br>- BS: Trái phiếu doanh nghiệp<br>- EF: Quỹ ETF<br>- EW: Chứng quyền<br>- FU: Hợp đồng tương lai<br>- ST: Cổ phiếu|
|» basicPrice|number(double)|false|none|Giá tham chiếu ngày giao dịch|
|» ceilingPrice|number(double)|false|none|Giá trần ngày giao dịch|
|» floorPrice|number(double)|false|none|Giá sàn ngày giao dịch|
|» securityStatus|string|false|none|Trạng thái giao dịch của mã chứng khoán<br>- HALT: Ngừng giao dịch<br>- NO_HALT: Không ngừng giao dịch|
|» symbolAdminStatusCode|string|false|none|Trạng thái quản lý hành chính mã chứng khoán<br>- CR: Kiểm soát và hạn chế giao dịch<br>- CTR: Kiểm soát<br>- NRM: Bình thường<br>- RES: Hạn chế giao dịch<br>- WFR: Cảnh báo vi phạm BCTC<br>- WID: Cảnh báo vi phạm CBTT<br>- WOV: Cảnh báo vi phạm khác|
|» symbolTradingMethodStatusCode|string|false|none|Trạng thái cơ chế giao dịch mã chứng khoán<br>- NRM: Bình thường<br>- NWE: Niêm yết mới (biên độ đặc biệt)<br>- NWN: Niêm yết mới (biên độ thường)<br>- SLS: Giao dịch đặc biệt sau tạm ngưng<br>- SNE: Giao dịch đặc biệt không có giao dịch dài hạn|
|» symbolTradingSanctionStatusCode|string|false|none|Tình trạng giao dịch của mã chứng khoán<br>- NRM: Bình thường<br>- SUS: Tạm ngừng giao dịch<br>- DTL: Hủy niêm yết để chuyển sàn<br>- TFR: Ngưng giao dịch do hạn chế|

Status Code **400**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» code|string|false|none|none|
|» message|string|false|none|none|
|» status|integer|false|none|none|

Status Code **500**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» code|string|false|none|none|
|» message|string|false|none|none|
|» status|integer|false|none|none|


## Chi tiết mã chứng khoán

### Base URLs:
- **https://openapi.dnse.com.vn**

<span id="getInstruments"></span>

### `GET /instruments`

<h3 id="getinstruments-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|symbol|query|string|true|Danh sách mã chứng khoán|
|marketId|query|string|false|Mã thị trường niêm yết|
|securityGroupId|query|string|false|Nhóm chứng khoán|
|indexName|query|string|false|Chỉ số thị trường |
|limit|query|integer|false|Số bản ghi trên mỗi trang|
|page|query|integer|false|Phân trang hiện tại|
|X-API-Key|header|string|true|API Key được cấp khi đăng ký dịch vụ|
|X-Aux-Date|header|string|true|Thời gian thực hiện yêu cầu|
|X-Signature|header|string|true|Chữ ký xác thực yêu cầu|

#### Detailed descriptions

**marketId**: Mã thị trường niêm yết
- STO: Cổ phiếu sàn HOSE
- STX: Cổ phiếu sàn HNX
- UPX: Cổ phiếu sàn UPCOM
- DVX: Phái sinh
- HCX: Trái phiếu doanh nghiệp

**securityGroupId**: Nhóm chứng khoán
- ST: Cổ phiếu
- EF: Quỹ ETF
- EW: Chứng quyền
- FU: Hợp đồng tương lai
- BS: Trái phiếu

**indexName**: Chỉ số thị trường 
- VN30: Top 30 cổ phiếu sàn HOSE
- VN100: Top 100 cổ phiếu sàn HOSE
- HNX30: Top 30 cổ phiếu sàn HNX

> Code samples

```shell
# You can also use wget
curl -X GET https://openapi.dnse.com.vn/instruments?symbol=SSI%2CSHS%2CACB%2CHUT%2CDSE \
  -H 'Accept: application/json' \
  -H 'X-API-Key: eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==' \
  -H 'X-Aux-Date: Mon, 19 Jan 2026 07:45:23 +0000' \
  -H 'X-Signature: your_signature'

```

```http
GET https://openapi.dnse.com.vn/instruments?symbol=SSI%2CSHS%2CACB%2CHUT%2CDSE HTTP/1.1
Host: openapi.dnse.com.vn
Accept: application/json
X-API-Key: eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==
X-Aux-Date: Mon, 19 Jan 2026 07:45:23 +0000
X-Signature: your_signature

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "X-API-Key": []string{"eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ=="},
        "X-Aux-Date": []string{"Mon, 19 Jan 2026 07:45:23 +0000"},
        "X-Signature": []string{"your_signature"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "https://openapi.dnse.com.vn/instruments", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

```javascript

const headers = {
  'Accept':'application/json',
  'X-API-Key':'eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==',
  'X-Aux-Date':'Mon, 19 Jan 2026 07:45:23 +0000',
  'X-Signature':'your_signature'
};

fetch('https://openapi.dnse.com.vn/instruments?symbol=SSI%2CSHS%2CACB%2CHUT%2CDSE',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'X-API-Key': 'eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==',
  'X-Aux-Date': 'Mon, 19 Jan 2026 07:45:23 +0000',
  'X-Signature': 'your_signature'
}

r = requests.get('https://openapi.dnse.com.vn/instruments', params={
  'symbol': 'SSI,SHS,ACB,HUT,DSE'
}, headers = headers)

print(r.json())

```

```java
URL obj = new URL("https://openapi.dnse.com.vn/instruments?symbol=SSI%2CSHS%2CACB%2CHUT%2CDSE");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

> Example responses

> 200 Response

```json
{
  "data": [
    {
      "symbol": "ACB",
      "marketId": "STO",
      "securityGroupId": "ST",
      "symbolType": "",
      "listedDate": "2020-12-09",
      "shortName": "Ngân hàng Á Châu",
      "name": "Ngân hàng TMCP Á Châu",
      "indexName": [
        "VN100",
        "VN30"
      ]
    }
  ],
  "total": 5,
  "page": 1,
  "pageSize": 100
}
```

<h3 id="getinstruments-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» data|[object]|false|none|Danh sách thông tin mã chứng khoán|
|»» symbol|string|false|none|Mã chứng khoán|
|»» marketId|string|false|none|Mã thị trường niêm yết<br>- STO: Cổ phiếu sàn HOSE<br>- STX: Cổ phiếu sàn HNX<br>- UPX: Cổ phiếu sàn UPCOM<br>- DVX: Phái sinh<br>- HCX: Trái phiếu doanh nghiệp|
|»» securityGroupId|string|false|none|Nhóm chứng khoán<br>- ST: Cổ phiếu<br>- EF: Quỹ ETF<br>- EW: Chứng quyền<br>- FU: Hợp đồng tương lai<br>- BS: Trái phiếu|
|»» symbolType|string|false|none|Phân loại mã hợp đồng phái sinh theo thời gian đáo hạn (áp dụng cho DERIVATIVE)<br>- VN30F1M: HĐTL chỉ số VN30 1 tháng<br>- VN30F2M: HĐTL chỉ số VN30 2 tháng<br>- VN30F1Q: HĐTL chỉ số VN30 1 quý<br>- VN30F2Q: HĐTL chỉ số VN30 2 quý|
|»» listedDate|string|false|none|Ngày niêm yết|
|»» shortName|string|false|none|Tên viết tắt của tổ chức phát hành|
|»» name|string|false|none|Tên đầy đủ của tổ chức phát hành|
|»» indexName|[string]|false|none|Danh sách chỉ số mà mã chứng khoán thuộc về (nếu có)|
|» total|integer(int32)|false|none|Tổng số bản ghi|
|» page|integer(int32)|false|none|Trang hiện tại (bắt đầu từ 1)|
|» pageSize|integer(int32)|false|none|Số bản ghi trên mỗi trang|

Status Code **400**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» code|string|false|none|none|
|» message|string|false|none|none|
|» status|integer|false|none|none|

Status Code **500**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» code|string|false|none|none|
|» message|string|false|none|none|
|» status|integer|false|none|none|


## Lịch sử OHLC

### Base URLs:
- **https://openapi.dnse.com.vn**

<span id="getOhlcHistory"></span>

### `GET /price/ohlc`

<h3 id="getohlchistory-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|symbol|query|string|true|Mã chứng khoán |
|type|query|string|true|Loại thị trường|
|resolution|query|string|true|Khung thời gian nến|
|from|query|string|true|Thời gian bắt đầu|
|to|query|string|true|Thời gian kết thúc|
|X-API-Key|header|string|true|API Key được cấp khi đăng ký dịch vụ|
|X-Aux-Date|header|string|true|Thời gian thực hiện yêu cầu|
|X-Signature|header|string|true|Chữ ký xác thực yêu cầu|

#### Detailed descriptions

**type**: Loại thị trường
- STOCK: Cổ phiếu
- DERIVATIVE: Phái sinh
- INDEX: Chỉ số thị trường

**resolution**: Khung thời gian nến
- 1,3,5,15,30,1h,1D,1W

> Code samples

```shell
# You can also use wget
curl -X GET https://openapi.dnse.com.vn/price/ohlc?symbol=ACB&type=STOCK&resolution=15&from=1773657310&to=1773830110 \
  -H 'Accept: application/json' \
  -H 'X-API-Key: eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==' \
  -H 'X-Aux-Date: Mon, 19 Jan 2026 07:45:23 +0000' \
  -H 'X-Signature: your_signature'

```

```http
GET https://openapi.dnse.com.vn/price/ohlc?symbol=ACB&type=STOCK&resolution=15&from=1773657310&to=1773830110 HTTP/1.1
Host: openapi.dnse.com.vn
Accept: application/json
X-API-Key: eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==
X-Aux-Date: Mon, 19 Jan 2026 07:45:23 +0000
X-Signature: your_signature

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "X-API-Key": []string{"eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ=="},
        "X-Aux-Date": []string{"Mon, 19 Jan 2026 07:45:23 +0000"},
        "X-Signature": []string{"your_signature"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "https://openapi.dnse.com.vn/price/ohlc", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

```javascript

const headers = {
  'Accept':'application/json',
  'X-API-Key':'eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==',
  'X-Aux-Date':'Mon, 19 Jan 2026 07:45:23 +0000',
  'X-Signature':'your_signature'
};

fetch('https://openapi.dnse.com.vn/price/ohlc?symbol=ACB&type=STOCK&resolution=15&from=1773657310&to=1773830110',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'X-API-Key': 'eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==',
  'X-Aux-Date': 'Mon, 19 Jan 2026 07:45:23 +0000',
  'X-Signature': 'your_signature'
}

r = requests.get('https://openapi.dnse.com.vn/price/ohlc', params={
  'symbol': 'ACB',  'type': 'STOCK',  'resolution': '15',  'from': '1773657310',  'to': '1773830110'
}, headers = headers)

print(r.json())

```

```java
URL obj = new URL("https://openapi.dnse.com.vn/price/ohlc?symbol=ACB&type=STOCK&resolution=15&from=1773657310&to=1773830110");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

> Example responses

> 200 Response

```json
{
  "t": [
    1773715500
  ],
  "o": [
    23.8
  ],
  "h": [
    23.8
  ],
  "l": [
    23.7
  ],
  "c": [
    23.75
  ],
  "v": [
    2530900
  ],
  "nextTime": 0
}
```

<h3 id="getohlchistory-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» t|[integer]|false|none|Danh sách thời gian nến|
|» o|[number]|false|none|Danh sách giá mở cửa của nến theo thời gian tương ứng|
|» h|[number]|false|none|Danh sách giá cao nhất trong nến theo thời gian tương ứng|
|» l|[number]|false|none|Danh sách giá thấp nhất trong nến theo thời gian tương ứng|
|» c|[number]|false|none|Danh sách giá đóng cửa của nến theo thời gian tương ứng|
|» v|[integer]|false|none|Danh sách khối lượng giao dịch theo thời gian tương ứng|
|» nextTime|integer(int32)|false|none|Timestamp của cây nến tiếp theo (nếu có), 0 nếu không còn|

Status Code **400**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» code|string|false|none|none|
|» message|string|false|none|none|
|» status|integer|false|none|none|

Status Code **500**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» code|string|false|none|none|
|» message|string|false|none|none|
|» status|integer|false|none|none|


## Dữ liệu khớp gần nhất

### Base URLs:
- **https://openapi.dnse.com.vn**

<span id="getLatestTrades"></span>

### `GET /price/{symbol}/trades/latest`

<h3 id="getlatesttrades-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|boardId|query|string|true|Mã bảng giao dịch|
|X-API-Key|header|string|true|API Key được cấp khi đăng ký dịch vụ|
|X-Aux-Date|header|string|true|Thời gian thực hiện yêu cầu|
|X-Signature|header|string|true|Chữ ký xác thực yêu cầu|
|symbol|path|string|true|Mã chứng khoán |

#### Detailed descriptions

**boardId**: Mã bảng giao dịch
- G1: Lô chẵn
- G4: Lô lẻ
- T1: Thỏa thuận trong giờ (9h - 14h45)
- T3: Thỏa thuận sau giờ (14h45 - 15h)
- T4: Thỏa thuận lô lẻ trong giờ (9h - 14h45)
- T6: Thỏa thuận lô lẻ sau giờ  (14h45 - 15h)

> Code samples

```shell
# You can also use wget
curl -X GET https://openapi.dnse.com.vn/price/{symbol}/trades/latest?boardId=G1 \
  -H 'Accept: application/json' \
  -H 'X-API-Key: eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==' \
  -H 'X-Aux-Date: Mon, 19 Jan 2026 07:45:23 +0000' \
  -H 'X-Signature: your_signature'

```

```http
GET https://openapi.dnse.com.vn/price/{symbol}/trades/latest?boardId=G1 HTTP/1.1
Host: openapi.dnse.com.vn
Accept: application/json
X-API-Key: eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==
X-Aux-Date: Mon, 19 Jan 2026 07:45:23 +0000
X-Signature: your_signature

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "X-API-Key": []string{"eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ=="},
        "X-Aux-Date": []string{"Mon, 19 Jan 2026 07:45:23 +0000"},
        "X-Signature": []string{"your_signature"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "https://openapi.dnse.com.vn/price/{symbol}/trades/latest", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

```javascript

const headers = {
  'Accept':'application/json',
  'X-API-Key':'eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==',
  'X-Aux-Date':'Mon, 19 Jan 2026 07:45:23 +0000',
  'X-Signature':'your_signature'
};

fetch('https://openapi.dnse.com.vn/price/{symbol}/trades/latest?boardId=G1',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'X-API-Key': 'eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==',
  'X-Aux-Date': 'Mon, 19 Jan 2026 07:45:23 +0000',
  'X-Signature': 'your_signature'
}

r = requests.get('https://openapi.dnse.com.vn/price/{symbol}/trades/latest', params={
  'boardId': 'G1'
}, headers = headers)

print(r.json())

```

```java
URL obj = new URL("https://openapi.dnse.com.vn/price/{symbol}/trades/latest?boardId=G1");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

> Example responses

> 200 Response

```json
{
  "trades": [
    {
      "marketId": "STO",
      "boardId": "G1",
      "isin": "VN000000ACB8",
      "symbol": "ACB",
      "matchPrice": 23.35,
      "matchQtty": 20,
      "side": "SELL",
      "avgPrice": 23.435,
      "totalVolumeTraded": 427980,
      "grossTradeAmount": 100.295515,
      "highestPrice": 23.55,
      "lowestPrice": 23.35,
      "openPrice": 23.5,
      "time": "2026-03-19 04:08:29"
    }
  ]
}
```

<h3 id="getlatesttrades-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» trades|[object]|false|none|none|
|»» marketId|string|false|none|Mã thị trường niêm yết mã chứng khoán<br>- DVX: Phái sinh sàn HNX<br>- HCX: Trái phiếu doanh nghiệp HNX<br>- STO: Cổ phiếu sàn HOSE<br>- STX: Cổ phiếu sàn HNX<br>- UPX: Cổ phiếu sàn Upcom|
|»» boardId|string|false|none|Mã bảng giao dịch<br>- G1: Lô chẵn<br>- G4: Lô lẻ<br>- T1: Thỏa thuận trong giờ (9h - 14h45)<br>- T3: Thỏa thuận sau giờ (14h45 - 15h)<br>- T4: Thỏa thuận lô lẻ trong giờ (9h - 14h45)<br>- T6: Thỏa thuận lô lẻ sau giờ (14h45 - 15h)|
|»» isin|string|false|none|Mã định danh quốc tế|
|»» symbol|string|false|none|Mã chứng khoán|
|»» matchPrice|number(double)|false|none|Giá khớp lệnh của giao dịch|
|»» matchQtty|integer(int32)|false|none|Khối lượng khớp của giao dịch|
|»» side|string|false|none|Phía giao dịch chủ động<br>- BUY: Bên mua chủ động<br>- SELL: Bên bán chủ động|
|»» avgPrice|number(double)|false|none|Giá khớp trung bình|
|»» totalVolumeTraded|integer(int32)|false|none|Tổng khối lượng đã giao dịch trong ngày|
|»» grossTradeAmount|number(double)|false|none|Tổng giá trị giao dịch trong nngày|
|»» highestPrice|number(double)|false|none|Giá khớp cao nhất trong ngày|
|»» lowestPrice|number(double)|false|none|Giá khớp thấp nhất trong ngày|
|»» openPrice|number(float)|false|none|Giá mở cửa|
|»» time|string|false|none|Thời gian khớp lệnh (theo định dạng yyyy-MM-dd HH:mm:ss)|

Status Code **400**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» code|string|false|none|none|
|» message|string|false|none|none|
|» status|integer|false|none|none|

Status Code **500**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» code|string|false|none|none|
|» message|string|false|none|none|
|» status|integer|false|none|none|


## Lịch sử khớp lệnh

### Base URLs:
- **https://openapi.dnse.com.vn**

<span id="getHistoryTrades"></span>

### `GET /price/{symbol}/trades`

<h3 id="gethistorytrades-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|boardId|query|string|true|Mã bảng giao dịch|
|from|query|string|true|Thời gian bắt đầu (timestamp)|
|to|query|string|true|Thời gian kết thúc (timestamp)|
|limit|query|number|false|none|
|X-API-Key|header|string|true|API Key được cấp khi đăng ký dịch vụ|
|X-Aux-Date|header|string|true|Thời gian thực hiện yêu cầu|
|X-Signature|header|string|true|Chữ ký xác thực yêu cầu|
|symbol|path|string|true|Mã chứng khoán|

#### Detailed descriptions

**boardId**: Mã bảng giao dịch
- G1: Lô chẵn
- G4: Lô lẻ
- T1: Thỏa thuận trong giờ (9h - 14h45)
- T3: Thỏa thuận sau giờ (14h45 - 15h)
- T4: Thỏa thuận lô lẻ trong giờ (9h - 14h45)
- T6: Thỏa thuận lô lẻ sau giờ  (14h45 - 15h)

> Code samples

```shell
# You can also use wget
curl -X GET https://openapi.dnse.com.vn/price/{symbol}/trades?boardId=G1&from=1773282637&to=1773289837 \
  -H 'Accept: application/json' \
  -H 'X-API-Key: eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==' \
  -H 'X-Aux-Date: Mon, 19 Jan 2026 07:45:23 +0000' \
  -H 'X-Signature: your_signature'

```

```http
GET https://openapi.dnse.com.vn/price/{symbol}/trades?boardId=G1&from=1773282637&to=1773289837 HTTP/1.1
Host: openapi.dnse.com.vn
Accept: application/json
X-API-Key: eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==
X-Aux-Date: Mon, 19 Jan 2026 07:45:23 +0000
X-Signature: your_signature

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "X-API-Key": []string{"eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ=="},
        "X-Aux-Date": []string{"Mon, 19 Jan 2026 07:45:23 +0000"},
        "X-Signature": []string{"your_signature"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "https://openapi.dnse.com.vn/price/{symbol}/trades", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

```javascript

const headers = {
  'Accept':'application/json',
  'X-API-Key':'eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==',
  'X-Aux-Date':'Mon, 19 Jan 2026 07:45:23 +0000',
  'X-Signature':'your_signature'
};

fetch('https://openapi.dnse.com.vn/price/{symbol}/trades?boardId=G1&from=1773282637&to=1773289837',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'X-API-Key': 'eyJvcmciOiJkbnNlIiwiaWQiOiI5YmMzYmViN2JjY2U0MmE0Yjk1NDE0MTA2YTMzODIxNyIsImgiOiJtdXJtdXIxMjgifQ==',
  'X-Aux-Date': 'Mon, 19 Jan 2026 07:45:23 +0000',
  'X-Signature': 'your_signature'
}

r = requests.get('https://openapi.dnse.com.vn/price/{symbol}/trades', params={
  'boardId': 'G1',  'from': '1773282637',  'to': '1773289837'
}, headers = headers)

print(r.json())

```

```java
URL obj = new URL("https://openapi.dnse.com.vn/price/{symbol}/trades?boardId=G1&from=1773282637&to=1773289837");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

> Example responses

> 200 Response

```json
{
  "trades": [
    {
      "marketId": "STO",
      "boardId": "G1",
      "isin": "VN000000ACB8",
      "symbol": "ACB",
      "matchPrice": 22.85,
      "matchQtty": 10,
      "side": "UNSPECIFIED",
      "avgPrice": 22.958,
      "totalVolumeTraded": 764430,
      "grossTradeAmount": 175.49849,
      "highestPrice": 23.25,
      "lowestPrice": 22.75,
      "openPrice": 23,
      "time": "2026-03-12 04:29:57"
    }
  ],
  "nextPageToken": "NDYwMTQ1XzIwMjYtMDMtMTJUMDQ6MjU6MDUuNTZa"
}
```

<h3 id="gethistorytrades-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» trades|[object]|false|none|Danh sách giao dịch|
|»» marketId|string|false|none|Mã thị trường|
|»» boardId|string|false|none|Mã bảng giao dịch|
|»» isin|string|false|none|Mã định danh quốc tế|
|»» symbol|string|false|none|Mã chứng khoán|
|»» matchPrice|number(double)|false|none|Giá khớp gần nhất|
|»» matchQtty|integer(int32)|false|none|Khối lượng khớp gần nhất|
|»» side|string|false|none|Chiều giao dịch|
|»» avgPrice|number(double)|false|none|Giá khớp trung bình|
|»» totalVolumeTraded|integer(int32)|false|none|Tổng khối lượng giao dịch trong ngày|
|»» grossTradeAmount|number(double)|false|none|Tổng giá trị giao dịch trong ngày|
|»» highestPrice|number(float)|false|none|Giá cao nhất trong ngày|
|»» lowestPrice|number(float)|false|none|Giá thấp nhất trong ngày|
|»» openPrice|integer(int32)|false|none|Giá mở cửa|
|»» time|string|false|none|Thời gian ghi nhận|
|» nextPageToken|string|false|none|Token dùng để lấy trang tiếp theo|

Status Code **400**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» code|string|false|none|none|
|» message|string|false|none|none|
|» status|integer|false|none|none|

Status Code **500**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|» code|string|false|none|none|
|» message|string|false|none|none|
|» status|integer|false|none|none|


---
title: "DNSE Market Data API - Tick Stats Per Side (Volume Profile)"
description: "Documentation for GetKrxTicksStatsPerSideBySymbol GraphQL endpoint to retrieve Volume Profile grouped by price level."
---

# DNSE Tick Stats Per Side (Volume Profile)

This endpoint is used to retrieve the "Bước giá" (Volume Profile) data. It returns the total accumulated volume traded at each price level, split by Buy Volume and Sell Volume, for a given symbol on a specific date.

## Endpoint Context

- **Endpoint:** `https://api.dnse.com.vn/price-api/query`
- **Method:** `POST`
- **Headers Required:**
  - `content-type: application/json`
  - `origin: https://entradex.dnse.com.vn`

## Request Payload

This API uses a GraphQL structure within a standard JSON body.
You must provide the `symbol`, `date` (format: `YYYY-MM-DD`), and `board` (e.g., `2` for VN30F).

```json
{
  "operationName": "GetKrxTicksStatsPerSideBySymbol",
  "query": "query GetKrxTicksStatsPerSideBySymbol { GetKrxTicksStatsPerSideBySymbol(symbol: \"41I1G4000\", date: \"2026-04-06\", board: 2) { price accumulatedVol buyVol sellVol unknownVol } }",
  "variables": {}
}
```

### Curl Example

```bash
curl 'https://api.dnse.com.vn/price-api/query' \
  -H 'content-type: application/json' \
  -H 'origin: https://entradex.dnse.com.vn' \
  --data-raw '{"operationName":"GetKrxTicksStatsPerSideBySymbol","query":"\n      query GetKrxTicksStatsPerSideBySymbol {\n        GetKrxTicksStatsPerSideBySymbol(symbol: \"41I1G4000\" , date: \"2026-04-06\", board: 2){\n            price\n            accumulatedVol\n            buyVol\n            sellVol\n            unknownVol\n        }\n      }\n      ","variables":{}}'
```

## Response Structure

The API returns a JSON object. Inside `data.GetKrxTicksStatsPerSideBySymbol`, it provides an array of price level objects. 

**Note: The array does not guarantee sorting.** If you are displaying this in an Order Book UI, you should sort the array by `price` descending to match typical Depth of Market (DOM) formats.

```json
{
  "data": {
    "GetKrxTicksStatsPerSideBySymbol": [
      {
        "price": 1854.5,
        "accumulatedVol": 71,
        "buyVol": 71,
        "sellVol": 0,
        "unknownVol": 0
      },
      ...
    ]
  }
}
```

### Field Definitions

- `price`: The price level where trades occurred.
- `accumulatedVol`: The total volume matched at this specific price level during the given date.
- `buyVol`: The portion of `accumulatedVol` that was initiated by proactive Market Buy orders (Khớp chủ động Mua).
- `sellVol`: The portion of `accumulatedVol` that was initiated by proactive Market Sell orders (Khớp chủ động Bán).
- `unknownVol`: Volume where the aggressor side could not be reliably determined.

## Client-side Handling Example (React/TypeScript)

When implementing the Volume Profile ("Bước giá") in a frontend client, you can use the following snippet to fetch and sort the data:

```typescript
const fetchVolumeProfile = async (symbol: string) => {
    const d = new Date();
    // Logic for weekend fallback in live market context
    if (d.getDay() === 0) d.setDate(d.getDate() - 2);
    else if (d.getDay() === 6) d.setDate(d.getDate() - 1);
    
    const dateStr = d.toISOString().split('T')[0];

    const queryPayload = {
        operationName: "GetKrxTicksStatsPerSideBySymbol",
        query: `query GetKrxTicksStatsPerSideBySymbol { GetKrxTicksStatsPerSideBySymbol(symbol: "${symbol}" , date: "${dateStr}", board: 2){ price accumulatedVol buyVol sellVol unknownVol } }`,
        variables: {}
    };

    const res = await fetch('https://api.dnse.com.vn/price-api/query', {
        method: 'POST',
        headers: { 'content-type': 'application/json', 'origin': 'https://entradex.dnse.com.vn' },
        body: JSON.stringify(queryPayload)
    });
    
    const json = await res.json();
    if (json?.data?.GetKrxTicksStatsPerSideBySymbol) {
        // Sort descending by price to match orderbook standards
        const sorted = json.data.GetKrxTicksStatsPerSideBySymbol.sort((a, b) => b.price - a.price);
        return sorted;
    }
    return [];
};
```


