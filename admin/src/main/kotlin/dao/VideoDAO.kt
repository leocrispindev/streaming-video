package dao

import model.VideoInfo
import org.jetbrains.exposed.sql.SqlExpressionBuilder.eq
import org.jetbrains.exposed.sql.deleteWhere
import org.jetbrains.exposed.sql.insert
import org.jetbrains.exposed.sql.selectAll
import org.jetbrains.exposed.sql.transactions.transaction
import org.jetbrains.exposed.sql.update
import service.database.DbConn
import service.database.VideoInfoTb

class VideoDAO {

    private var db = DbConn.getDatabase()

    fun insert(video: VideoInfo) : Int {

        val insertid =  transaction (db = db){
            VideoInfoTb.insert {
                it[titulo] = video.titulo
                it[descricao] = video.descricao
                it[category] = video.category
                it[indexless] = video.indexless
                it[duration] = video.duration
            } get VideoInfoTb.id

        }

        return insertid.value
    }

    fun update(video : VideoInfo) {
        transaction {
            VideoInfoTb.update(where = {VideoInfoTb.id eq video.id}) {
                it[titulo] = video.titulo
                it[descricao] = video.descricao
                it[category] = video.category
            }
        }
    }

    fun get() : ArrayList<VideoInfo>{

        val query = VideoInfoTb.selectAll()
        val result = ArrayList<VideoInfo>()

        transaction (db= db){
            query.forEach {row ->


                val videoInfo = VideoInfo(
                    id = row[VideoInfoTb.id].value,
                    titulo = row[VideoInfoTb.titulo],
                    descricao = row[VideoInfoTb.descricao],
                    category = row[VideoInfoTb.category],
                    duration = row[VideoInfoTb.duration],
                    indexless = row[VideoInfoTb.indexless]
                )

                result.add(videoInfo)
            }
        }


        return result
    }

    fun delete(id: Int) {
        transaction(db = db) {
            VideoInfoTb.deleteWhere(op = {VideoInfoTb.id eq id})

        }
    }

}