var Priority = [], dbData = new Map()

function test() {
   window.location.search = "sort=score"
}

function indexLoad() {
   document.getElementById("uplImg").value = null
   getDBData(40);
}

async function getDBData(lim) {
   dbData = null
   let thePromise = new Promise(function (resolve) {
      const dbXML = new XMLHttpRequest();
      dbXML.onload = function () {
         if (this.status != 200) {
            document.getElementById("galleryWrap").style.display = none;
            resolve(null)
            return
         }
         console.log(this.responseText)
         let respon = JSON.parse(this.responseText)
         resolve(respon)
      }
      dbXML.open("GET", "/getDBData/"+window.location.search);
      dbXML.send();
   });
   dbData = await thePromise
   //console.log(dbData)
   UpdatePriority()
   loadImgs(lim)
}

function UpdatePriority() {
   const urlParams = new URLSearchParams(window.location.search);
   let sortBy = [], mapData = new Map()
   Priority = []
   for (id in dbData) {
      let upAt = 0
      switch(urlParams.get("sort")){
         case "score": upAt = parseInt(dbData[id].Score); break;
         default: upAt = dbData[id].UpdatedAt; break;
      }
      //console.log(upAt)
      sortBy.push(upAt)
      if (mapData[upAt] == undefined) {
         mapData[upAt] = []
      }
      mapData[upAt].push(id)
   }

   //sortBy.sort()
   sortBy.sort((a, b) => a-b)

   if(urlParams.get('asc')!="1"){ sortBy.reverse() }
   for (i in sortBy) {
      Priority[i] = mapData[sortBy[i]].shift()
   }
}

function printPriority() { console.log(Priority) }
function printDBData() { console.log(dbData) }

var ofstA = 0

function loadImgs(lim) {
   let ofst = ofstA
   for (let i = ofst; i < ofst + lim; i++) {
      if (Priority[i] == undefined) {
         document.getElementById("loadImgBtn").style.display = "none"
         //document.getElementById("loadImgBtn").ariaDisabled = true // doesn't work ???? (disable display none)
         document.getElementById("galleryEnd").innerText = "hmm, seems like you've reached the end of the gallery."
         console.log("all imgs loaded")
         break
      }
      let imgID = Priority[i]
      document.getElementById("gallery").innerHTML +=
         "<a " +
         "href='/img/" + imgID +  "' " +
         "class='imgGallery'" +
         "draggable='false'" +
         ">" +

         "<div " +
         "class='imgScore ff-mn'" +
         ">" +
         "<label>" + dbData[imgID].Score +"</label>"+
         "</div>" +

         "<img " +
         "src='/gallery/" + imgID + dbData[imgID].FileExt + "' " +
         "alt='" + imgID + "' " +
         "draggable='false'" +
         "/>" +

         "<div " +
         "class='imgInfo'" +
         "id='imgInfo" + imgID + "'" +
         ">" +
         dbData[imgID].Name + "<br/>" + dbData[imgID].Owner +
         "</div>" +

         "</a> ";
      ofstA++
   }
}

function imgUpl() {
   document.getElementById("uplImg").click()
   document.getElementById("uplBtn").focus()
}

function isOverflown(elementID) {
   elm = document.getElementById(elementID)
   //console.log(elm.scrollHeight, elm.clientHeight)
   return elm.scrollHeight > elm.clientHeight || elm.scrollWidth > elm.clientWidth;
}

function RandVal(min, max) { // inclusive min & max
   min = Math.floor(min);
   max = Math.floor(max);
   return Math.floor(Math.random() * (max - min + 1) + min);
}

function invert() {
   let light = document.getElementsByClassName("light")
   let dark = document.getElementsByClassName("dark")
   for(let i=0; i<light.length; i++){
      light[i].classList.add("dark")
      light[i].classList.remove("light")
   }
   for(let i=0; i<dark.length; i++){
      dark[i].classList.add("light")
      dark[i].classList.remove("dark")
   }
}



