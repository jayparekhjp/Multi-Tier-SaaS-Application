var express=require("express");
var bodyParser=require('body-parser');
const port = 8000;
var app = express();
var Client = require('node-rest-client').Client;
app.set('view engine', 'ejs');

var request = require('request');
var http = require('http');

app.use(express.static(__dirname + '/public'));

app.use(bodyParser.urlencoded({extended:true}));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.get('/', function (req, res) {

   res.render('index',{
   });
})  

app.get('/cart', function (req, res) {
    var client = new Client();
    client.get("http://demo8655652.mockable.io/menus", function (data, response) {
        // console.log(data[0]['name']);
        res.render('menu',{
          data:data
        });
    });
});

app.listen(port, function () {
  console.log("Server is running on "+ port +" port");
});
