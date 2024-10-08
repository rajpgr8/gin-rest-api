```
go run .

TEST:
Go to api-test folder:
1.
GET  get-event.http -> click on 'Send Request'
POST create-event.http -> click on 'Send Request'

2.
POST create-user.http -> click on 'Send Request'

3.
POST login.http -> click on 'Send Request'
```
```
Login API Response:
{
  "message": "Login successful!",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QyQGV4YW1wbGUuY29tIiwiZXhwIjoxNzI0NjY4MzI2LCJ1c2VySWQiOjB9.FoNXi6S8jrHPRM0_pQUmXtcyasAnOAxfDTGOVcDKc_E"
}
```

![image](https://github.com/user-attachments/assets/08487fbe-dbd2-492e-8dbb-5bbc08304168)



In authentication, when the user successfully logs in using their credentials, a JSON Web Token will be returned. 

Whenever the user wants to access a protected route or resource, the user agent should send the JWT, typically in the **Authorization header** using the **Bearer schema**. The content of the header should look like the following:
```
Authorization: Bearer <jwt-token>
```
#### Deploy to K8S
```
kubectl apply -f Deployment.yaml
kubectl apply -f Service.yaml
kubectl port-forward svc/gin-rest-api-svc 8080:8080

Test:
Go to api-test folder, and test the endpoints.
```
