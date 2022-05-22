const offerButton = document.getElementById("offerButton");

document.loggedIn.then(value => {
    if(!value){offerButton.style.display = "none"}
})

function saveOffer(){
    fetch(`/private/api/offer`, {
        method: "POST",
        body: JSON.stringify(
            {
                "GoodID": parseInt(location.pathname.split("/").last()),
                "Value": parseFloat(document.getElementById("price").value),
                "Quantity": parseFloat(document.getElementById("quantity").value)
            }
        )
    })
        .then(_ => location.reload())
        .catch(reason => console.error(reason))
}