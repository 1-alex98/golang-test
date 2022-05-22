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

function offers(dataServer){
    let offerTable = document.getElementById("offerTable");
    for(let offer of dataServer){
        offerTable.innerHTML+=`
<tr>
    <td>${escapeHtml(offer["Quantity"])}</td>
    <td>${escapeHtml(offer["Value"])}<i class="bi bi-currency-euro"></i></td>
    <td>
        <button type="button" class="btn btn-primary">
            Buy
        </button>
    </td>
</tr>
            
        `
    }
}