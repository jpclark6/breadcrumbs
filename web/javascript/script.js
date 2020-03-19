function logLocation() {
    var x = document.getElementById("changethis");
    x.innerHTML = "Getting location..."
    document.getElementById("litter").classList.toggle('hidden');
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(showPosition);
    } else {
        x.innerHTML = "Geolocation is not supported by this browser.";
    }
}

function showPosition(position) {
    var x = document.getElementById("changethis");
    x.innerHTML = "Latitude: " + position.coords.latitude +
        "<br>Longitude: " + position.coords.longitude;
    lat = document.getElementById("lat")
    lat.value = position.coords.latitude
    long = document.getElementById("long")
    long.value = position.coords.longitude
}

function getLocation() {
    x.innerHTML = "Finding nearby messages..."
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(findMessages);
    } else {
        x.innerHTML = "Geolocation is not supported by this browser.";
    }
}

function findMessages(position) {
    var lat = position.coords.latitude
    var long = position.coords.longitude
    var url = "/getbreadcrumbs"
    var params = {
        lat: lat,
        long: long
    }
    axios.get(url, {params: params})
    .then(data=>{
        buildList(data)
    })
    .catch(err=>console.log(err))
}

function buildList(data) {
    var updatedList = []
    data.data.messages.forEach(message => {
        date = new Date(message.CreatedAt)
        var displayData = {
            Distance: message.Distance.toString(),
            Breadcrumb: message.Text,
            CreatedAt: date.toString()
        }
        updatedList.push(displayData)
    });

    var table = document.querySelector("table")
    var newdata = Object.keys(updatedList[0])
    generateTable(table, updatedList)
    generateTableHead(table, newdata)
}

function generateTableHead(table, data) {
    let thead = table.createTHead();
    let row = thead.insertRow();
    for (let key of data) {
        let th = document.createElement("th");
        let text = document.createTextNode(key);
        th.appendChild(text);
        row.appendChild(th);
    }
}

function generateTable(table, data) {
    for (let element of data) {
        let row = table.insertRow();
        for (key in element) {
            let cell = row.insertCell();
            let text = document.createTextNode(element[key]);
            cell.appendChild(text);
        }
    }
}
