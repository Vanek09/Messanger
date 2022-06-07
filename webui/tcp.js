const Net = require('net');
require('dotenv').config();
// The port number and hostname of the server.
const port = process.env.BROKER_API_PORT;
const host = process.env.BROKER_API_HOST;
const client = new Net.Socket();



function communicate(request, res, user, users, destination) {
  client.connect({ port: port, host: host }, function() {
    // If there is no error, the server has accepted the request and created a new 
    // socket dedicated to us.
    console.log('TCP connection established with the server.');
    console.log(request)
    client.write(JSON.stringify(request));
    // The client can now send data to the server by writing to its socket.
    
  });
  let response = null
  

  // The client can also receive data from the server by reading from its socket.
  client.once('data', function(chunk) {
      console.log(`Data received from the server: ${chunk.toString()}.`);
      response = JSON.parse(chunk.toString())
      res.render('todo.ejs', {users: users, messages: response, user: user, destination: destination});
      // Request an end to the connection after the data has been received.
      client.end();
  });

  client.on('end', function() {
      
  });
  
}

// client.connect({ port: port, host: host }, function() {
//   // If there is no error, the server has accepted the request and created a new 
//   // socket dedicated to us.
//   console.log('TCP connection established with the server.');

//   // The client can now send data to the server by writing to its socket.
  
// });

// function render_res(request, res, user, users, destination) {
//   let response = null
  
//   console.log(JSON.stringify(request))
//   client.write(JSON.stringify(request));
//   // The client can also receive data from the server by reading from its socket.
//   client.once('data', function(chunk) {
//       console.log(`Data received from the server: ${chunk.toString()}.`);
//       messages = JSON.parse(chunk.toString())
//       console.log(messages)
//       res.render('todo.ejs', {users: users, messages: messages, user: user, destination: destination})
//       // Request an end to the connection after the data has been received.
//       // client.end();
//   });

//   client.on('end', function() {
//       console.log('Requested an end to the TCP connection');
//       return response
//   });
  
// }

module.exports = {

    recieveMessages: function (res, user, users, destination) {
      request = {
        command: "getMessages",
        from: user,
        to: destination
      }
      communicate(request, res, user, users, destination)
      // render_res(request, res, user, users, destination)
    }
};