var express = require('express');
var app = express();

app.get("/", function(req, res) {
	res.send("Hello From Docker Container !");
});

app.listen(8080, console.log("server is running on localhost:8080"));