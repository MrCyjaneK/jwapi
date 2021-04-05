import QtQuick 2.7
import Ubuntu.Components 1.3
//import QtQuick.Controls 2.2
import QtQuick.Layouts 1.3
import Qt.labs.settings 1.0
import io.thp.pyotherside 1.3
import QtWebEngine 1.7

MainView {
    id: root
    objectName: 'mainView'
    applicationName: 'jwstudy.anon'
    automaticOrientation: true

    width: units.gu(45)
    height: units.gu(75)

    WebEngineView {
        id: webview
        url: "data:text/html,<script>alert('Hey! Thanks for testing JW Study app! Please note that this version is experimantal, and may not work well. Also - from the things that matter, publications are downloaded on the go, so after opening it for the first time wait a bit.')</script><h1>Loading...</h1><script>setInterval(() => {fetch('http://127.0.0.1:8080/api/ping').then(response => response.text()).then(resp => {if (resp == 'pong') {window.location.href = 'http://127.0.0.1:8080'}}); }, 500);</script>"
        settings.showScrollBars: false
        zoomFactor: 2
        anchors {
            left: parent.left
            top: parent.top
            right: parent.right
            bottom: Qt.inputMethod.visible? showKeyboard.top: parent.bottom
        }
    }
    Python {
        id: python

        Component.onCompleted: {
            addImportPath(Qt.resolvedUrl('../src/'));
            importModule('go', function() {
                console.log('module imported');
                python.call('go.load', [], function(returnValue) {
                    console.log(returnValue);
                })
            });
        }
        onError: {
            console.log('python error: ' + traceback);
        }
    }
}