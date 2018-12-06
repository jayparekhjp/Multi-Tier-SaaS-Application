var app = document.getElementById('root');

// const logo = document.createElement('img');
// logo.src = 'images/logo1.png';

const container = document.createElement('div');
container.setAttribute('class', 'container');

//app.appendChild(logo);
app.appendChild(container);

var request = new XMLHttpRequest();
request.open('GET', 'http://localhost:3002/cart');
request.onload = function () {

  // Begin accessing JSON data here
  var data = JSON.parse(this.response);
  if (request.status >= 200 && request.status < 400) {
      console.log("hi")
    data.forEach(cart => {
      const card = document.createElement('div');
      card.setAttribute('class', 'card');

      const h1 = document.createElement('h1');
      h1.textContent = cart.items;

      const h2 = document.createElement('h1');
      h1.textContent = cart.price;

    //   const p = document.createElement('p');
    //   movie.description = movie.description.substring(0, 300);
    //   p.textContent = `${movie.description}...`;

      container.appendChild(card);
      card.appendChild(h1);
      card.appendChild(h2);
    });
  } else {
    const errorMessage = document.createElement('marquee');
    errorMessage.textContent = `Gah, it's not working!`;
    app.appendChild(errorMessage);
  }
}

request.send();