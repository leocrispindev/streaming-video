package service

import model.VideoDTO
import model.VideoInfo

interface Admin {

    fun insert(video: VideoDTO): Int
    fun update(video: VideoDTO)

    fun index(videoInfo: VideoInfo) : Int// Indexa o conteudo

    fun getAll() : ArrayList<VideoInfo>

    fun delete(id : Int)
}