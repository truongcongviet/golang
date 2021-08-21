## The Friend Management RESTFul Api

### Project Features

* Build Api enpoints for Friend management project
* Write unit test for API endpoints

### Project Dependencies
1. Golang (1.16.3)
2. Chi framework (v5.0.2)
3. Install and run PostgreSQL on your localhost for storing data
4. ...


### How to use from this sample project
##### Clone the repository
```
git clone https://github.com/quan12yt/friend-management-golang-restapi.git
cd friend-management-golang-restapi
```

##### Install dependencies 
```
cd friend-management-golang-restapi
make build
```

##### Run project
* Local
```
cd friend-management-golang-restapi
make run
```
* Docker
```
make docker-run
```


### RestApi Enpoints

````
1, Retrieve the friends list for an email address :  http://localhost:8080/api/friends
 *Example Request
    {
    "email" : "quan12yt@gmail.com"
    }
 *Success Response Example
    {
    "success": true,
    "friends": [
        "quang@gmail.com",
        "tonhut@gmail.com"
    ],
    "count": 2
    }
  *Error Response Example
   {
    "success": false,
    "text": "Invalid email format",
    "timestamp": "2021-05-06 14:20:59"
  }
 -------------------------------------------------------------
2, Create a friend connection : http://localhost:8080/api/add
  *Example Request
    {
      "friends":[
        "anh.tran@s3corp.com.vn",
        "chi.vo@s3corp.com.vn"
        ]
    }
  *Success Response Example
    {
      true
    }
  *Error Response Example
   {
    "success": false,
    "text": "2 emails are already being friend",
    "timestamp": "2021-05-06 14:21:44"
  }
  -------------------------------------------------------------
3, Retrieve the common friends list between two email addresses :  http://localhost:8080/api/common
  *Example Request
      {
    "friends" : [
        "quan12yt@gmail.com",
        "tonhut@gmail.com"
      ]
    }
 *Success Response Example
   {
    "success": true,
    "friends": [
        "quang@gmail.com"
    ],
    "count": 1
  }
 *Error Response Example
   {
    "success": false,
    "text": "invalid email format",
    "timestamp": "2021-05-06 14:23:41"
  }
  -------------------------------------------------------------
4, Create subscribe to updates from an email address : http://localhost:8080/api/subscribe
  *Example Request
    {
       "requestor": "quang@gmail.com",
        "target": "quan12yt@gmail.com"
  }
  *Success Response Example
   {
    "success": true
  }
  *Error Response Example
   {
    "success": false,
    "text": "invalid email format",
    "timestamp": "2021-05-06 14:23:41"
  }
  -------------------------------------------------------------
5, Block updates from an email address: http://localhost:8080/api//block
  *Example Request
    {
       "requestor": "quang@gmail.com",
        "target": "quan12yt@gmail.com"
  }
  *Success Response Example
   {
    "success": true
  }
  *Error Response Example
   {
    "success": false,
    "text": "invalid email format",
    "timestamp": "2021-05-06 14:23:41"
  }
  -------------------------------------------------------------
6, Create API to retrieve all email addresses that can receive updates from an email address :  http://localhost:8080/api/retrieve
  *Example Request
    {
    "sender": "quan12yt@gmail.com",
    "text" : "ahdad la@gmail.com ahdad la@gmail.com ahdad la@gmail.com"
  }
  *Success Response Example
    {
    "success": true,
    "recipients": [
        "la@gmail.com",
        "len@gmail.com",
        "quang@gmail.com",
        "tonhut@gmail.com"
    ]
  }
  *Error Response Example
   {
    "success": false,
    "text": "sender must not empty",
    "timestamp": "2021-05-06 14:27:42"
  }
````
