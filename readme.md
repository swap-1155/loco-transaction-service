# LOCO.GG - Assessment - Backend Developer

This is code base created for an interview for loco.gg
The code is written using GoLang, DB used is mySql.
The code is ment to run on a local system only, using postman or other API tools.

### Things to consider changing on different devices.

1. in main.go
   - 23: router.Run("localhost:8080") 
   - this is the local host endpoint that will be used for API calls, this can be changed to ome remote API endpoint.

2. in util\connectToDB\ConnectToDB.go
   - 25: conString := "root:password@tcp(localhost:3306)/testing" 
   - ref - "{user}:{password}@tcp({db_hostname:port})/{schema}"
   - this represnts the local DB connect and can be updated to some remote value as well.

