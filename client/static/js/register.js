console.log("Register script loaded");

import { fetchPosts } from './posts.js';

document.getElementById("signUpLink").addEventListener("click", function(e) {
    const form = document.getElementById('myLogInForm');
    const errorDiv = document.getElementById('errorMessage');
    const successDiv = document.getElementById('successMessage');
    errorDiv.style.display = 'none';
    successDiv.style.display = 'none';
    e.preventDefault();
    document.getElementById("logInSection").style.display = "none";
    document.getElementById("signUpSection").style.display = "block";
    console.log("Switched to Sign Up section");
    
})

document.addEventListener('DOMContentLoaded', function() {
    const signUpForm = document.getElementById('mySignUpForm');
    if (signUpForm) {
        signUpForm.addEventListener('submit', handleSignUp);
    }
});



async function handleSignUp(event) {
    event.preventDefault();
    const form = document.getElementById('mySignUpForm');
    const errorDiv = document.getElementById('errorMessage');
    const successDiv = document.getElementById('successMessage');
    
    // Hide previous messages
    errorDiv.style.display = 'none';
    successDiv.style.display = 'none';
    
    // Convert form data to URL-encoded string instead of FormData
    const formData = new FormData(form);
    const urlEncodedData = new URLSearchParams(formData).toString();
    
    try {
        const response = await fetch('/api/signup', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: urlEncodedData
        });
        
        const result = await response.json();
        
        if (result.success) {
            successDiv.textContent = result.message;
            successDiv.style.display = 'block';
            
            setTimeout(() => {
                document.getElementById("signUpSection").style.display = "none";
                document.getElementById("logInSection").style.display = "block";
                form.reset();
            }, 2000);
            
        } else {
            errorDiv.textContent = result.error;
            errorDiv.style.display = 'block';
        }
        
    } catch (error) {
        errorDiv.textContent = 'Network error. Please try again.';
        errorDiv.style.display = 'block';
        console.error('Error:', error);
    }
    
    return false;
}

// log in form section
document.addEventListener('DOMContentLoaded', function() { 
    const logInForm = document.getElementById('myLogInForm');
    if (logInForm) {
        logInForm.addEventListener('submit', handleLogIn);
    }
    console.log("Log In form event listener added");
});

async function handleLogIn(event) {
    event.preventDefault();

    const form = document.getElementById('myLogInForm');
    const errorDiv = document.getElementById('logerrorMessage');
    const successDiv = document.getElementById('logsuccessMessage');
    
    // Hide previous messages
    errorDiv.style.display = 'none';
    successDiv.style.display = 'none';
    // Convert form data to URL-encoded string instead of FormData
    const formData = new FormData(form);
    const urlEncodedData = new URLSearchParams(formData).toString();
    
    try {
        const response = await fetch('/api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: urlEncodedData
        });

        const result = await response.json();

        if (result.success) {
            successDiv.textContent = result.message;
            successDiv.style.display = 'block';                
            setTimeout(() => {
                document.getElementById("logInSection").style.display = "none";
                form.reset();
                fetchPosts();
            }, 2000);
        } else {
            errorDiv.textContent = result.error;
            errorDiv.style.display = 'block';
        }
    } catch (error) {
        errorDiv.textContent = 'Network error. Please try again.';
        errorDiv.style.display = 'block';
        console.error('Error:', error);
    }
    return false;
}