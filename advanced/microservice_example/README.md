# A Simple Microservice Example 
This folder contains a simple microservice example. This microservice hosts a homepage and an about page displaying simple strings. Later, it is containerized using a Dockerfile. 


## Run the following commands 

<br>

``` 
docker build -t gowebapi . 
``` 

```
docker run -p 8080:8080 -tid gowebapi 
```

