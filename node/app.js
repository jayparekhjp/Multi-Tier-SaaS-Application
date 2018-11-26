var express=require("express");
var bodyParser=require('body-parser');
const port = 8012;
// var connection = require('./config');
var app = express();
app.set('view engine', 'ejs');

var request = require('request');
var http = require('http');
var cookie = require('cookie');
var Client = require('node-rest-client').Client;


app.use(express.static(__dirname + '/public'));

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

app.use(bodyParser.urlencoded({extended:true}));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));


app.get('/', function (req, res) {
    var cookies = parseCookies(req);  
   var name = cookies.username;
   console.log(name);
   res.render('index',{
    name : name
   });
})  
 
app.get('/login', function (req, res) {  
   res.sendFile( __dirname + "/views/" + "login.html" );  
})  
 
app.get('/restraunts', function (req, res) {
    var cookies = parseCookies(req);  
   var name = cookies.username;
   console.log(name);
   res.render('search',{
    name : name
   });
})  

app.post('/search', function (req, res) {
  var client = new Client();
  var pin = req.body.pin;
  var args = {
      parameters: { "pin": pin } // request headers
  };
  client.get("http://demo8655652.mockable.io/restraunts",args, function (data, response) {
      // parsed response body as js object
      // var data = JSON.parse(data);
      console.log(data);
      console.log(data["restraunts"][0]["name"]);
      res.render('search',{
        data:data["restraunts"]
      });
  });
   // console.log(req.body.pin);

})  
/* route to handle login and registration */
// app.post('/api/register',registerController.register);
// app.post('/api/authenticate',authenticateController.authenticate);
 
// console.log(authenticateController);
// app.post('/controllers/register-controller', registerController.register);
// app.post('/controllers/authenticate-controller', authenticateController.authenticate);
app.listen(port, function () {
  console.log("Server is running on "+ port +" port");
});


