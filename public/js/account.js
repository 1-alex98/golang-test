let selectedId;
function saveAmountPopUp(id){
    selectedId = id
}

function saveAmount(){
    fetch(selectedId === -1 ? `/private/api/credit`: `/private/api/account/${selectedId}`, {
        method: "PUT",
        body: JSON.stringify(
            {
                "Value": parseFloat(document.getElementById("amountInput").value)
            }
        )
    })
        .then(_ => location.reload())
        .catch(reason => console.error(reason))
}