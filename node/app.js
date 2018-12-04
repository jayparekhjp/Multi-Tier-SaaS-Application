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

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));


app.get('/viewcart', function (req, res) {
    var client = new Client();
    var args = {
        data: { "id": "2" },
        headers: {"Content-Type":"application/json"}// request headers
    };
    client.get("http://localhost:3000/api/cart/itemDisplay",args, function (data, response) {
        res.render('cart',{
          data:data
        });
    });
});

app.post('/deleteitem',function(req,res){
  var args = {
      data: { 
          id : req.body.userid,
          iid : req.body.itemid,
          rid : req.body.resid 
      },
      headers: { "Content-Type": "application/json" }
  };
  console.log(args);
  client.delete("http://localhost:3000/api/cart/cartDelete", args, function (data, response) {
      
          res.redirect('/viewcart');
      }
  );
});


app.get('/payment', function (req,res) {
    //var client = new Client();
    //client.post("http://localhost:3002/orders", function (data, response) {
        //console.log(data[0]['name']);
        res.render('payment');
    });


app.post('/order', function (req, res) {
    var client = new Client();
    var args = {
        headers: { "Content-Type": "application/json" }
    };    
     client.post("http://localhost:3002/orders",args, function (data, response) {
      // console.log(data[0]['name']);
    });
});
    
app.listen(port, function () {
  console.log("Server is running on "+ port +" port");
});

