if (!Array.prototype.last){
    Array.prototype.last = function(){
        return this[this.length - 1];
    };
}

fetch(`/api/goods/${location.pathname.split("/").last()}/offers`)
    .then(response => response.json())
    .then(data => {
        offers(data)
    })
    .catch(err => console.log(err));

function escapeHtml(html){
    const text = document.createTextNode(html);
    const p = document.createElement('p');
    p.appendChild(text);
    return p.innerHTML;
}

async function offers(dataServer){
    let offerTable = document.getElementById("offerTable");

    const loggedIn = await document.loggedIn

    for(let offer of dataServer){
        offerTable.innerHTML+=`
<tr>
    <td>${escapeHtml(offer["Quantity"])}</td>
    <td>${escapeHtml(offer["Value"])}<i class="bi bi-currency-euro"></i></td>
    <td>
        <button type="button" class="btn btn-primary${!loggedIn?' disabled':''}" onclick="alert('not implemented')">
            <i class="bi bi-cart-plus"></i> Buy
        </button>
        <button type="button" class="btn btn-primary disabled" onclick="alert('not implemented')">
            <i class="bi bi-trash3"></i> Withdraw
        </button>
    </td>
</tr>
            
        `
    }
}