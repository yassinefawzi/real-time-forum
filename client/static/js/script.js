let username
let password


document.getElementById('myLogInForm').addEventListener('submit', function(e) {
    e.preventDefault(); // prevent page reload
  
    username = document.getElementById('myUserName').value;
    password = document.getElementById('myPassword').value;
});
// Log the values to the console 

console.log(`Title: ${username}, Content: ${password}`);