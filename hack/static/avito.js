const addString = document.getElementById("addString")
const saveChangesDiv = document.getElementById("saveChangesDiv")
const body = document.getElementById("body")
const saveChanges = document.getElementById("saveChanges")
let c = 2

addString.addEventListener('click', function(){
    let newInputStrings = document.createElement('div')
    newInputStrings.id = "inputStrings" + String(c)
    newInputStrings.className = "inputStrings"
    newInputStrings.innerHTML = '<input id = "mcID' + String(c) + '" name = "mcID' + String(c) + '" placeholder= "CategoryID" class = "inputs" style = "margin-left:0.5vw">\n' +
                        '        <input id = "lcID' + String(c) + '" name = "lcID' + String(c) + '" placeholder= "LocationID" class = "inputs">\n' +
                        '        <input id = "prID' + String(c) + '" name = "prID' + String(c) + '" placeholder= "Price" class = "inputs">'
    body.insertBefore(newInputStrings,saveChangesDiv)
    window.scrollBy(0, 100);
    c += 1
})

saveChanges.onclick = function(){
    let arr = []
    let data
    for (let i = 1;i < c + 1;i ++){
        data = {mcID:document.getElementById(`mcID${String(i)}`).value,
                lcID:document.getElementById(`lcID${String(i)}`).value,
                prID:document.getElementById(`prID${String(i)}`).value}
        arr.push(data)
    }
    console.log(arr)
}