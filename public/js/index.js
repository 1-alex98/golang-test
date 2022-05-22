let navLoginLink = document.getElementById("navLoginLink");
let accountLink = document.getElementById("accountLink");

function logout(evt) {
    evt.preventDefault()
    fetch('/logout', {method: 'post'}).then(_ => {})
    location.reload()
}

fetch('/private/me')
    .then(response => response.json())
    .then(data => {
        navLoginLink.innerText= "Logout being: "+data["email"]
        navLoginLink.onclick = logout
        document.loggedIn= true
    })
    .catch(_ => {
        document.loggedIn= false
        accountLink.style.display = "none"
    });