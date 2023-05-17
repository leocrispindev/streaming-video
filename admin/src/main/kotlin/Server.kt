import controller.VideoHandler
import controller.VideoHandlerImpl
import io.ktor.application.call
import io.ktor.html.respondHtml
import io.ktor.http.HttpStatusCode
import io.ktor.request.*
import io.ktor.response.*
import io.ktor.routing.*
import io.ktor.server.engine.embeddedServer
import io.ktor.server.netty.Netty
import kotlinx.html.*
import model.VideoDTO

fun HTML.index() {
    head {
        title("Hello from Ktor!")
    }
    body {
        div {
            +"Hello from Ktor"
        }
    }
}

fun main() {

    val handler = VideoHandlerImpl()

    embeddedServer(Netty, port = 8080, host = "127.0.0.1") {
        routing {
            get("/") {
                call.respondHtml(HttpStatusCode.OK, HTML::index)
            }

            post("/insert"){
                var body = call.receive<VideoDTO>()

               val response =  handler.insert(body)

                call.respond(response.status, response)
            }
        }
    }.start(wait = true)
}