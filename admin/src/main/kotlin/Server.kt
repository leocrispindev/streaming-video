import com.google.gson.Gson
import controller.VideoHandler
import controller.VideoHandlerImpl
import io.ktor.application.call
import io.ktor.html.respondHtml
import io.ktor.http.*
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
            +"Hello from Ktor Leonardo"
        }
    }
}

fun main() {

    val handler = VideoHandlerImpl()

    embeddedServer(Netty, port = 8080, host = "127.0.0.1") {
        routing {
            get("/") {
                val response = handler.getAll()
                call.respondText(Gson().toJson(response), ContentType.Application.Json, response.status)
            }

            post("/insert"){
                var body = call.receive<String>()
                var videoDto = Gson().fromJson(body, VideoDTO::class.java)
               val response =  handler.insert(videoDto)

                call.respondText(Gson().toJson(response), ContentType.Application.Json, response.status)
            }

            put("/update"){
                var body = call.receive<String>()
                var videoDto = Gson().fromJson(body, VideoDTO::class.java)
                val response =  handler.update(videoDto)

                call.respondText(Gson().toJson(response), ContentType.Application.Json, response.status)
            }

            delete ("/delete"){
                var id = call.request.queryParameters.get("id")
                val response = handler.delete(id)

                call.respondText(Gson().toJson(response), ContentType.Application.Json, response.status)
            }

        }
    }.start(wait = true)
}