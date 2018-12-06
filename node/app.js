var express=require("express");
var bodyParser=require('body-parser');
const port = 8012;
// var connection = require('./config');
var app = express();
app.set('view engine', 'ejs');
app.set('views','./views');

var Client = require('node-rest-client').Client;
var client = new Client();

var request = require('request');
var Cookies = require('cookies')
var http = require('http');
var cookie = require('cookie');
var Client = require('node-rest-client').Client;
var keys = ['keyboard cat']


app.use(express.static(__dirname + '/public'));
app.use(bodyParser.urlencoded({ extended: true })); 


// var authenticateController=require('./controllers/authenticate-controller');
// var registerController=require('./controllers/register-controller');
function parseCookies (request) {
    var list = {},
        rc = request.headers.cookie;

    rc && rc.split(';').forEach(function( cookie ) {
        var parts = cookie.split('=');
        list[parts.shift().trim()] = decodeURI(parts.join('='));
    });

    return list;
}

//app.use(bodyParser.urlencoded({extended:true}));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));


app.get('/summary', function (req, res) {
    var client = new Client();
    var args = {
    data: {
     "id":"2"
     //"id":req.body.userid
     },
     headers: {"Content-Type":"application/json"}
    };
    client.get("http://34.216.22.59:3000/api/cart/itemDisplay",args, function (data, response) {
        // console.log(data[0]['name']);
        res.render('summary',{
          data:data
        });
    });
});


app.get('/payment', function (req,res) {
    //var client = new Client();
    //client.post("http://18.222.209.245:3002/orders", function (data, response) {
        //console.log(data[0]['name']);
        res.render('payment');
    });


app.post('/order', function (req, res) {
    var client = new Client();
     //client.delete()
         
     client.post("http://18.222.209.245:3002/orders", function (data, response) {
      // console.log(data[0]['name']);
    });
});
    

app.listen(port, function () {
  console.log("Server is running on "+ port +" port");
});


