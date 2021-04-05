package x.x.jwstudy

import android.annotation.SuppressLint
import android.os.Bundle
import android.webkit.WebChromeClient
import android.webkit.WebView
import android.webkit.WebViewClient
import androidx.appcompat.app.AppCompatActivity
import java.io.BufferedReader
import java.io.File
import java.io.InputStreamReader
import java.util.concurrent.Executors


class MainActivity : AppCompatActivity() {
    private lateinit var webview: WebView
    @SuppressLint("SetJavaScriptEnabled")
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        supportActionBar?.hide()
        webview = findViewById(R.id.webview)
        webview.settings.javaScriptEnabled = true
        webview.settings.domStorageEnabled = true
        webview.settings.loadWithOverviewMode = true
        webview.settings.useWideViewPort = true
        webview.settings.builtInZoomControls = true
        webview.settings.displayZoomControls = false
        webview.settings.setSupportZoom(true)
        webview.settings.defaultTextEncodingName = "utf-8"
        webview.webChromeClient = WebChromeClient()
        webview.webViewClient = WebViewClient()
        webview.loadUrl("data:text/html,<a href=\"http://127.0.0.1:8080\">open</a><script>alert('Hey! Thanks for testing JW Study app! Please note that this version is experimantal, and may not work well. Also - from the things that matter, publications are downloaded on the go, so after opening it for the first time wait a bit.')</script><h1>Loading...</h1><script>setInterval(() => {fetch('http://127.0.0.1:8080/api/ping').then(response => response.text()).then(resp => {if (resp == 'pong') {window.location.href = 'http://127.0.0.1:8080'}}); }, 500);</script>")
        rund()
        //webview.loadUrl("data:text/html,<script>alert('Hey! Thanks for testing JW Study app! Please note that this version is experimantal, and may not work well. Also - from the things that matter, publications are downloaded on the go, so after opening it for the first time wait a bit.')</script><h1>Loading...</h1><script>setInterval(() => {fetch('http://127.0.0.1:8080/api/ping').then(response => response.text()).then(resp => {if (resp == 'pong') {window.location.href = 'http://127.0.0.1:8080'}}); }, 500);</script>")
    }
    private fun exec(command: String, params: String): String {
        try {
            val process = ProcessBuilder()
                    .directory(File(filesDir.parentFile!!, "lib"))
                    .command(command, params)
                    .redirectErrorStream(true)
                    .start()
            val reader = BufferedReader(
                    InputStreamReader(process.inputStream)
            )
            val text = reader.readText()
            reader.close()
            process.waitFor()
            return text
        } catch (e: Exception) {
            return e.message ?: "IOException"
        }
    }
    private fun rund() {
        val myPool = Executors.newFixedThreadPool(5)
        myPool.submit { exec("./jwlib.bin", "") }
    }
}