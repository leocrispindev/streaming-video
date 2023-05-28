package controller

import io.ktor.http.*
import model.VideoDTO
import model.response.Response
import service.AdminImpl

class VideoHandlerImpl : VideoHandler{

    val videoAdmin = AdminImpl
    override fun insert(videoDto: VideoDTO): Response {

        try{
            var id = videoAdmin.insert(videoDto)

            return Response(HttpStatusCode.OK, "Success", Any())
        }catch (e : Exception) {
            return Response(HttpStatusCode.BadRequest, e.message.toString(), Any())
        }
    }

    override fun update(videoDto: VideoDTO): Response {
        try{
            videoAdmin.update(videoDto)

            return Response(HttpStatusCode.OK, "Success", Any())
        }catch (e : Exception) {
            return Response(HttpStatusCode.InternalServerError, e.message.toString(), Any())
        }
    }

    override fun getAll(): Response {
        try{
            return Response(HttpStatusCode.OK, "Success", videoAdmin.getAll())
        }catch (e : Exception) {
            return Response(HttpStatusCode.BadRequest, e.message.toString(), Any())
        }
    }

    override fun delete(id : String?): Response {
        try{
            if (id.isNullOrEmpty()) {
                return Response(HttpStatusCode.BadRequest,"Empty ID parameter", Any())
            }

            return Response(HttpStatusCode.OK, "Success", videoAdmin.delete(id.toInt()))

        }catch (e : NumberFormatException) {
            return Response(HttpStatusCode.BadRequest, "Invalid ID parameter", Any())
        }
        catch (e : Exception) {
            return Response(HttpStatusCode.InternalServerError, e.message.toString(), Any())
        }
    }

}