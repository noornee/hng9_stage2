## Task Description
Create an (POST) api endoint that takes the following sample json:
`{ “operation_type”: Enum <addition | subtraction | multiplication> , “x”: Integer, “y”: Integer }`
 Operation can either be addition, subtraction or mutiplication
 x can be a number and Integer datatype
 y can be a number and Integer datatype
 Based on the operation sent, perform a simple arithmetic operation on x and y
 Return a response with the result of the operation and your slack username
 `{ “slackUsername”: String, "operation_type" : Enum. value, “result”: Integer }`


## test live server
send a `POST` request with the required payload to
https://hng9stage1.herokuapp.com/calculate


## run server locally

### NOTE
make sure you have a PORT environment variable set
`export PORT="8000"`


go to the root directory of the project folder and execute

`go run *.go`

the server would run on port `:8000`

`curl localhost:8000/calculate -d '{"operation_type": "multiplication", "x": 10, "y": 10}'
` 

expected response:

`
{
    "operation_type": "multiplication",
    "result": 100,
    "slackUsername": "noornee"
}
`

# visual 
![Screenshot_20221105-093204](https://user-images.githubusercontent.com/71889751/200111200-42fe0102-28f2-4a60-8d8f-beeb24def2fc.png)
