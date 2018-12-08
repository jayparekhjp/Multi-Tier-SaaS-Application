# fa18-281-slack-overflow
## Architecture diagram
<img src=https://github.com/nguyensjsu/fa18-281-slack-overflow/blob/master/Architecture.png>

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
Cassandra
GoLang
Node.js
Cassandra
MongoDB
session and cookie

## Go api Functionality
  Users:
  Create the registration and login functions for users, and allow the userId to be passed through all the requests sent by a certain user via session. The user information is stored in all the database.
  Menu:
  Restaurants in different locations have varied menus. For online orders, users can view menu via typing Zip code of the restaurants.  
  Cart:
  While viewing items in a menu, users can add items with a specific count and add to the cart for this restaurant. The cart information is stored in the counter database.  
  Payment:
  When users checking out items in shopping carts, they can make payment by typing user's payment information and then make a payment. After clicking the button 'make payment', the order history will be passed to the next goapi.
  Orders:
  After a user make a payment, the order history is kept in the buigerOrder database. At the same page, users can view his own order history by clicking on "List all your order history".

## node.js
The app.js is designed to send requests to goapis and render or redirect pages to separate ejs files.  

## ejs pages
admin.ejs:
List all the user information and only available for admin user
cart.ejs:
Show the cart for a certain user
header.ejs:
Gives an overview of how the website works and shows the users the navigation bar to home page, restaurants page and description page
home.ejs:
Home page of our website
insertAfterPayment.ejs:
Shows the most recent payment of the user.
list.ejs:
Shows all the payment history of the user.
loin.ejs:
Login page
menu.ejs:
List menus for restaurants which have different zip code.
payment.ejs:
For users to make a payment
search.ejs:
For users to search and view more information of restuarants.
signup.ejs:
Sign up page
summary.ejs:
For users to view the summary of the items in the cart

## WOW factors we have worked on
  Swarm is used as container orchestration tool for managing containerized infrastructure at scale
## CAP Theorem
CAP Theorem is the concept that a distributed database system can only satisfy 2 of the 3 properties: Consistency, Availability and Partition Tolerance. CAP theorem is used to make decision on which databases that will be used in a project based on the very specific user case.
### Partition Tolerance
A system is considered to be partition tolerant if it continues running unless failure of the entire network. Data is stored horizontally among nodes that baring with partial outage. Counting on reliable and powerful machines are way more expensive, so modern distributed system has to trade off between consistency and availability.
### High Consistency
Every request comes to the server gets a response of success or failure, which means that a user can either read or write to the database or not. For high consistency, servers reject some requests on some nodes.
### High Availability
When database are configured or designed to be with high availability, all the requests sent to all the nodes can not be rejected. But to achieve the metric, the data in every node may not match to each other because they sacrifice the Conflict serializability.

## MongoDB
MongoDB has tunable consistency for read and write operations. MongoDB uses write concerns to specifies the level of acknowledgement required for writes, and uses read preferences for sending requests to members of a replica set. Developers can also tune the read concern levels for consistency and isolation properties of the data read from replica sets and replica set shards.
For example, if the primary node which allows write operations is down for some reason, a new primary node will be elected among the other members in the cluster. Before that, write operations are not allowed because of loss of the primary node. Once the partition is removed, the previous primary node will be synchronized with the updated data, which is called eventual consistency.

## Cassandra
Cassandra database is peer-peer model meaning that each node is coordinating both write and read operations. If a node that received a write goes down, other write will be redirected to other node that has the same shard key. But this can further configured via replication factor(how many copies of data) and consistency level (read and write).


## Weekly progress
###Week1
* Meet up with the group members and decide the topic we are going to use for the project
* Talk to each other and specify each one's capabilities and experiences
* Brain storm of any technology that could be implemented in our project
* Research those technologies that more members are agreed to use for the project

###Week2
* Decide which website we are going to clone the business logic
* Agreed on burger online order for our team
* Allocate the functionality and the related database that each member should work on
* Decide the data structure that will be used for our database and unify the data we are passing among GO APIs
* Started coding for GO APIs

###Week3
* Formed the initial version of the design architecture
* Finished the first version of our GO APIs
* Discuss the changes that we need to make for integrating everyone's backend
* Finish testing the backend via postman
* Start working on the front end and the node.js that calls the GO APIs.

###Week4
* Dockerize the go api and try running on the AWS instances where we run the database clusters
* Front-end in Express is done.
* Made further changes for integrating the GO APIs.
* Improved the UI.
* Integrated the front end.

###Week5
* Made modifications foe the database cluster for better architecture design
* Finished integrating GO API.
* Worked on the Wow factor.
* Wrapped up everything and finished the journal
* Finished the slides
* Decide the demo flow of our website


## Challenges
* integrating everyone's work took much more time than we thought
* Figuring out the design of the architecture and the functionality of the whole project
* connecting the GO API, the node.js and the front end

## Further improvement that we can continue working on
* Kong gateway could be added for better safety of our system
