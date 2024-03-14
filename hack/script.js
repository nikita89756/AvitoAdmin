console.log("JS Loaded")
const selectMatrix = document.getElementById("selectMatrix")
const url = "http://localhost:8080/test"

selectMatrix.onfocus = function (){
    fetch(url,{method:"GET"})
        .then(response => response.json())
        .then(data => console.log(data));
}

