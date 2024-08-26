##### We want to restrict event creation to logged-in users only, ensuring that only authenticated users can create events.
Once user is logged in a JWT (JSON Web Token) will be returned. Now if user wants to create an event (protected route), the JWT should be sent in the request header. Like ```--header 'Authorization: <JWT-TOKEN>'```
```
go run .
```

```
TEST: Go to api-test folder:
1.
POST login.http -> click on 'Send Request'

=>
Login API Response:
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 26 Aug 2024 09:54:14 GMT
Content-Length: 199
Connection: close

{
  "message": "Login successful!",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QyQGV4YW1wbGUuY29tIiwiZXhwIjoxNzI0NjczMjU0LCJ1c2VySWQiOjB9.yj7Sr04IwglYwmmprIcKKLaiRZ0WMBwm5xugkeV0Es8"
}
```

```
2. Try to create an event without passing JWT or with wrong JWT
POST login.http -> click on 'Send Request' (without passing authorization header)

=>
Login API Response:
HTTP/1.1 401 Unauthorized
Content-Type: application/json; charset=utf-8
Date: Mon, 26 Aug 2024 09:56:41 GMT
Content-Length: 29
Connection: close

{
  "message": "Not authorized."
}
```

```
3. Create an event with JWT token got from previous login request (See point 1 above). 
POST login.http -> click on 'Send Request'
Example:
POST http://localhost:8080/events
content-type: application/json
authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QyQGV4YW1wbGUuY29tIiwiZXhwIjoxNzI0NjczMDg5LCJ1c2VySWQiOjB9.3lcSTvX0sxYmLmI6j36gWm0r6rNFXnTFAmsXqzKqWQA

=>
Login API Response:
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Mon, 26 Aug 2024 09:58:09 GMT
Content-Length: 176
Connection: close

{
  "event": {
    "ID": 0,
    "Name": "Test event 2",
    "Description": "Another test event",
    "Location": "A test location",
    "DateTime": "2025-01-01T15:30:00Z",
    "UserID": 0
  },
  "message": "Event created!"
}

You can also use CURL for this POST request:

curl --location 'http://localhost:8080/events' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QyQGV4YW1wbGUuY29tIiwiZXhwIjoxNzI0NjczMDg5LCJ1c2VySWQiOjB9.3lcSTvX0sxYmLmI6j36gWm0r6rNFXnTFAmsXqzKqWQA' \
--header 'Content-Type: application/json' \
--data '{
  "name": "Test event 2",
  "description": "Another test event",
  "location": "A test location",
  "dateTime": "2025-01-01T15:30:00.000Z"
}'
```

![image](https://github.com/user-attachments/assets/075309e7-6e37-4c7e-88a4-4ea75b032022)

In authentication, when the user successfully logs in using their credentials, a JSON Web Token will be returned. 

Whenever the user wants to access a protected route or resource, the user agent should send the JWT, typically in the **Authorization header** using the **Bearer schema**. The content of the header should look like the following:
```
Authorization: Bearer <jwt-token>
```
