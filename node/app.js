var express=require("express");
var bodyParser=require('body-parser');
const port = 8012;
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
    var cookies = new Cookies(req, res, { keys: keys})
    var client = new Client();
    var args = {
        data: { "id": cookies.get('userid', { signed: true })  },
        headers: {"Content-Type":"application/json"}
    };
    var userid = cookies.get('userid', { signed: true })
    if(userid === undefined){
      res.redirect('/login');
    }
    client.get("http://Project-132974579.us-west-2.elb.amazonaws.com:3000/api/cart/itemDisplay",args, function (data, response) {
        if(data  === null){
          data = [];
        }
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
  // console.log(args);

  client.delete("http://Project-132974579.us-west-2.elb.amazonaws.com:3000/api/cart/cartDelete", args, function (data, response) {
          res.redirect('/viewcart');
      }
  );
});


app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.get('/', function (req, res) {
   var cookies = parseCookies(req);  
   var name = cookies.username;
   res.render('home',{
    name : 'hello'
   });
})  
 
// Jay Parekh-Users ----------------------------------------------------------------------------------------------------

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
    // console.log(args);
    client.post("http://cmpe281-1995605336.us-west-1.elb.amazonaws.com:3000/api/users/login", args, function (data, response) {
        // console.log(data);
        // res.send(data);
        if (data == "Password Incorrect"){
            res.redirect('/login');
        }else if (data){    
            var cookies = new Cookies(req, res, { keys: keys})
            cookies.set('userid', data, {signed: true})
            res.redirect('/restraunts');
        }
    });
});


app.get('/signup',function(req,res){
    res.render('signup', {
        title: 'Signup | Counter Burger'
    });
});

app.post('/users/signupSubmit',function(req,res){
    // if (req.password == req.confirmPassword){
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
        // console.log(args);
        client.post("http://cmpe281-1995605336.us-west-1.elb.amazonaws.com:3000/api/users/signup", args, function (data, response) {
            // console.log(data);
            if (data){
                var cookies = new Cookies(req, res, { keys: keys})
                cookies.set('userid', data, {signed: true})
                res.redirect('/restraunts');
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
      parameters: { "zip": pin }
  };
  var cookies = new Cookies(req, res, { keys: keys })
  var userid = cookies.get('userid', { signed: true })
  if(userid === undefined){
    res.redirect('/login');
  }
  client.get("http://GOAPI-1977895044.us-west-1.elb.amazonaws.com:3000/restraunts",args, function (data, response) { // CHANGE to broadcsat address for docker
      // console.log(data);
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
  var userid = cookies.get('userid', { signed: true })
  if(userid === undefined){
    res.redirect('/login');
  }
  cookies.set('restraunt_id', res_id, { signed: false })
  var userid = cookies.get('userid', { signed: true })
  client.get("http://GOAPI-1977895044.us-west-1.elb.amazonaws.com:3000/menus",args, function (data, response) {
      // console.log(userid);
      res.render('menu',{
        data:data,
        userid:userid
        });
    });
});


app.post('/summary', function (req, res) {
    var client = new Client();
    var cookies = new Cookies(req, res, { keys: keys })
    var userid = cookies.get('userid', { signed: true })
    if(userid === undefined){
      res.redirect('/login');
    }
    var args = {
    data: {
     "id":cookies.get('userid', { signed: true })
     },
     headers: {"Content-Type":"application/json"}
    };
    client.get("http://Project-132974579.us-west-2.elb.amazonaws.com:3000/api/cart/itemDisplay",args, function (data, response) {
        var amount = 0;
        if(data !== null){
          for (var i = data.length - 1; i >= 0; i--) {
            amount+=data[i]['price']
          }
        }
        cookies.set('total', amount, {signed: true})
        res.render('summary',{
          data:data
        });
    });
});


app.post('/payment', function (req,res) {
  var cookies = new Cookies(req, res, { keys: keys })
  var userid = cookies.get('userid', { signed: true })
  if(userid === undefined){
    res.redirect('/login');
  }
  var client = new Client();
  res.render('payment',{ 
      userid : req.body.id,
      total : req.body.total, 
  });
});

app.post('/order', function (req, res) {
     var client = new Client();
     var cookies = new Cookies(req, res, { keys: keys })
     var userid = cookies.get('userid', { signed: true })
     var total = cookies.get('total', { signed: true })
      if(userid === undefined){
        res.redirect('/login');
      }
     //client.delete()
     // console.log(req.body)
     var args = {
        data: { 
            name: req.body.name,
            number: req.body.cardnumber,
            month: req.body.month,
            year: req.body.year,
            cvv: req.body.cvv,
            total: total,
            id: userid
        },
        headers: { "Content-Type": "application/json" }
    };
    client.post("http://cmpe281-1432011132.us-east-2.elb.amazonaws.com:3000/orders", args, function (data, response) {
          // console.log(data);
          res.render('insertAfterPayment',{
            data:data,
            userid : userid
          });
    }); 
});

app.get('/list', function (req, res) {
  var client = new Client();
  var cookies = new Cookies(req, res, { keys: keys })
  var userid = cookies.get('userid', { signed: true })
  if(userid === undefined){
    res.redirect('/login');
  }
  var args = {
        data: { 
        "id": userid
      },
      headers: {"Content-Type":"application/json"}
    };
  client.get("http://cmpe281-1432011132.us-east-2.elb.amazonaws.com:3000/details",args, function (data, response) { 
      // console.log(data);
      res.render('list',{
        data:data
      });
  });
});

app.get('/logout', function (req, res) {
  res.clearCookie("userid");
  res.redirect('/login');
})


app.listen(process.env.PORT || port, function () {
  console.log("Server is running on "+ port +" port");
});

