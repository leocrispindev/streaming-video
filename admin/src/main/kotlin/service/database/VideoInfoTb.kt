package service.database

import org.jetbrains.exposed.dao.id.IntIdTable

object VideoInfoTb : IntIdTable(name = "video_info") {
    val title = varchar("title", 100)
    val synopsis = varchar("synopsis", 255)
    val category = integer("category")
    val duration = integer("duration")
    val indexless = bool("indexless")
    val extension = varchar("extension", 50)
}
