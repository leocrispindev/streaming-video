package model.response

import io.ktor.http.*

data class Response(var status: HttpStatusCode, val message: String, val data: Any)
