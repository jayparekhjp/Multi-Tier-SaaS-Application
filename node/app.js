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


app.get('/', function (req, res) {
   var cookies = parseCookies(req);  
   var name = cookies.username;
   console.log(name);
   res.render('home',{
    name : 'hello'
   });
})  
 
// Jay Parekh-Users ----------------------------------------------------------------------------------------------------
app.get('/users/ping',function(req,res){
    console.log('Ping Called');
    client.get("http://13.56.115.80:3000/api/ping", function (data, response) {
    // client.methods.pingUserAPI(function (data, response) {
        console.log(data);
        res.send(data);
    });
});

app.get('/login',function(req,res){
    res.render('login');
});

app.get('/login2',function(req,res){
    res.render('login2');
});

app.get('/signup2',function(req,res){
    res.render('signup2');
});

app.post('/users/loginSubmit',function(req,res){
    var args = {
        data: { 
            username: req.body.username,
            password: req.body.password 
        },
        headers: { "Content-Type": "application/json" }
    };
    // console.log(args);
    client.post("http://http://13.56.115.80:3000/api/users/login", args, function (data, response) {
        console.log(data);
        // res.send(data);
        if (data){
            var cookies = new Cookies(req, res, { keys: keys})
            cookies.set('userid', data, {signed: true})
            res.redirect('/restraunts');
        }
    });
    // console.log(req.body.username);
});

// app.get('/admin/login',function(req,res){
//     // res.send('Admin Login Page Under Development.');
//     var users;
//     res.render('admin', {
//         title : 'Admin Dashboard',
//         users: users
//     });
// });

app.get('/users/signup',function(req,res){
    res.render('signup', {
        title: 'Signup | Counter Burger'
    });
});

app.post('/users/signupSubmit',function(req,res){
    var args = {
        data: { username: req.body.username,
            name: req.body.name,
            email: req.body.email,
            contact: req.body.contact,
            address: req.body.address,
            password: req.body.password 
        },
        headers: { "Content-Type": "application/json" }
    };
    console.log(args);
    client.post("http://http://13.56.115.80:3000/api/users/signup", args, function (data, response) {
        // parsed response body as js object
        console.log(data);
        if (data == 1){
            res.send('Username Already Taken');
        }else if (data){
            var cookies = new Cookies(req, res, { keys: keys})
            cookies.set('userid', data, {signed: true})
            res.redirect('/');
        }
    });
});

app.put('/users/changePassword',function(req,res){
    res.send('ChangePassword Page Under Development.');
});

app.delete('/users/deleteUser',function(req,res){
    res.send('DeleteUser Page Under Development.');
}); 
// Jay Parekh - Users ---------------------------------------------------------------------------------------------------

/*app.get('/restraunts', function (req, res) {
    var cookies = parseCookies(req);  
   var name = cookies.username;
   console.log(name);
   res.render('search',{
    name : name
   });
})  */

app.get('/restraunts', function (req, res) {
  var client = new Client();
  var pin = req.query.pin;
  var args = {
      parameters: { "zip": pin } // request headers
  };
  // console.log(pin)
  var cookies = new Cookies(req, res, { keys: keys })
  // var userid = cookies.get('userid');
  var userid = cookies.get('userid', { signed: true })
  console.log(userid);
  if(userid === undefined){
    res.redirect('/login');
  }
  client.get("http://localhost:3000/restraunts",args, function (data, response) { // CHANGE to broadcsat address for docker
      console.log(data);
      res.render('search',{
        data:data
      });
  });
});


app.get('/menu', function (req, res) {
  var client = new Client();
  var res_id = req.query.id;
  var args = {
      parameters: { "restraunt_id": res_id } // request headers
  };
  var cookies = new Cookies(req, res, { keys: keys })
<<<<<<< HEAD
  cookies.set('restraunt_id', res_id, { signed: true })

=======
  cookies.set('restraunt_id', res_id, { signed: false })
  var userid = cookies.get('userid', { signed: true })
  // console.log(userid)
>>>>>>> master
  client.get("http://localhost:3000/menus",args, function (data, response) {
      console.log(userid);
      res.render('menu',{
        data:data,
        userid:userid
      });
  });

});

app.listen(port, function () {
  console.log("Server is running on "+ port +" port");
});


