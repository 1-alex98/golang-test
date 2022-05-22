let navLoginLink = document.getElementById("navLoginLink");

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
    })
    .catch(err => console.log(err));