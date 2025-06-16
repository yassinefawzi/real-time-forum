document.addEventListener('DOMContentLoaded', function() {
    console.log('Posts script loaded');
    const postForm = document.getElementById('myCreateForm');
    if (postForm) {
        postForm.addEventListener('submit', handlePostCreation);
    }
});

async function handlePostCreation(event) {
    event.preventDefault();

    const form = document.getElementById('myCreateForm');
    const errorDiv = document.getElementById('createErrorMessage');
    const successDiv = document.getElementById('createSuccessMessage');
    const postsDiv = document.getElementById('feedPost');

    // Hide previous messages
    errorDiv.style.display = 'none';
    successDiv.style.display = 'none';

    // Convert form data to URL-encoded string instead of FormData
    const formData = new FormData(form);
    const urlEncodedData = new URLSearchParams(formData).toString();
    
    try {
        const response = await fetch('/api/createpost', {
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
                document.getElementById("createPost").style.display = "none";
                fetchPosts();
                postsDiv.style.display = "grid";
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

document.getElementById('createicon').addEventListener('click', function() {
    document.getElementById("createPost").style.display = "block";
    document.getElementById("feedPost").style.display = "none";
    document.getElementById("myCreateForm").reset();
    document.getElementById("createErrorMessage").style.display = "none";
    document.getElementById("createSuccessMessage").style.display = "none";
});

export async function fetchPosts() {
    console.log("Fetching posts...");
    try {
        const response = await fetch('/api/posts');
        if (!response.ok) throw new Error('Failed to fetch posts');
        const posts = await response.json();

        const feedPost = document.getElementById('feedPost');
        feedPost.style.display = 'grid';
        feedPost.innerHTML = '';

        posts.forEach(post => {
            const postDiv = document.createElement('div');
            postDiv.className = 'posts';
            postDiv.id = `${post.id}`;
            postDiv.innerHTML = `
                
                <h2 class="postTitle">${post.title}</h2>
                <p class="postCategory">Category: ${post.category}</p>
                <p class="postContent">${post.content}</p>
            `;

            feedPost.appendChild(postDiv);
        });
    } catch (err) {
        console.error('Error loading posts:', err);
    }
}

document.addEventListener("click", function(e) {
    const post = e.target.closest(".posts");
    if (post) {
        document.getElementById("feedPost").style.display = "none";
        singlePost(post);
    }
});

async function singlePost(post) {
    try {
        const response = await fetch(`/api/singlepost/${post.id}`);
        if (!response.ok) throw new Error('Failed to fetch single post');
        const postData = await response.json();

        const single = document.getElementById('singlePost');
        single.style.display = 'grid';
        single.innerHTML = `
            <h2 class="postTitle">${postData.title}</h2>
            <p class="postCategory">Category: ${postData.category}</p>
            <p class="postContent">${postData.content}</p>
        `;
        Array.from(document.getElementsByClassName('postContent')).forEach(el => {
        el.style.display = 'block'});

    } catch (error) {
        console.error('Error displaying single post:', error);
    }
}
    