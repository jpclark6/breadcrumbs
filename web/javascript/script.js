var x = document.getElementById("changethis");
function logLocation() {
    x.innerHTML = "Getting location..."
    document.getElementById("litter").classList.toggle('hidden');
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(showPosition);
    } else {
        x.innerHTML = "Geolocation is not supported by this browser.";
    }
}

function showPosition(position) {
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
    var url = "/findbreadcrumb"
    var params = {
        lat: lat,
        long: long
    }
    axios.get(url, {params: params})
    .then(data=>{
        console.log(data)
        text = data.data.messages
        x.innerHTML = text
    })
    .catch(err=>console.log(err))
}

