let navLoginLink = document.getElementById("navLoginLink");
let accountLink = document.getElementById("accountLink");
document.loggedIn= new Promise((resolve, _) => {
    fetch('/private/me')
        .then(response => response.json())
        .then(data => {
            navLoginLink.innerText= "Logout being: "+data["email"]
            navLoginLink.onclick = logout
            resolve(true)
        })
        .catch(_ => {
            resolve(false)
            accountLink.style.display = "none"
        });

});

function logout(evt) {
    evt.preventDefault()
    fetch('/logout', {method: 'post'}).then(_ => {})
    location.reload()
}

