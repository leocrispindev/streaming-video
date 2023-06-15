package service

import model.VideoDTO
import model.VideoInfo

interface Admin {

    fun insert(video: VideoDTO): Int
    fun update(video: VideoDTO)

    fun getAll() : ArrayList<VideoInfo>

    fun delete(id : Int)
}