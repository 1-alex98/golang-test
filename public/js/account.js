let selectedId;
function saveAmountPopUp(id, evt){
    selectedId = id
}

function saveAmount(){
    fetch("/private/api/account/selectedId", {
        method: "POST",
        body: JSON.stringify(
            {
                "Value": document.getElementById("amountInput").value
            }
        )
    })
        .then(_ => location.reload())
        .catch(reason => console.error(reason))
}