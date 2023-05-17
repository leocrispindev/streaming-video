package model.response

import io.ktor.http.*

data class Response(val status: HttpStatusCode, val message: String, val data: Object)
