// watchtower


// This is the target URL
// /reader.html#/api/publications/w_202101/
// Current date: 25.03.2021
//    February 2021      
// Su Mo Tu We Th Fr Sa  
// 1  2  3  4  5  6  
// 7  8  9 10 11 12 13  
// 14 15 16 17 18 19 20  
// 21 22 23 24 25 26 27  
// 28                   
// 
// March 2021       
// Su Mo Tu We Th Fr Sa  
//    [1] 2  3  4  5  6  
//  7  8  9 10 11 12 13  
// 14 15 16 17 18 19 20  
// 21 22 23 24 25 26 27  
// 28 29 30 31           
//
// April 2021       
// Su Mo Tu We Th Fr Sa  
//              1  2  3  
// [4] 5  6  7  8  9 10  
// 11 12 13 14 15 16 17  
// 18 19 20 21 22 23 24  
// 25 26 27 28 29 30     
// 
// Articles from the w_202101
// watchtower is being studied
// - from 01.03
// - to 04.04
//
// - - - - - - - - - - - -
//
// with mwb things are easier
// 21.01 is used in 21.01 and 21.02
function setMeeting() {
    // Watchtower
    date = getThisSunday()
    date.setMonth(date.getMonth() - 2)
    month = ("00" + (date.getMonth() + 1)).slice(-2)
    year = date.getFullYear()
    wt = "w_"+year+month
    fetch("/api/publications_json/"+wt+"/toc.xhtml")
    .then(response => response.json())
    .then((response) => { 
        toc = response.html.body.section.nav[0].ol.li[1].a['-href']
        wtlink = document.getElementById("wt_link")
        wtlink.href = "/reader.html#/api/publications/"+wt+"/"+toc
    })
    // Now I need:
    // mwb_202105
    mwbdate = getThisSunday()
    if (((mwbdate.getMonth()) % 2) === 0) {
        mwbdate.setMonth(mwbdate.getMonth() - 1)
    } else {
        mwbdate.setMonth(mwbdate.getMonth())
    }
    mwbmonth = ("00" + (mwbdate.getMonth())).slice(-2)
    mwbyear = mwbdate.getFullYear()
    mwb = "mwb_"+mwbyear+mwbmonth
    console.log(mwb)
    mwblink = document.getElementById("mwb_link")
    mwblink.href = "/reader.html#/api/publications/"+mwb+"/"
}

function getThisSunday() {
    var t = new Date();
    t.setDate(t.getDate() - t.getDay());
    return t;
}
setMeeting()