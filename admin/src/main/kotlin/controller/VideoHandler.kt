package controller

import model.VideoDTO
import model.response.Response

interface VideoHandler {
    fun insert(videoDto : VideoDTO): Response

}