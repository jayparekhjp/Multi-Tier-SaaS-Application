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

function parseCookies (request) {
    var list = {},
        rc = request.headers.cookie;

    rc && rc.split(';').forEach(function( cookie ) {
        var parts = cookie.split('=');
        list[parts.shift().trim()] = decodeURI(parts.join('='));
    });

    return list;
}


app.get('/viewcart', function (req, res) {
    var client = new Client();
    var cookies = new Cookies(req, res, { keys: keys })
    var args = {
        data: { "id": cookies.get('userid', { signed: true }) },
        headers: {"Content-Type":"application/json"}// request headers
    };
    client.get("http://34.219.121.214:3000/api/cart/itemDisplay",args, function (data, response) {
        res.render('cart',{
          data:data
        });
    });
});

/*app.post('/additem', function (req, res) {
    var client = new Client();
    var args = {
        data: { 
            id : req.body.userid,
            iid : req.body.itemid,
            rid : req.body.resid,
            res : req.body.res,
            iname : req.body.iname,
            price:parseFloat(req.body.price)
        },
        headers: {"Content-Type":"application/json"}// request headers
    };
    client.post("http://34.219.121.214:3000/api/cart/itemSave",args, function (data, response) {
                res.redirect('')
    }) 
    }
);
*/
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
  client.delete("http://34.219.121.214:3000/api/cart/cartDelete", args, function (data, response) {
      
          res.redirect('/viewcart');
      }
  );
});


//app.use(bodyParser.urlencoded({extended:true}));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));


app.post('/summary', function (req, res) {
    var client = new Client();
    var cookies = new Cookies(req, res, { keys: keys })
    var args = {
    data: {
     "id":cookies.get('userid', { signed: true })
     },
     headers: {"Content-Type":"application/json"}
    };
    client.get("http://34.219.121.214:3000/api/cart/itemDisplay",args, function (data, response) {
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

app.get('/', function (req, res) {
   var cookies = parseCookies(req);  
   var name = cookies.username;
   res.render('home',{
    name : 'hello'
   });
})  
 
// Jay Parekh-Users ----------------------------------------------------------------------------------------------------
app.get('/users/ping',function(req,res){
    console.log('Ping Called');
    client.get("http://cmpe281-1995605336.us-west-1.elb.amazonaws.com:3000/api/ping", function (data, response) {
    // client.methods.pingUserAPI(function (data, response) {
        console.log(data);
        res.send(data);
    });
});

app.get('/login',function(req,res){
    var cookies = new Cookies(req, res, { keys: keys })
    var userid = cookies.get('userid', { signed: true })
    if(userid !== undefined){
      res.redirect('/restraunts');
    }else{
      res.render('login');
    }
});

app.get('/login-test',function(req,res){
    res.render('login-test');
});

app.get('/signup-test',function(req,res){
    res.render('signup-test');
});

app.post('/users/loginSubmit',function(req,res){
    var args = {
        data: { 
            username: req.body.username,
            password: req.body.password 
        },
        headers: { "Content-Type": "application/json" }
    };
    console.log(args);
    client.post("http://cmpe281-1995605336.us-west-1.elb.amazonaws.com:3000/api/users/login", args, function (data, response) {
        console.log(data);
        // res.send(data);
        if (data == "Password Incorrect"){
            res.redirect('/login');
        }else if (data){    
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

app.get('/signup',function(req,res){
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
    client.post("http://cmpe281-1995605336.us-west-1.elb.amazonaws.com:3000/api/users/signup", args, function (data, response) {
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

app.get('/restraunts', function (req, res) {
  var client = new Client();
  var pin = req.query.pin;
  var args = {
      parameters: { "zip": pin } // request headers
  };
  // console.log(pin)
  var cookies = new Cookies(req, res, { keys: keys })
  var userid = cookies.get('userid', { signed: true })
  if(userid === undefined){
    res.redirect('/login');
  }
  client.get("http://GOAPI-1977895044.us-west-1.elb.amazonaws.com:3000/restraunts",args, function (data, response) { // CHANGE to broadcsat address for docker
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
  cookies.set('restraunt_id', res_id, { signed: false })
  // cookies.set('userid', 1, {signed: true})
  var userid = cookies.get('userid', { signed: true })
  client.get("http://GOAPI-1977895044.us-west-1.elb.amazonaws.com:3000/menus",args, function (data, response) {
      console.log(userid);
      res.render('menu',{
        data:data,
        userid:userid
        });
    });
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