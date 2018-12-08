# fa18-281-slack-overflow

## Work Distribution
|Team Member Name|Functionality|
|-|-|
|Jay Parekh|Users|
|Atish Maitreya|Menu|
|Jaykumar Patel|Cart|
|Yesha Dave|Payment|
|Wen Li|Orders|

## Technologies Used:
AWS Web Services
Heroku (Frond-end deployment)
Docker
Kong
Cassandra
GoLang
Node.js
Cassandra
MongoDB
session and cookie

## Go api Functionality
  Users: Create the registration and login functions for users, and allow the userId to be passed through all the requests sent by a certain user via session. The user information is stored in all the database.
  Menu: Restaurants in different locations have varied menus. For online orders, users can view menu via typing Zip code of the restaurants.  
  Cart: While viewing items in a menu, users can add items with a specific count and add to the cart for this restaurant. The cart information is stored in the counter database.  
  Payment: When users checking out items in shopping carts, they can make payment by typing user's payment information and then make a payment. After clicking the button 'make payment', the order history will be passed to the next goapi.
  Orders: After a user make a payment, the order history is kept in the buigerOrder database. At the same page, users can view his own order history by clicking on "List all your order history".

## node.js
The app.js is designed to send requests to goapis and render or redirect pages to separate ejs files.  

## ejs pages
admin.ejs: List all the user information and only available for admin user
cart.ejs: Show the cart for a certain user
description.ejs:
header.ejs: Gives an overview of how the website works and shows the users the navigation bar to home page, restaurants page and description page
home.ejs: Home page of our website
index.ejs:
insertAfterPayment.ejs: Shows the most recent payment of the user.
list.ejs: Shows all the payment history of the user.
loin.ejs: Login page
menu.ejs: List menus for restaurants which have different zip code.
payment.ejs: For users to make a payment
search.ejs: For users to search and view more information of restuarants.
signup.ejs: Sign up page
summary.ejs: For users to view the summary of the items in the cart
## the WOW factors we have worked on
Swarm is used as container orchestration tool for managing containerized infrastructure at scale

## Weekly progress
### Week1
* Meet up with the group members and decide the topic we are going to use for the project
* Talk to each other and specify each one's capabilities and experiences
* Brain storm of any technology that could be implemented in our project
* Research those technologies that more members are agreed to use for the project

### Week2
* Decide which website we are going to clone the business logic
* Agreed on burger online order for our team
* Allocate the functionality and the related database that each member should work on
* Decide the data structure that will be used for our database and unify the data we are passing among GO APIs
* Started coding for GO APIs

### Week3
* Formed the initial version of the design architecture
* Finished the first version of our GO APIs
* Discuss the changes that we need to make for integrating everyone's backend
* Finish testing the backend via postman
* Start working on the front end and the node.js that calls the GO APIs.

### Week4
* Dockerize the go api and try running on the aws instances where we run the database clusters
* Front-end in Express is done.
* Made further changes for integrating the GO APIs.
* Improved the UI.
* Integrated the front end.

### Week5
* Made modifications foe the database cluster for better architecture design
* Finished integrating GO API.
* Worked on the Wow factor.
* Wrapped up everything and finished the journal
* Finished the slides
* Decide the demo flow of our website
*
## Architecture diagram
<img src=https://github.com/nguyensjsu/fa18-281-slack-overflow/blob/master/Architecture.png>

## Challenges
* integrating everyone's work took much more time than we thought
* Figuring out the design of the architecture and the functionality of the whole project
* connecting the GO API, the node.js and the front end

## Further improvement that we can continue working on
* Kong gateway could be added for better safety of our system
