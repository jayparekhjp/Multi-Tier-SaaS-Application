var express = require('express');
var bodyParser = require('body-parser');
var path = require('path');
var Client = require('node-rest-client').Client;

var client = new Client();
var app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: false}));

app.use(express.static(path.join(__dirname, 'public')));

app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, './views'));


// User API routes and functions
client.registerMethod("pingUserAPI", "http://localhost:3000/api/ping", "GET");
client.registerMethod("login", "http://localhost:3000/api/users/login", "GET");
client.registerMethod("admin", "http://localhost:3000/api/users", "GET");

app.get('/',function(req,res){
    console.log('Home Page Called.');
    res.send('Home Page');
});

app.get('/users/ping',function(req,res){
    console.log('Ping Called');
    client.methods.pingUserAPI(function (data, response) {
        console.log(data);
        res.send(data);
    });
});

app.get('/users/login',function(req,res){
    res.render('login', {
        title: 'Login | Counter Burger'
    });
});

app.get('/users/loginSubmit',function(req,res){
    // client.methods.login(function (data, response) {
    //     console.log(data);
    //     res.send(data);
    // });
    res.send(res.body);
    console.log(res.body);
});

// Dummy user data
var users = [
    {
        id: 1,
        name: "jay"
    },
    {
        id: 2,
        name: "parekh",
    }
]

app.get('/admin/login',function(req,res){
    // res.send('Admin Login Page Under Development.');
    res.render('admin', {
        title : 'Admin Dashboard',
        users: users
    });
});

app.get('/users/signup',function(req,res){
    res.send('SignUp Page Under Development.');
});

app.post('/users/signupSubmit',function(req,res){
    res.send('SignUp Page Under Development.');
});

app.put('/users/changePassword',function(req,res){
    res.send('ChangePassword Page Under Development.');
});

app.delete('/users/deleteUser',function(req,res){
    res.send('DeleteUser Page Under Development.');
});


app.listen('5000');
console.log('App listening on port 5000');