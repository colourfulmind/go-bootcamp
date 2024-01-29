## Tasty Discoveries

Run "godoc" to access the following links.
- [Cmd Package](http://localhost:6060/pkg/main/cmd/app/)
- [Internal Package APIServer](http://localhost:6060/pkg/main/internal/app/apiserver)
- [Internal Package DB](http://localhost:6060/pkg/main/internal/app/db)
- [Internal Package ESClient](http://localhost:6060/pkg/main/internal/app/esclient)
- [Internal Package Server](http://localhost:6060/pkg/main/internal/app/server)
- [Internal Package Store](http://localhost:6060/pkg/main/internal/app/store)


### Elasticsearch Integration for Restaurant Dataset

A dataset of restaurants in Moscow, Russia, has been provided.
Elasticsearch was employed to create an index for the dataset,
and the dataset is uploaded using the Bulk API.

Each entry contains the following fields:

- ID
- Name
- Address
- Phone
- Longitude
- Latitude

Check Index:

```bash
~$ curl -s -XGET "http://localhost:9200/places"
```

Query Entry by ID:

```bash
~$ curl -s -XGET "http://localhost:9200/places/_doc/1"
```

### Simplest Interface

An HTML UI is created to interact with the database.
The page displays a list of names, addresses, and phone numbers, allowing users to view it in a browser.

The underlying database is abstracted behind an interface,
providing functionality to retrieve a list of entries and paginate through them.

The HTTP application is hosted on port 8888,
responding with a list of restaurants and offering simple pagination.
For example, querying "http://127.0.0.1:8888/?page=5" yields a page structure like the following:

```html
<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>Places</title>
    <meta name="description" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1">
</head>

<body>
<h5>Total: 13649</h5>
<ul>
    <li>
        <div>Teremok</div>
        <div>gorod Moskva, ulitsa Arbat, dom 32</div>
        <div>(495) 916-00-24</div>
    </li>
    <!-- ... other entries ... -->
    <li>
        <div>Stolovaja pri shkole 2097</div>
        <div>gorod Moskva, Aerodromnaja ulitsa, dom 9</div>
        <div>(495) 395-32-25</div>
    </li>
</ul>
<a href="/?page=1">First</a>
<a href="/?page=4">Previous</a>
<a href="/?page=6">Next</a>
<a href="/?page=1365">Last</a>
</body>
</html>
```

The "Previous" link disappears on the first page, and the "Next" link disappears on the last page.

In case the 'page' parameter is specified with a wrong value,
the page returns an HTTP 400 error and plain text with an error message:

```
Invalid 'page' value: 'foo'
```

### Proper API

Another handler is implemented that responds with Content-Type: application/json
and provides a JSON version of the same data.

For example, when accessing http://127.0.0.1:8888/api/places?page=11, the response should look like this:

```json
{
  "name": "Places",
  "total": 13649,
  "places": [
    {
      "id": 100,
      "name": "Bil'jardnyj Bar-klub Polkovnik i Bajron",
      "address": "gorod Moskva, prospekt Budennogo, dom 1/1",
      "phone": "(495) 365-22-24",
      "location": {
        "lon": 37.721607349974896,
        "lat": 55.7785020631572
      }
  }
  // other entries
  ],
  "prev_page": 10,
  "next_page": 12,
  "last_page": 1365
}
```

### Closest Restaurants

The next implementation is a key functionality - searching for the three closest restaurants.
To achieve this, query sorting was configured using Elasticsearch.
The sorting configuration:

```json
{
  "sort": [
    {
      "_geo_distance": {
        "location": {
          "lat": 55.674,
          "lon": 37.666
        },
        "order": "asc",
        "unit": "km",
        "mode": "min",
        "distance_type": "arc",
        "ignore_unmapped": true
      }
    }
  ]
}
```

In this configuration, "lat" and "lon" represent your current coordinates. 
For example, with the URL http://127.0.0.1:8888/api/recommend?lat=55.537&lon=37.722, 
the application returns JSON in the following format:

```json
{
  "name": "Recommendation",
  "places": [
    {
      "id": 5719,
      "name": "Kafe «Zagor'e»",
      "address": "gorod Moskva, 28-j kilometr Moskovskoj Kol'tsevoj Avtodorogi, vladenie 7",
      "phone": "(926) 680-24-45",
      "location": {
        "lon": 37.6741268859941,
        "lat": 55.5719712615328
      }
    },
    {
      "id": 5633,
      "name": "Kafe-zakusochnaja «Burger King»",
      "address": "gorod Moskva, 28-j kilometr Moskovskoj Kol'tsevoj Avtodorogi, vladenie 2",
      "phone": "(495) 916-56-16",
      "location": {
        "lon": 37.673778198821104,
        "lat": 55.5727021798031
      }
    },
    {
      "id": 2972,
      "name": "KONECh.ST.VODIT. 7 TROL.",
      "address": "gorod Moskva, Lipetskaja ulitsa, dom 19",
      "phone": "(495) 329-67-51",
      "location": {
        "lon": 37.686302196638735,
        "lat": 55.579274995602425
      }
    }
  ]
}
```

### JWT Authentication

To enhance security, a simple form of authentication using JSON Web Tokens (JWT) is implemented.

http://127.0.0.1:8888/api/get_token sole purpose is to generate a token and return it. 
The response format is as follows:

```bash
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJKV1QifQ.DeYqCCnuM-7srwUhk2jNIBr-tCoJmqbbS3RKmC9R_lQ"
}
```

By default, when querying this API from the browser without a valid token, it fails with an HTTP 401 error.
However, it works when an Authorization: Bearer header is specified by the client.
You can verify this using tools like cURL.

```
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJKV1QifQ.DeYqCCnuM-7srwUhk2jNIBr-tCoJmqbbS3RKmC9R_lQ" -XGET "http://127.0.0.1:8888/api/recommend?lat=55.674&lon=37.666"
```
