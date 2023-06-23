package model

data class VideoInfo(
    var id: Int?,
    val title: String,
    var synopsis: String,
    val category: Int,
    val duration: Int,
    var indexless: Boolean,
    var extension: String
) {
    constructor(id: Int?, title: String, synopsis: String, category: Int, extension: String) :
            this(id, title, synopsis, category, 0, false, extension)
}
