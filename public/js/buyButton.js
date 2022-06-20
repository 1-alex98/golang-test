if (!Array.prototype.last){
    Array.prototype.last = function(){
        return this[this.length - 1];
    };
}
let lastId, lastPrice;

const quantityInput = document.getElementById("buyQuantity");
const priceText= document.getElementById("priceText");

function buyDialog(id, price , quantity){
    lastId = id
    lastPrice = price
    quantityInput.max = quantity
    updatePrice()
}

function updatePrice(){
    const value = quantityInput.value;
    priceText.innerText = (parseFloat(value) * lastPrice).toString();
}

function buyOffer(){
    const value = parseFloat(quantityInput.value);
    sendBuy(value);
}

function sendBuy(quantity){
    fetch(`/private/api/offer/${lastId}/order`,{
        method: "POST",
            body: JSON.stringify(
            {
                "Quantity": quantity
            }
        )
    })
        .then(_ => location.reload())
        .catch(err => console.log(err));

}