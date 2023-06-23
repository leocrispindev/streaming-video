package model

data class Document(var action : Int, val id: Int, val key: String, val repository: String, var videoInfo: VideoInfo?)
