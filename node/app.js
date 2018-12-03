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
var http = require('http');
var cookie = require('cookie');


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
   res.render('/login');  
})  
 

// Mock Users login
// client.registerMethod("jsonMethod", "http://demo7713207.mockable.io/api/users", "GET");

// Mock Users signup
// client.registerMethod("jsonMethod", "http://demo7713207.mockable.io/api/users", "POST");

client.registerMethod("jsonMethod", "http://localhost:3000/api/ping", "GET");
 
client.methods.jsonMethod(function (data, response) {
    console.log(data);
});

/* route to handle login and registration */
// app.post('/api/register',registerController.register);
// app.post('/api/authenticate',authenticateController.authenticate);
 
// console.log(authenticateController);
// app.post('/controllers/register-controller', registerController.register);
// app.post('/controllers/authenticate-controller', authenticateController.authenticate);
app.listen(port, function () {
  console.log("Server is running on "+ port +" port");
});


