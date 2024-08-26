```
go run .

TEST:
Go to api-test folder.
GET  get-event.http -> click on 'Send Request'
POST create-event.http -> click on 'Send Request'
```

#### Deploy to K8S
```
kubectl apply -f Deployment.yaml
kubectl apply -f Service.yaml
kubectl port-forward svc/gin-rest-api-svc 8080:8080

Test:
Go to api-test folder, and test the endpoints.
```
