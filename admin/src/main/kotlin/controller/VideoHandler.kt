package controller

import model.VideoDTO
import model.response.Response

interface VideoHandler {
    fun insert(videoDto : VideoDTO): Response

    fun update(videoDto: VideoDTO) : Response

    fun getAll() : Response

    fun delete(id : String?): Response
}