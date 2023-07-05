package com.br

import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import com.br.watcher.services.configureRouting
import com.br.watcher.services.configureSockets

fun main() {
    embeddedServer(Netty, port = 8080, host = "0.0.0.0", module = Application::module)
        .start(wait = true)
}

fun Application.module() {
    configureSockets()
    configureRouting()
}
