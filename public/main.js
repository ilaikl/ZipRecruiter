var allIcons = document.querySelectorAll(".fa-star");
allIcons.forEach(i => {
    if (localStorage.getItem(i.dataset.id) == 'true')
        i.classList.toggle("fas")
});

// Get the modal
var modal = document.getElementById("myModal");

// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];

// When the user clicks the button, open the modal 
function openModal(jobTitleElement) {
    document.getElementById("modal-description").innerHTML = jobTitleElement.childNodes[1].innerHTML
    modal.style.display = "block";
}
let isOpenSearch=false;
function openSearch(d) {
    isOpenSearch=true;
    document.getElementById("matching-jobs").style.top = "-44px";
    document.getElementById("matching-jobs").style.position= "relative"; 
    document.getElementById("matching-jobs").style.zIndex =  "-1";
    document.getElementById("keyword").style.visibility = "visible"
    document.getElementById("keyword").style.display = "unset";
    document.getElementById("location").style.visibility = "visible"
    document.getElementById("location").style.display = "unset";
    document.getElementById("search-jobs-button").style.visibility = "visible"
    document.getElementById("search-jobs-button").style.display = "unset";
    d.style.visibility = "hidden"
    d.style.display = "none";
}
// When the user clicks on <span> (x), close the modal
function closeModal() {
    modal.style.display = "none";
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function (event) {
    if (event.target == modal) {
        modal.style.display = "none";
        return
    }

    if (isOpenSearch && event.target != document.getElementById("search-jobs-form") && document.getElementById("open-search").style.visibility == "hidden" && (event.target != document.getElementById("open-search")) && (event.target != document.getElementById("keyword")) && (event.target != document.getElementById("location"))) {
        document.getElementById("matching-jobs").attributeStyleMap.clear()
        document.getElementById("keyword").attributeStyleMap.clear()
        document.getElementById("location").attributeStyleMap.clear()
        document.getElementById("search-jobs-button").attributeStyleMap.clear()
        document.getElementById("open-search").attributeStyleMap.clear()
        isOpenSearch=false;
    }

}

function saveJob(i) {
    i.classList.toggle("fas")
    let isFaved = localStorage.getItem(i.dataset.id)
    if (isFaved == undefined || isFaved == null) {
        localStorage.setItem(i.dataset.id, false)
    }
    isFaved = isFaved != 'true'
    localStorage.setItem(i.dataset.id, isFaved)
}