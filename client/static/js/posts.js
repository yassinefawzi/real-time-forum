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
                //document.getElementById("posts").style.display = "block";
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