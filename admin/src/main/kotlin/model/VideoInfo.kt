package model

data class VideoInfo(
    val id: Int?,
    val titulo: String,
    var descricao: String,
    val category: Int,
    val duration: Double,
    var indexless: Boolean
) {
    constructor(id: Int?, titulo: String, descricao: String, category: Int) :
            this(id, titulo, descricao, category, 0.0, false)
}
