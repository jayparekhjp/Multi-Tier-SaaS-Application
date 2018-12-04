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
// client.registerMethod("pingUserAPI", "http://localhost:3000/api/ping", "GET");
// client.registerMethod("login", "http://localhost:3000/api/users/login/", "GET");
client.registerMethod("admin", "http://localhost:3000/api/users", "GET");
client.registerMethod("signup", "http://localhost:3000/api/users/signup", "POSt");

app.get('/',function(req,res){
    console.log('Home Page Called.');
    res.send('Home Page');
});

app.get('/users/ping',function(req,res){
    console.log('Ping Called');
    client.get("http://localhost:3000/api/ping", function (data, response) {
    // client.methods.pingUserAPI(function (data, response) {
        console.log(data);
        res.send(data);
    });
});

app.get('/users/login',function(req,res){
    res.render('login', {
        title: 'Login | Counter Burger'
    });
});

app.post('/users/loginSubmit',function(req,res){
    var existingUser = {
        username: req.body.username,
        password: req.body.password
    }
    console.log(existingUser);
    // client.methods.login(existingUser, function (data, response) {
    client.get("http://localhost:3000/api/users/login/", function (data, response) {
        console.log(data);
        res.send(data);
    });
    console.log(req.body.username);
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
    res.render('signup', {
        title: 'Signup | Counter Burger'
    });
});

app.post('/users/signupSubmit',function(req,res){
    // if req.body.password != req.body.confirmPassword (
    //     res.send('Password and Confirm Password needs to be same.')
    // )
    // var newUser = {
    //     username: req.body.username,
    //     name: req.body.name,
    //     email: req.body.email,
    //     contact: req.body.contact,
    //     address: req.body.address,
    //     password: req.body.password
    // }
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
    var message;
    client.post("http://localhost:3000/api/users/signup", args, function (data, response) {
        // parsed response body as js object
        console.log(data);
        res.send(data);
    });
});

app.put('/users/changePassword',function(req,res){
    res.send('ChangePassword Page Under Development.');
});

app.delete('/users/deleteUser',function(req,res){
    res.send('DeleteUser Page Under Development.');
});


app.listen('5000');
console.log('App listening on port 5000');