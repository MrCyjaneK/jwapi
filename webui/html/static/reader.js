function chapterChange(publication, chapter) {
    window.location.href = chapter // atm we use full URL
}
/*
// Swipe to change page:
var container = document.getElementById('bookcontent');

  container.addEventListener("touchstart", startTouch, false);
  container.addEventListener("touchmove", moveTouch, false);

  // Swipe Up / Down / Left / Right
  var initialX = null;
  var initialY = null;

  function startTouch(e) {
    initialX = e.touches[0].clientX;
    initialY = e.touches[0].clientY;
  };

  function moveTouch(e) {
    if (initialX === null) {
      return;
    }

    if (initialY === null) {
      return;
    }

    var currentX = e.touches[0].clientX;
    var currentY = e.touches[0].clientY;

    var diffX = initialX - currentX;
    var diffY = initialY - currentY;

    if (Math.abs(diffX) > Math.abs(diffY)) {
      // sliding horizontally
      if (diffX > 0) {
        // swiped left
        window.location.href = document.getElementById('next').href
      } else {
        // swiped right
        window.location.href = document.getElementById('prev').href
      }  
    } else {
      // sliding vertically
      if (diffY > 0) {
        // swiped up
        console.log("swiped up");
      } else {
        // swiped down
        console.log("swiped down");
      }  
    }

    initialX = null;
    initialY = null;

    e.preventDefault();
  };
*/

function highlightMenu(el) {
  console.log(el)
}
mouseXPosition = 0;
$(document).ready(function () {
  
    $("#bookcontent").mousedown(function (e1) {
        mouseXPosition = e1.pageX;//register the mouse down position
    });

    $("#bookcontent").mouseup(function (e2) {
        var highlighted = false;
        var selection = window.getSelection();
        var selectedText = selection.toString();
        var startPoint = window.getSelection().getRangeAt(0).startOffset;
        var endPoint = window.getSelection().getRangeAt(0).endOffset;
        var anchorTag = selection.anchorNode.parentNode;
        var focusTag = selection.focusNode.parentNode;
        if ((e2.pageX - mouseXPosition) < 0) {
            focusTag = selection.anchorNode.parentNode;
            anchorTag = selection.focusNode.parentNode;
        }
        if (selectedText.length === (endPoint - startPoint)) {
            highlighted = true;

            if (anchorTag.className !== "highlight") {
                highlightSelection();
            } else {
                var afterText = selectedText + "<span class='highlight'>" + anchorTag.innerHTML.substr(endPoint) + "</span>";
                anchorTag.innerHTML = anchorTag.innerHTML.substr(0, startPoint);
                anchorTag.insertAdjacentHTML('afterend', afterText);
            }

        }else{
            if(anchorTag.className !== "highlight" && focusTag.className !== "highlight"){
                highlightSelection();  
                highlighted = true;
            }
            
        }


        if (anchorTag.className === "highlight" && focusTag.className === 'highlight' && !highlighted) {
            highlighted = true;

            var afterHtml = anchorTag.innerHTML.substr(startPoint);
            var outerHtml = selectedText.substr(afterHtml.length, selectedText.length - endPoint - afterHtml.length);
            var anchorInnerhtml = anchorTag.innerHTML.substr(0, startPoint);
            var focusInnerHtml = focusTag.innerHTML.substr(endPoint);
            var focusBeforeHtml = focusTag.innerHTML.substr(0, endPoint);
            selection.deleteFromDocument();
            anchorTag.innerHTML = anchorInnerhtml;
            focusTag.innerHTml = focusInnerHtml;
            var anchorafterHtml = afterHtml + outerHtml + focusBeforeHtml;
            anchorTag.insertAdjacentHTML('afterend', anchorafterHtml);


        }

        if (anchorTag.className === "highlight" && !highlighted) {
            highlighted = true;
var Innerhtml = anchorTag.innerHTML.substr(0, startPoint);
            var afterHtml = anchorTag.innerHTML.substr(startPoint);
            var outerHtml = selectedText.substr(afterHtml.length, selectedText.length);
            selection.deleteFromDocument();
            anchorTag.innerHTML = Innerhtml;
            anchorTag.insertAdjacentHTML('afterend', afterHtml + outerHtml);
         }
        
        if (focusTag.className === 'highlight' && !highlighted) {
            highlighted = true;
var beforeHtml = focusTag.innerHTML.substr(0, endPoint);
            var outerHtml = selectedText.substr(0, selectedText.length - beforeHtml.length);
            selection.deleteFromDocument();
            focusTag.innerHTml = focusTag.innerHTML.substr(endPoint);
            outerHtml += beforeHtml;
            focusTag.insertAdjacentHTML('beforebegin', outerHtml );


        }
        if (!highlighted) {
            highlightSelection();
        }
        $('.highlight').each(function(){
            if($(this).html() == ''){
                $(this).remove();
            }
        });
        selection.removeAllRanges();
    });
});

function highlightSelection() {
    var selection;

    //Get the selected stuff
    if (window.getSelection)
        selection = window.getSelection();
    else if (typeof document.selection != "undefined")
        selection = document.selection;
    console.log(selection)
    //Get a the selected content, in a range object
    var range = selection.getRangeAt(0);

    //If the range spans some text, and inside a tag, set its css class.
    if (range && !selection.isCollapsed) {
        if (selection.anchorNode.parentNode == selection.focusNode.parentNode) {
            var span = document.createElement('span');
            span.className = 'highlight';
            span.textContent = selection.toString();
            selection.deleteFromDocument();
            range.insertNode(span);
//                        range.surroundContents(span);
        }
    }
}
