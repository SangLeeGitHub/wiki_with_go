### Wiki Web backend (Part 1)

#### Technology used : Go

#### 0. Install Go compiler 
`Install Go compiler from https://golang.org/dl`

#### 1. Cloning Git
`git clone https://github.com/hotdeveloper/wiki_with_go.git`

`cd wiki_with_go`

#### 2. Build a executable binary file
`go build`

#### 3. Running executable as background.
`./wiki_with_go &`

#### 4. Test the backend app
In CLI, if you have curl or wget command on your system,  
○ curl http://localhost:9090/articles/  
○ curl http://localhost:9090/articles/rest_api  
○ curl -X PUT http://localhost:9090/articles/wiki -d 'A wiki is a knowledge base website'  
