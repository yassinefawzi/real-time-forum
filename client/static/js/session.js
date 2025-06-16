console.log("Session script loaded");

import { fetchPosts } from './posts.js';

document.addEventListener('DOMContentLoaded', checkSession);


async function checkSession() {
    try {
        const response = await fetch('/api/checksession');
        const result = await response.json();

        if (result.loggedIn) {
            document.getElementById("logInSection").style.display = "none";
            document.getElementById("feedPost").style.display = "grid";
            document.getElementById("logout").style.display = "block";
            document.getElementById("createicon").style.display = "block";
            console.log("User is logged in:", result.username);
            fetchPosts();
        } else {
            document.getElementById("logInSection").style.display = "block";
            document.getElementById("feedPost").style.display = "none";
        }
    } catch (error) {
        console.error('Failed to check session:', error);
    }
}