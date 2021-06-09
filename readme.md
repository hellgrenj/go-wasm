Playing around with Wasm in Go.  

You need:
* Go 1.16.5 or later  
* Node 15.11.0 or later (to use ``npx``)

Compile:  
```GOOS=js GOARCH=wasm go build -o main.wasm```

Run:  
```npx serve``` and then open localhost:5000 in your browser

