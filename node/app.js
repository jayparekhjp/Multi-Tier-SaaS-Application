var express=require("express");
var bodyParser=require('body-parser');
const port = 8012;
// var connection = require('./config');
var app = express();
app.use(express.static(__dirname + '/public'));

// var authenticateController=require('./controllers/authenticate-controller');
// var registerController=require('./controllers/register-controller');
 
app.use(bodyParser.urlencoded({extended:true}));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.get('/', function (req, res) {  
   res.sendFile( __dirname + "/views/" + "index.html" );  
})  
 
app.get('/login', function (req, res) {  
   res.sendFile( __dirname + "/views/" + "login.html" );  
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
