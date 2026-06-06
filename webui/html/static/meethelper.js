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
    // Show loading state immediately while fetches are in flight
    document.getElementById("mwb_link").textContent = "Life and Ministry — Loading..."
    document.getElementById("wt_link").textContent = "Watchtower — Loading..."

    // Watchtower
    // Save the original Sunday before rolling back 2 months for the publication code
    var thisSunday = getThisSunday()
    var date = new Date(thisSunday)
    date.setMonth(date.getMonth() - 2)
    var month = ("00" + (date.getMonth() + 1)).slice(-2)
    var year = date.getFullYear()
    var wt = "w_"+year+month
    fetch("/api/publications_json/"+wt+"/toc.xhtml")
    .then(response => response.json())
    .then((response) => {
        var items = response.html.body.section.nav[0].ol.li
        var result = getCurrentWTArticle(items, thisSunday)
        if (!result) throw new Error("Watchtower TOC has unexpected structure")
        var wtlink = document.getElementById("wt_link")
        wtlink.href = "/reader.html#/api/publications/"+wt+"/"+result.href
        wtlink.textContent = "Watchtower — " + result.title
    })
    .catch(() => {
        var wtlink = document.getElementById("wt_link")
        wtlink.textContent = "Watchtower — not available offline (open it once with internet)"
        wtlink.href = "/publications.html"
    })

    // MWB
    // MWB is bimonthly: mwb_YYYY01 covers Jan+Feb, mwb_YYYY03 covers Mar+Apr, etc.
    // getMonth() is 0-indexed: odd 0-indexed = 2nd month of pair → go back 1
    var mwbdate = getThisSunday()
    if ((mwbdate.getMonth() % 2) === 1) {
        mwbdate.setMonth(mwbdate.getMonth() - 1)
    }
    var mwbmonth = ("00" + (mwbdate.getMonth() + 1)).slice(-2)
    var mwbyear = mwbdate.getFullYear()
    var mwb = "mwb_"+mwbyear+mwbmonth
    fetch("/api/publications_json/"+mwb+"/toc.xhtml")
    .then(response => response.json())
    .then((response) => {
        var items = response.html.body.section.nav[0].ol.li
        var result = getCurrentMWBWeek(items)
        if (!result) throw new Error("MWB TOC has unexpected structure")
        var mwblink = document.getElementById("mwb_link")
        mwblink.href = "/reader.html#/api/publications/"+mwb+"/"+result.href
        mwblink.textContent = "Life and Ministry — " + result.title
    })
    .catch(() => {
        var mwblink = document.getElementById("mwb_link")
        mwblink.textContent = "Life and Ministry — not available offline (open it once with internet)"
        mwblink.href = "/publications.html"
    })
}

// Returns the href of the current week's entry in the MWB table of contents.
// items[0] = publication cover, items[1..N-1] = weekly entries, items[N] = page nav.
// Parses the start date from items[1]'s title (e.g. "May 4-10") and computes
// the week offset from today so the link always opens to the current week.
function getCurrentMWBWeek(items) {
    if (!items || items.length < 2) return null;
    var fallback = {href: items[1].a['-href'], title: items[1].a['#content']};
    var firstTitle = items[1].a['#content'] || '';
    var match = firstTitle.match(/([A-Za-z]+)\s+(\d+)/);
    if (!match) return fallback;
    var monthNames = {
        January:0, February:1, March:2,   April:3,
        May:4,     June:5,     July:6,    August:7,
        September:8, October:9, November:10, December:11
    };
    var monthNum = monthNames[match[1]];
    if (monthNum === undefined) return fallback;
    var today = new Date();
    var firstWeekStart = new Date(today.getFullYear(), monthNum, parseInt(match[2]));
    if (firstWeekStart > today) firstWeekStart.setFullYear(today.getFullYear() - 1);
    var weekOffset = Math.floor((today - firstWeekStart) / 604800000); // ms per week
    var maxWeek = items.length - 2; // exclude cover (0) and page nav (last)
    var weekIndex = Math.min(Math.max(1, 1 + weekOffset), maxWeek);
    return {href: items[weekIndex].a['-href'], title: items[weekIndex].a['#content']};
}

// Returns the href of the current week's study article in the Watchtower TOC.
// items[0] = cover, items[1] = Table of Contents, items[2..N-2] = study articles, items[N-1] = page nav.
// Finds the first Sunday of the study month (= current month, before the 2-month rollback),
// then computes the week offset to determine the correct article.
function getCurrentWTArticle(items, thisSunday) {
    if (!items || items.length < 3) return null;
    var fallback = {href: items[2].a['-href'], title: items[2].a['#content']};
    var firstDay = new Date(thisSunday.getFullYear(), thisSunday.getMonth(), 1);
    var dow = firstDay.getDay(); // 0=Sun, 1=Mon, ..., 6=Sat
    var firstSunday = new Date(firstDay);
    firstSunday.setDate(1 + (dow === 0 ? 0 : 7 - dow));
    var weekOffset = Math.floor((thisSunday - firstSunday) / 604800000); // ms per week
    var maxArticle = items.length - 3; // articles are items[2..length-2], skip cover+toc+pagenav
    var articleIndex = Math.min(Math.max(0, weekOffset), maxArticle);
    return {href: items[2 + articleIndex].a['-href'], title: items[2 + articleIndex].a['#content']};
}

function getThisSunday() {
    var t = new Date();
    t.setDate(t.getDate() - t.getDay());
    return t;
}
setMeeting()