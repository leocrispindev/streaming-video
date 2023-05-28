package service.database

import org.jetbrains.exposed.dao.id.IntIdTable

object VideoInfoTb : IntIdTable(name = "video_info") {
    val titulo = varchar("titulo", 100)
    val descricao = varchar("descricao", 255)
    val category = integer("category")
    val duration = double("duration")
    val indexless = bool("indexless")
}
