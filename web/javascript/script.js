function logLocation() {
    addLoading()
    const container = document.getElementById('bc')
    container.innerHTML = ""
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(showPosition);
    } else {
        loading.innerHTML = "Geolocation is not supported by this browser.";
    }
}

async function addLoading() {
    var node = document.createElement("h2");   
    node.setAttribute("id", "loadingtext");              // Create a <li> node
    var textnode = document.createTextNode("");          // Create a text node
    node.appendChild(textnode);                          // Append the text to <li>
    var loading = document.getElementById("loading")
    loading.appendChild(node);
    loading.classList.remove("hidden")
    while (true) {
        await sleep(500);
        if (document.getElementById("loadingtext")) {
            document.getElementById("loadingtext").innerHTML = "Finding location."
        }
        await sleep(500);
        if (document.getElementById("loadingtext")) {
            document.getElementById("loadingtext").innerHTML = "Finding location.."
        }
        await sleep(500);
        if (document.getElementById("loadingtext")) {
            document.getElementById("loadingtext").innerHTML = "Finding location..."
        }
    }
}

function removeLoading() {                            // Append the text to <li>
    var loading = document.getElementById("loading")
    loading.innerHTML = ""
    loading.classList.add("hidden")
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function showPosition(position) {
    removeLoading()
    document.getElementById("long-descriptions").classList.add("hidden")
    document.getElementById("litter").classList.remove('hidden');
    lat = document.getElementById("lat")
    lat.value = position.coords.latitude.toFixed(9)
    long = document.getElementById("long")
    long.value = position.coords.longitude.toFixed(9)
}

function getLocation() {
    addLoading()
    document.getElementById("litter").classList.add('hidden');
    document.getElementById("long-descriptions").classList.add("hidden")
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(findMessages);
    } else {
        removeLoading()
        loading.innerHTML = "Geolocation is not supported by this browser.";
    }
}

function findMessages(position) {
    removeLoading()
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
            CreatedAt: date
        }
        updatedList.push(displayData)
    });

    var table = document.getElementById("bc")
    table.classList.remove("hidden")
    generateTable(table, updatedList)
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
    const apiResult = data

    const container = document.getElementById('bc')
    container.innerHTML = ""

    if (apiResult.length > 0) {
        apiResult.forEach((result, idx) => {
            const card = document.createElement('div');
            card.classList = 'card-body';
    
            const content = `
                <div class="card">
                    <div>
                        <div class="card-body">
                            <h3>${result.Breadcrumb}</h3>
                            <p class="grey-text">${result.Distance} miles away</p>
                            <p class="grey-text">Dropped at ${result.CreatedAt.toLocaleTimeString('en-US')} on ${result.CreatedAt.toDateString()}</p>
                        </div>
                    </div>
                </div>
            `;
    
            container.innerHTML += content;
        })
    } else {
        const card = document.createElement('div');
        card.classList = 'card-body';

        const content = `
                <div class="card">
                    <div>
                        <div class="card-body">
                            <h3>Sorry, no breadcrumbs around here.</h3>
                        </div>
                    </div>
                </div>
            `;

        // Append newyly created card element to the container
        container.innerHTML = content;
    }
}
