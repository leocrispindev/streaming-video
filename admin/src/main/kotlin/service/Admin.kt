package service

import model.VideoInfo

interface Admin {

    fun insert(videoInfo: VideoInfo): Int
    fun update(videoInfo: VideoInfo): Int // Deve retornar o objeto

    fun upload() // Realiza o upload do arquivo

    fun index(videoInfo: VideoInfo) : Int// Indexa o conteudo
}