if (!Array.prototype.last){
    Array.prototype.last = function(){
        return this[this.length - 1];
    };
}

fetch(`/api/goods/${location.pathname.split("/").last()}/course`)
    .then(response => response.json())
    .then(data => {
        draw(data)
    })
    .catch(err => console.log(err));

function draw(dataServer){
    const data = [
        {
            x: dataServer.map(ele => ele["CreatedAt"]),
            y: dataServer.map(ele => ele["Value"]),
            type: 'scatter'
        }
    ];

    Plotly.newPlot('coursePlot', data);
}