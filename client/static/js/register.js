document.getElementById("signUpLink").addEventListener("click", function(e) {
    e.preventDefault();
    document.getElementById("logInSection").style.display = "none";
    document.getElementById("signUpSection").style.display = "block";
    console.log("Switched to Sign Up section");
})