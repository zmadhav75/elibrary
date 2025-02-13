document.getElementById('subscribeForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const email = document.getElementById('email').value;
    
    try {
        const response = await fetch('/api/subscribe', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({ email })
        });
        
        if (response.ok) {
            alert('Subscription successful!');
        } else {
            alert('Subscription failed');
        }
    } catch (error) {
        console.error('Error:', error);
    }
});